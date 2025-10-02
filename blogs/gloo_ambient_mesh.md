# From Istio Sidecars to Gloo Ambient Mesh: Simplifying Service Mesh at Scale

## The Sidecar Tax

If you've run Istio at scale, you know the pain. Every pod gets an Envoy sidecar. Every deployment doubles your resource consumption. Every application restart requires coordination between your app container and the sidecar. The operational complexity compounds with each service you onboard.

We were running Istio across multiple EKS clusters, handling 500K+ TPS, and the sidecar model was bleeding us dry—not just in infrastructure costs, but in operational overhead and developer friction.

Then we discovered Gloo Ambient Mesh.

## Why We Had to Change

Our Istio deployment was showing cracks:

**Resource overhead was brutal.** Each pod ran two containers—the application and the Envoy sidecar. At our scale, we were running thousands of sidecars consuming CPU and memory that could have been used for actual workloads. Our cluster costs were 30-40% higher than they needed to be.

**Application lifecycle management was painful.** Coordinating shutdowns between application containers and sidecars led to race conditions. Rolling updates became complicated choreography. Developers had to understand sidecar behavior to troubleshoot application issues.

**Cross-cluster communication was a nightmare.** We needed our services to communicate across EKS clusters using Pulsar for pub-sub messaging. With Istio sidecars, the networking complexity was overwhelming. Each sidecar needed configuration for cross-cluster service discovery, certificate management was a distributed mess, and DAPR sidecars added another layer of proxies to the stack.

Something had to give.

## Enter Gloo Ambient Mesh

Gloo Ambient Mesh takes a fundamentally different approach. Instead of injecting sidecars into every pod, it runs service mesh functionality at the node level using two components:

**Ztunnel** (Zero Trust Tunnel) handles L4 networking—mTLS, traffic routing, and telemetry—without modifying application pods. It runs as a DaemonSet, one instance per node.

**Waypoint proxies** provide L7 capabilities—advanced routing, retries, circuit breaking—but only for services that need them. They're deployed on-demand, not universally.

The promise was compelling: dramatically reduced resource consumption, simplified operational model, and easier integration with our existing DAPR and Pulsar infrastructure.

## The Migration Trade-offs

Before we committed, we mapped out what we'd gain and lose:

| Dimension | Istio Sidecars | Gloo Ambient Mesh | Impact |
|-----------|----------------|-------------------|---------|
| **Resource Model** | Sidecar per pod | Ztunnel per node | 60-70% reduction in proxy resource usage |
| **Pod Count** | 2 containers per pod | 1 container per pod | Simplified scheduling, faster startups |
| **mTLS Overhead** | In-pod proxy termination | Node-level termination | Slightly higher network latency (~1-2ms) |
| **L7 Features** | Always available | Optional waypoint proxies | Deploy complexity only where needed |
| **Config Complexity** | Per-service sidecar annotations | Mesh-wide policies + selective waypoints | Easier to manage at scale |
| **Upgrade Path** | Rolling restart of all services | Update DaemonSet, services unaffected | Zero-downtime mesh upgrades |

The trade-offs were clear. We'd gain massive resource efficiency and operational simplicity. We'd lose some per-pod isolation and potentially add minimal network latency. For our use case, this was a no-brainer.

## The Migration Strategy

We couldn't flip a switch and move thousands of services overnight. We needed a gradual, low-risk migration path.

**Phase One: Parallel Run**

We deployed Gloo Ambient Mesh alongside our existing Istio infrastructure. Both service meshes ran in parallel, with careful namespace segregation. This let us test Ambient Mesh behavior without risking production traffic.

We started with non-critical services—internal tools, development environments, and low-traffic APIs. These guinea pig services helped us understand Ambient Mesh's operational characteristics and identify gaps in our runbooks.

**Phase Two: DAPR Integration Testing**

Our DAPR sidecars needed to work seamlessly with Ambient Mesh. This was critical because DAPR handled our Pulsar pub-sub abstraction. We had to verify:

- DAPR sidecar-to-sidecar communication through Ztunnel
- mTLS certificate trust chains between DAPR and Ambient Mesh
- Service discovery when DAPR components talked across the mesh
- Performance characteristics under load

The good news? DAPR didn't care about the underlying service mesh implementation. As long as network policies and mTLS were configured correctly, DAPR worked transparently. This validated our architectural decision to use DAPR as an abstraction layer.

**Phase Three: Progressive Service Migration**

We migrated services in waves, starting with those that would benefit most from reduced resource consumption. The process was remarkably simple:

Remove Istio sidecar injection annotations from deployments, redeploy the service (now it's just the application container), and Ztunnel automatically handles mesh traffic for the pod. No code changes. No configuration rewrites.

For services needing L7 features, we deployed waypoint proxies strategically. Only about 20% of our services actually needed advanced routing capabilities—the rest worked perfectly with L4-only Ztunnel.

**Phase Four: Cross-Cluster Communication**

This is where things got interesting. We needed services in different EKS clusters to communicate via Pulsar, all secured by the service mesh.

## Enabling Cross-Cluster Pulsar with DAPR

Our architecture had multiple EKS clusters running workloads. Services needed to publish and subscribe to Pulsar topics regardless of which cluster they lived in. Previously, with Istio sidecars, this required complex cross-cluster service mesh federation.

Gloo Ambient Mesh made this dramatically simpler.

### The Pulsar Proxy Pattern

We deployed Pulsar proxy instances in each cluster. These proxies acted as local gateways to our centralized Pulsar clusters. Services in any cluster would:

- Call DAPR's pub-sub API (no knowledge of Pulsar)
- DAPR would connect to the local Pulsar proxy
- The proxy would route messages to the appropriate Pulsar broker
- All traffic would flow through Ztunnel with mTLS

The beauty of this design was layering:

**DAPR abstracted the pub-sub semantics.** Services didn't know they were using Pulsar. They certainly didn't know about proxies or cross-cluster routing.

**Pulsar proxies handled multi-cluster routing.** They knew which brokers owned which topics and routed accordingly. They handled connection pooling and load balancing.

**Ambient Mesh secured everything.** Every connection—service to DAPR, DAPR to proxy, proxy to broker—was mTLS encrypted and authenticated by Ztunnel.

### The Configuration Simplification

Here's what made this architecture powerful: **we eliminated per-service configuration complexity.**

With Istio sidecars, each service needed configuration for:
- Service discovery across clusters
- Certificate management for cross-cluster mTLS
- Egress rules for external Pulsar brokers
- Traffic routing policies for multi-cluster scenarios

With Gloo Ambient Mesh and Pulsar proxies, services needed... nothing. They called DAPR's API. Everything else was infrastructure-level configuration that we managed centrally.

| Configuration Aspect | Istio Sidecar Approach | Gloo Ambient + Pulsar Proxy | Improvement |
|----------------------|------------------------|----------------------------|-------------|
| **Per-Service Config** | ServiceEntry, DestinationRule, VirtualService | None—mesh handles it | 90% reduction in service-specific config |
| **Cross-Cluster mTLS** | Manual certificate distribution | Automatic via Ztunnel | Zero manual certificate management |
| **Pulsar Connection** | Each service connects directly | Proxy handles connection pooling | Reduced connection overhead |
| **Service Discovery** | Requires explicit multi-cluster setup | Local proxy, centralized routing | Simplified network topology |
| **Failure Isolation** | Sidecar failures affect specific pods | Node-level failures, blast radius contained | Improved fault isolation |

### Performance Implications

We were nervous about performance. Adding a proxy layer could introduce latency. Moving mTLS termination from pod-level to node-level might slow things down.

| Performance Metric | Istio Sidecars | Gloo Ambient + Pulsar Proxy | Delta |
|--------------------|----------------|----------------------------|-------|
| **P50 Pub-Sub Latency** | 11ms | 13ms | +2ms (18% increase) |
| **P99 Pub-Sub Latency** | 52ms | 54ms | +2ms (4% increase) |
| **Cross-Cluster Latency** | 28ms | 30ms | +2ms (7% increase) |
| **CPU per Pod** | 0.8 cores (app + sidecar) | 0.5 cores (app only) | -0.3 cores (37% reduction) |
| **Memory per Pod** | 768MB | 512MB | -256MB (33% reduction) |
| **Cluster-wide CPU** | Baseline | -35% | Massive savings |
| **Connection Count** | Per-pod connections to Pulsar | Pooled via proxy | 80% reduction |

The latency increase was minimal—2ms at P50, negligible at higher percentiles. But the resource savings were game-changing. We reduced cluster-wide CPU consumption by 35% and memory by a similar margin.

The Pulsar proxy pattern actually improved connection efficiency. Instead of every pod maintaining connections to Pulsar brokers, proxies pooled connections. At our scale, this reduced overall connection count by 80%.

## Operational Wins

Beyond raw performance, the operational improvements were substantial:

**Mesh upgrades became trivial.** Previously, upgrading Istio meant rolling restarts of every service to get new sidecar versions. With Ambient Mesh, we updated the Ztunnel DaemonSet and services continued running unchanged.

**Debugging became easier.** With sidecars, tracing network issues required understanding three components: the app, the sidecar, and the remote service. With Ambient Mesh, network behavior was centralized at the node level. We could troubleshoot connectivity issues without restarting application pods.

**Developer onboarding simplified.** New services got mesh capabilities automatically. No annotations to remember, no sidecar behavior to explain. Deploy a pod, Ztunnel handles the rest.

**Security posture improved.** Centralized policy enforcement at the node level meant consistent security controls across all workloads. We couldn't accidentally deploy a pod without mesh protection—Ztunnel covered everything on the node.

## The DAPR Synergy

The combination of DAPR, Gloo Ambient Mesh, and Pulsar proxies created an elegant architecture:

**DAPR provided application-level abstraction.** Services called standard pub-sub APIs without knowing about Pulsar, proxies, or mesh networking.

**Pulsar proxies provided messaging-level routing.** They understood Pulsar's topic topology and handled cross-cluster message routing.

**Ambient Mesh provided network-level security.** mTLS encryption and authentication happened transparently at the infrastructure layer.

Each layer did one thing well. Each layer was independently upgradeable. The system was loosely coupled yet tightly integrated.

## Lessons Learned

**Sidecars aren't always necessary.** For many workloads, node-level service mesh capabilities are sufficient. Deploy L7 proxies only where you need advanced features.

**Abstraction layers compound value.** DAPR abstracting Pulsar plus Ambient Mesh abstracting networking created a system where application teams could move fast without understanding infrastructure complexity.

**Migration can be gradual.** We didn't need a big-bang cutover. Running both service meshes in parallel let us migrate at our own pace with minimal risk.

**Resource efficiency matters at scale.** The 35% reduction in cluster resources paid for the entire migration effort in three months.

**Simplicity is a feature.** Reducing per-service configuration complexity made our platform easier to operate and reduced the chance of misconfigurations.

## Six Months Later

We're now running entirely on Gloo Ambient Mesh across our EKS clusters, handling 550K+ TPS with DAPR and Pulsar. Cross-cluster communication works flawlessly. Our infrastructure costs are down 30%. Developer satisfaction is up because they're not fighting service mesh complexity.

The move from Istio sidecars to Gloo Ambient Mesh wasn't just a technical upgrade—it was a philosophical shift. We moved from "every pod needs its own proxy" to "the infrastructure handles networking." We traded some per-pod isolation for massive operational simplicity.

And we've never looked back.

---

*Considering Ambient Mesh for your platform? The resource savings alone make it worth evaluating.*
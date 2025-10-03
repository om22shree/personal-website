# From Kafka Lessons to Pulsar: Building Our New Platform Right

## The Platform Nobody Wanted to Rebuild

Let's be honest, building a new messaging infrastructure from scratch while your legacy system handles 500K+ transactions per second is the kind of project that makes engineering leaders question their career choices. When we decided to build our next-generation platform on Pulsar instead of Kafka across our EKS clusters, the immediate question wasn't "why?" but "how do we avoid repeating past mistakes?"

The answer turned out to be DAPR, but not for the reasons we initially thought.

## Why Not Kafka This Time?

Kafka served our legacy platform well for years, but the operational scars run deep. Zookeeper management became a constant source of pain—cluster coordination failures, split-brain scenarios during network partitions, and the endless operational overhead of keeping it healthy. Multi-tenancy was a nightmare of ACLs and topic management. Geo-replication required custom tooling that our team constantly had to maintain.

For our new platform, we had a clean slate. We could learn from these lessons without the burden of migration.

Pulsar offered everything we wished Kafka had: native multi-tenancy, built-in geo-replication, and most importantly, a clean separation of compute and storage. No Zookeeper dependency. The business case was clear. The technical risk? Still significant, but manageable since we were building new rather than migrating.

We had a greenfield opportunity to build 200+ new microservices the right way. Instead of tightly coupling them to Pulsar clients from day one, we needed an abstraction layer that would give us flexibility as requirements evolved. That's where DAPR entered the picture.

## DAPR as the Foundation

DAPR's pub-sub abstraction became our architectural foundation. Instead of services directly importing Pulsar client libraries, they would call DAPR's HTTP or gRPC APIs. The underlying messaging system became an implementation detail—a swappable component.

This abstraction wasn't free, though. Let's break down the trade-offs:

| Aspect | What We Gained | What We Lost |
|--------|----------------|--------------|
| **Portability** | Switch messaging backends without code changes | Direct access to platform-specific features |
| **Development Velocity** | Uniform API across all services, no client library management | Some advanced client-side optimizations |
| **Operational Complexity** | Centralized configuration and policy management | Additional sidecar to monitor and manage |
| **Performance** | Consistent behavior and retry patterns | 3-5ms additional latency per operation |
| **Debugging** | Standardized observability hooks | Extra abstraction layer to understand |

The strategy was elegant in its simplicity. New services would be built using DAPR's pub-sub API from the start. Behind the scenes, DAPR would communicate with our Pulsar clusters. If we ever needed to switch backends or run hybrid scenarios, we could do so without rewriting application code.

We traded direct client control for portability and future flexibility. We lost some of Pulsar's advanced features like direct access to BookKeeper primitives. But what we gained was priceless: the ability to evolve our messaging infrastructure independently from application code.

## The Istio mTLS Complexity

Our EKS environment runs Istio with strict mTLS enforcement—every service-to-service call encrypted and authenticated. Adding DAPR sidecars into this mesh created a puzzle of certificate chains and trust boundaries.

DAPR has its own sidecar model, and Istio has its own. When you nest them together, you create multiple layers of proxies, each with its own security model. The performance implications alone could threaten our 500K+ TPS target.

| Security Consideration | Without DAPR + Istio | With DAPR + Istio |
|------------------------|---------------------|---------------------|
| **Authentication** | Manual service account management | Automated mTLS certificate rotation |
| **Authorization** | Application-level API keys | Network policy + service mesh rules |
| **Encryption** | TLS termination at application | End-to-end encryption across all hops |
| **Audit Trail** | Custom logging per service | Centralized access logs via Envoy |
| **Resource Overhead** | Single container per pod | 2 sidecars per pod (DAPR + Envoy) |
| **Network Latency** | Direct service-to-service | Additional proxy hops (~2-3ms each) |

We had to make a critical architectural decision: embrace the ambient mesh pattern and automate the hell out of it. Instead of fighting DAPR and Istio, we made them collaborators. This meant accepting some overhead—two sidecars per pod, additional network hops, more resource consumption—but gaining automated mutual authentication and zero-trust networking without touching application code.

The trade-off was worth it. Yes, we consume more CPU and memory per pod. Yes, latency increased by a few milliseconds. But we eliminated entire categories of security vulnerabilities and made our infrastructure auditable by default.

## The Build Phases

**Phase One** was establishing the foundation—setting up Pulsar clusters with proper namespace isolation, configuring BookKeeper for optimal storage performance, and establishing our DAPR component configurations. We started small with internal tools and non-critical services to validate our architecture.

**Phase Two** involved building our first production services. These early adopters helped us identify pain points: debugging became trickier with abstractions, and understanding message flow required new observability tools. We learned to tune DAPR's message delivery semantics, configure retry policies, and handle backpressure properly.

**Phase Three** was scaling out. As more teams adopted the platform, we refined our patterns, created reusable templates, and built comprehensive documentation. Services that followed our established patterns came online smoothly. Those that tried to work around the abstraction layer struggled.

**Phase Four** is where we are now—onboarding high-volume services and validating performance at scale. The architecture decisions we made early are paying off, but we're also discovering new challenges as we push the limits.

## Observability: Seeing Through the Abstraction

From day one, we knew observability would make or break this platform. We couldn't just check application logs—we needed to understand what was happening inside DAPR, between DAPR and Istio, between DAPR and Pulsar, and across the entire distributed trace.

But here's the kicker: **observability platforms are expensive**. We needed a cost-effective solution that didn't sacrifice visibility.

### The Cost-Optimized Observability Stack

We built our observability around three core components: **Prometheus for metrics, OpenTelemetry for instrumentation, and Cribl for intelligent data routing to Splunk**. This architecture keeps our licensing costs manageable while giving us operational insights we never had with Kafka.

**Prometheus as the Metrics Foundation**

DAPR exposes Prometheus-compatible metrics out of the box, which became our primary data source. We instrumented three critical metric streams:

First, **DAPR custom metrics**. We configured DAPR to expose granular operational metrics including message throughput per topic and per service, pub-sub operation latency broken down by publish, subscribe, and acknowledgment phases, sidecar resource utilization tracking CPU, memory, and network I/O, component health status for all connections, and retry and circuit breaker state to catch degradation patterns.

Second, **Pulsar custom metrics**. We instrumented our Pulsar brokers to expose topic-level statistics, broker resource utilization, replication lag for our geo-distributed setup, consumer lag metrics to identify bottlenecks, and producer acknowledgment latency. These metrics were scraped by Prometheus and correlated with DAPR metrics to give us complete pub-sub visibility.

Third, **application-level metrics**. Services expose their own business metrics through Prometheus client libraries, making it possible to correlate business events with infrastructure performance.

The key insight was granularity without explosion. We needed per-service, per-topic metrics to isolate issues quickly, but we couldn't afford to store every metric forever. We implemented tiered retention: high-resolution metrics for 7 days, downsampled metrics for 30 days, and aggregated metrics for 1 year.

**OpenTelemetry for Unified Instrumentation**

We adopted OpenTelemetry as our instrumentation layer, which gave us vendor-agnostic telemetry collection. DAPR's native OpenTelemetry support meant minimal configuration—we just enabled the OTel collector sidecar and configured exporters.

The challenge was trace and metric correlation at scale. With DAPR and Istio both in the proxy path, a single pub-sub operation can generate six or more spans. We configured OpenTelemetry to propagate trace context through the entire stack: Service A → DAPR sidecar → Envoy proxy → Pulsar broker → Envoy proxy → DAPR sidecar → Service B.

At 500K+ TPS, collecting every trace would bankrupt us. We implemented smart sampling: tail-based sampling that captures 1% of healthy transactions but 100% of errors and slow requests, head-based sampling for specific high-value flows we always want to see, and adaptive sampling that increases collection rates when error rates spike.

**Cribl: The Cost-Saving Game Changer**

This is where we control licensing costs. Cribl sits between our telemetry sources and Splunk, acting as an intelligent data router and processor. Here's what it enables:

We filter aggressively. Cribl drops noisy, low-value logs before they hit Splunk. Debug logs that developers need for troubleshooting go to S3 buckets instead of expensive Splunk indexes. Health check logs and routine polling operations are filtered entirely—we track their success rates in Prometheus instead.

We enrich strategically. Cribl enriches critical events with additional context before forwarding to Splunk, making investigations faster and reducing query complexity. This means fewer Splunk searches and lower computational costs.

We route intelligently. Security-relevant events go to Splunk for compliance and audit trails. Performance metrics go to Prometheus for real-time alerting and dashboards. Historical data for cold storage goes to S3 with lifecycle policies.

The result keeps our costs manageable while actually improving our operational insights compared to our legacy Kafka platform.

### Observability Trade-offs

| Observability Layer | Benefit | Cost | Cost Mitigation |
|---------------------|---------|------|-----------------|
| **Prometheus Metrics** | Real-time dashboards, instant alerting | ~5% CPU overhead per sidecar | Tiered retention reduces storage costs by 70% |
| **OpenTelemetry Traces** | End-to-end visibility across services | Trace storage and processing overhead | Smart sampling keeps costs manageable at scale |
| **Cribl Filtering** | 60% reduction in potential Splunk costs | Initial setup complexity, maintenance overhead | ROI achieved quickly, ongoing savings massive |
| **Custom Dashboards** | Rapid incident response | Engineering time to build and maintain | Reusable templates across teams amortize cost |
| **Automated Alerting** | Proactive issue detection | Alert fatigue if not tuned | Prometheus alert manager with smart routing |

We built custom Grafana dashboards that unified metrics from DAPR, Pulsar, and application layers:

- Pub-sub latency percentiles (P50, P90, P99) correlated with Pulsar broker health
- Message delivery success rates broken down by error category and topic
- DAPR component status across all namespaces with automatic anomaly detection
- Resource utilization trends with predictive alerts for capacity planning
- Comparative views showing performance across different service teams

The observability investment was significant—roughly 20% of our platform effort went into instrumentation and tooling. But it's paying dividends. We catch issues in pre-production before they affect users. We optimize configurations based on real data rather than guesswork. And we have visibility we never had with our legacy Kafka infrastructure.

## Performance Trade-offs at Scale

The abstraction layer costs us performance—there's no way around it. Direct Pulsar producers could achieve lower latency than going through DAPR's sidecar. We measured the impact carefully:

| Performance Metric | Target | Current | Status |
|--------------------|--------|---------|--------|
| **P50 Latency** | <15ms | 11ms | ✓ Exceeding target |
| **P99 Latency** | <60ms | 52ms | ✓ Exceeding target |
| **Max Throughput** | 500K TPS | 550K TPS | ✓ Exceeding target |
| **CPU per Pod** | <1.0 cores | 0.8 cores | ✓ Within budget |
| **Memory per Pod** | <1GB | 768MB | ✓ Within budget |

But here's what we gained: operational simplicity. Our developers don't worry about client library versions, connection pooling, and serialization formats. They call an HTTP endpoint. DAPR handles the rest.

We also gained flexibility. If Pulsar has issues, we can route specific topics to alternative backends without code changes. When we need to implement complex routing rules, we do it in DAPR configuration rather than scattered across codebases.

The resource overhead is significant—an extra sidecar per pod, each consuming memory and CPU. But developer velocity is dramatically higher than our legacy platform, and we're building operational muscle we never had with Kafka.

Interestingly, maximum throughput exceeds our targets. DAPR's built-in connection pooling and message batching optimizations help certain high-volume producers achieve excellent performance.

## The Istio Security Dividend

Strict mTLS with automated ambient mesh enrollment means every message passing through our system is encrypted and authenticated. We don't have to implement custom authentication logic. We don't have to rotate API keys or manage service credentials manually.

The trade-off? Complexity. Debugging network issues is harder. Understanding traffic flow requires deep knowledge of both DAPR and Istio. Our observability stack had to be comprehensive from day one.

But for a system handling sensitive transactions at massive scale, the security posture is non-negotiable. We're passing compliance audits that would have been nightmares without mTLS everywhere.

## Lessons from Building New

**Start with abstractions.** Building on DAPR from day one means we never had to retrofit abstractions later. Our services are inherently portable.

**Observability isn't optional.** With multiple layers of abstraction, understanding system behavior required massive upfront investment in tracing, metrics, and logging. We can't debug by reading code—we read distributed traces. This investment became our most valuable asset.

**Pulsar's architecture matters.** The separation of compute and storage is a game-changer operationally. No Zookeeper means one less complex distributed system to manage. BookKeeper's predictable performance makes capacity planning actually possible.

**Automation is foundational.** Manually configuring Istio and DAPR for 200+ services would be impossible. GitOps and infrastructure-as-code weren't afterthoughts—they were day-one requirements.

**Performance testing validates architecture.** We load-tested every phase in staging environments that mirror production. This validated our architecture decisions and gave us confidence to scale.

## Was It Worth It?

We're now handling 550K+ TPS during peak loads on our new platform. Developer satisfaction is high because they're not wrestling with client libraries or Zookeeper issues. Operational burden is manageable because we built observability and automation from the start.

The trade-offs were real—performance overhead, complexity, learning curve, significant observability investment. But the strategic value of building our platform on proven abstractions and learning from past mistakes has positioned us for sustainable growth.

Our observability stack gives us visibility that was impossible with our legacy Kafka infrastructure. When issues occur, we diagnose them in minutes instead of hours. Our on-call engineers actually sleep through the night.

Sometimes building new is better than migrating old. And sometimes, investing heavily in observability and abstraction from day one is what separates successful platforms from legacy nightmares.

---

*Building a new messaging platform? Learn from operational pain before choosing your stack, and invest in abstractions and observability from day one.*
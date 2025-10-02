# From Kafka to Pulsar: How DAPR Saved Our Migration at 500K+ TPS

## The Migration Nobody Wanted

Let's be honest—migrating your entire messaging infrastructure while maintaining 500K+ transactions per second in production is the kind of project that makes engineering leaders lose sleep. When we decided to move from Kafka to Pulsar across our EKS clusters, the immediate question wasn't "why?" but "how do we not break everything?"

The answer turned out to be DAPR, but not for the reasons we initially thought.

## Why Leave Kafka Behind?

Kafka served us well for years, but as we scaled, the cracks began to show. Multi-tenancy was a nightmare of ACLs and topic management. Geo-replication required custom tooling that our team constantly had to maintain. The operational overhead of managing Zookeeper clusters was becoming untenable.

Pulsar promised native multi-tenancy, built-in geo-replication, and a cleaner separation of compute and storage. The business case was clear. The technical risk? Astronomical.

We had over 200 microservices tightly coupled to Kafka clients. A big-bang migration was out of the question. A gradual migration meant running dual infrastructure—expensive and complex. We needed a decoupling layer, and that's where DAPR entered the picture.

## DAPR as the Decoupling Agent

DAPR's pub-sub abstraction became our secret weapon. Instead of services directly importing Kafka or Pulsar client libraries, they would call DAPR's HTTP or gRPC APIs. The underlying messaging system became an implementation detail—a swappable component.

This abstraction wasn't free, though. Let's break down the trade-offs:

| Aspect | What We Gained | What We Lost |
|--------|----------------|--------------|
| **Portability** | Switch messaging backends without code changes | Direct access to platform-specific features |
| **Development Velocity** | Uniform API across all services, no client library management | Some advanced client-side optimizations |
| **Operational Complexity** | Centralized configuration and policy management | Additional sidecar to monitor and manage |
| **Performance** | Consistent behavior and retry patterns | 3-5ms additional latency per operation |
| **Debugging** | Standardized observability hooks | Extra abstraction layer to understand |

The strategy was elegant in its simplicity. Services would be refactored to use DAPR's pub-sub API. Behind the scenes, we'd initially point DAPR at our Kafka clusters. Once a service was "DAPRified," we could switch its backend to Pulsar independently, test thoroughly, and move on to the next service.

We traded direct client control for portability. We lost some of Kafka's advanced features like exactly-once semantics at the client level. But what we gained was priceless: the ability to migrate services incrementally without rewriting application code.

## The Istio mTLS Complexity

Our EKS environment ran Istio with strict mTLS enforcement—every service-to-service call encrypted and authenticated. Adding DAPR sidecars into this mesh created a puzzle of certificate chains and trust boundaries.

DAPR has its own sidecar model, and Istio has its own. When you nest them together, you create multiple layers of proxies, each with its own security model. The performance implications alone could have killed our 500K+ TPS target.

| Security Consideration | Before DAPR + Istio | After DAPR + Istio |
|------------------------|---------------------|---------------------|
| **Authentication** | Manual service account management | Automated mTLS certificate rotation |
| **Authorization** | Application-level API keys | Network policy + service mesh rules |
| **Encryption** | TLS termination at application | End-to-end encryption across all hops |
| **Audit Trail** | Custom logging per service | Centralized access logs via Envoy |
| **Resource Overhead** | Single container per pod | 2 sidecars per pod (DAPR + Envoy) |
| **Network Latency** | Direct service-to-service | Additional proxy hops (~2-3ms each) |

We had to make a critical architectural decision: embrace the ambient mesh pattern and automate the hell out of it. Instead of fighting DAPR and Istio, we made them collaborators. This meant accepting some overhead—two sidecars per pod, additional network hops, more resource consumption—but gaining automated mutual authentication and zero-trust networking without touching application code.

The trade-off was worth it. Yes, we consumed more CPU and memory per pod. Yes, latency increased by a few milliseconds. But we eliminated entire categories of security vulnerabilities and made our infrastructure auditable by default.

## The Migration Phases

**Phase One** was the hardest—convincing teams to adopt DAPR when Kafka was working fine. We needed buy-in, so we started with new services and non-critical workloads. These early adopters helped us identify pain points: debugging became trickier with abstractions, and understanding message flow required new observability tools.

**Phase Two** involved migrating high-volume services still pointing at Kafka through DAPR. We learned to tune DAPR's message delivery semantics, configure retry policies, and handle backpressure properly. The abstraction layer introduced new failure modes we hadn't considered. Services that previously failed fast now had DAPR retry logic to contend with.

**Phase Three** was the actual backend swap. Service by service, we updated DAPR component configurations from Kafka to Pulsar. This is where the decoupling paid off. Most services required zero code changes. We could test each migration in isolation, roll back easily if needed, and move at our own pace.

**Phase Four** involved decommissioning Kafka infrastructure. This took longer than expected because of hidden dependencies and legacy consumers we'd forgotten about. But eventually, we were fully on Pulsar.

## Observability: Seeing Through the Abstraction

The moment we introduced DAPR sidecars, our existing observability stack became inadequate. We couldn't just check application logs anymore—we needed to understand what was happening inside DAPR, between DAPR and Istio, and across the entire distributed trace.

But here's the kicker: **observability platforms are expensive**. Our Splunk licensing costs were becoming unsustainable, and we needed a cost-effective solution that didn't sacrifice visibility.

### The Cost-Optimized Observability Stack

We built our observability around three core components: **Prometheus for metrics, OpenTelemetry for instrumentation, and Cribl for intelligent data routing to Splunk**. This architecture cut our licensing costs by 60% while actually improving our operational insights.

**Prometheus as the Metrics Foundation**

DAPR exposes Prometheus-compatible metrics out of the box, which became our primary data source. We instrumented three critical metric streams:

First, **DAPR custom metrics**. We configured DAPR to expose granular operational metrics including message throughput per topic and per service, pub-sub operation latency broken down by publish, subscribe, and acknowledgment phases, sidecar resource utilization tracking CPU, memory, and network I/O, component health status for all connections, and retry and circuit breaker state to catch degradation patterns.

Second, **Pulsar custom metrics**. We instrumented our Pulsar brokers to expose topic-level statistics, broker resource utilization, replication lag for our geo-distributed setup, consumer lag metrics to identify bottlenecks, and producer acknowledgment latency. These metrics were scraped by Prometheus and correlated with DAPR metrics to give us complete pub-sub visibility.

Third, **application-level metrics**. Services exposed their own business metrics through Prometheus client libraries, making it possible to correlate business events with infrastructure performance.

The key insight was granularity without explosion. We needed per-service, per-topic metrics to isolate issues quickly, but we couldn't afford to store every metric forever. We implemented tiered retention: high-resolution metrics for 7 days, downsampled metrics for 30 days, and aggregated metrics for 1 year.

**OpenTelemetry for Unified Instrumentation**

We adopted OpenTelemetry as our instrumentation layer, which gave us vendor-agnostic telemetry collection. DAPR's native OpenTelemetry support meant minimal configuration—we just enabled the OTel collector sidecar and configured exporters.

The challenge was trace and metric correlation at scale. With DAPR and Istio both in the proxy path, a single pub-sub operation could generate six or more spans. We configured OpenTelemetry to propagate trace context through the entire stack: Service A → DAPR sidecar → Envoy proxy → Pulsar broker → Envoy proxy → DAPR sidecar → Service B.

At 500K+ TPS, collecting every trace would have bankrupted us. We implemented smart sampling: tail-based sampling that captured 1% of healthy transactions but 100% of errors and slow requests, head-based sampling for specific high-value flows we always wanted to see, and adaptive sampling that increased collection rates when error rates spiked.

**Cribl: The Cost-Saving Game Changer**

This is where we slashed licensing costs. Cribl sits between our telemetry sources and Splunk, acting as an intelligent data router and processor. Here's what it enabled:

We filtered aggressively. Cribl dropped noisy, low-value logs before they hit Splunk. Debug logs that developers needed for troubleshooting went to S3 buckets instead of expensive Splunk indexes. Health check logs and routine polling operations were filtered entirely—we tracked their success rates in Prometheus instead.

We enriched strategically. Cribl enriched critical events with additional context before forwarding to Splunk, making investigations faster and reducing query complexity. This meant fewer Splunk searches and lower computational costs.

We routed intelligently. Security-relevant events went to Splunk for compliance and audit trails. Performance metrics went to Prometheus for real-time alerting and dashboards. Historical data for cold storage went to S3 with lifecycle policies.

The result was dramatic: we reduced Splunk data ingestion by 60%, cut query load by 40%, and actually improved mean time to resolution because engineers weren't drowning in noise.

### Observability Trade-offs

| Observability Layer | Benefit | Cost | Cost Mitigation |
|---------------------|---------|------|-----------------|
| **Prometheus Metrics** | Real-time dashboards, instant alerting | ~5% CPU overhead per sidecar | Tiered retention reduces storage costs by 70% |
| **OpenTelemetry Traces** | End-to-end visibility across services | Trace storage and processing overhead | Smart sampling keeps costs manageable at scale |
| **Cribl Filtering** | 60% reduction in Splunk licensing costs | Initial setup complexity, maintenance overhead | ROI achieved in 3 months, ongoing savings massive |
| **Custom Dashboards** | Rapid incident response | Engineering time to build and maintain | Reusable templates across teams amortized cost |
| **Automated Alerting** | Proactive issue detection | Alert fatigue if not tuned | Prometheus alert manager with smart routing |

We built custom Grafana dashboards that unified metrics from DAPR, Pulsar, and application layers:

- Pub-sub latency percentiles (P50, P90, P99) correlated with Pulsar broker health
- Message delivery success rates broken down by error category and topic
- DAPR component status across all namespaces with automatic anomaly detection
- Resource utilization trends with predictive alerts for capacity planning
- Before-migration vs. after-migration comparison views showing the impact

The observability investment was significant—roughly 20% of our migration effort went into instrumentation and tooling. But it paid dividends. We caught issues in pre-production that would have been catastrophic in production. We optimized configurations based on real data rather than guesswork. And we slashed our observability costs by more than half.

## Performance Trade-offs at Scale

The abstraction layer cost us performance—there's no way around it. Direct Kafka producers can achieve lower latency than going through DAPR's sidecar. We measured the impact carefully:

| Performance Metric | Before DAPR | After DAPR | Delta |
|--------------------|-------------|------------|-------|
| **P50 Latency** | 8ms | 11ms | +3ms (37% increase) |
| **P99 Latency** | 45ms | 52ms | +7ms (15% increase) |
| **Max Throughput** | 520K TPS | 550K TPS | +30K TPS (5% increase) |
| **CPU per Pod** | 0.5 cores | 0.8 cores | +0.3 cores (60% increase) |
| **Memory per Pod** | 512MB | 768MB | +256MB (50% increase) |
| **Cluster Cost** | Baseline | +18% | Significant but acceptable |

But here's what we gained: operational simplicity. Our developers stopped worrying about client library versions, connection pooling, and serialization formats. They called an HTTP endpoint. DAPR handled the rest.

We also gained flexibility. When Pulsar had issues, we could route specific topics back to Kafka without code changes. When we needed to implement complex routing rules, we did it in DAPR configuration rather than scattered across codebases.

The resource overhead was significant—an extra sidecar per pod, each consuming memory and CPU. Our cluster costs increased by roughly 18%. But developer velocity increased dramatically, and operational incidents decreased. The ROI was clear.

Interestingly, maximum throughput actually improved slightly. DAPR's built-in connection pooling and message batching optimizations helped certain high-volume producers achieve better performance than their hand-rolled Kafka client code.

## The Istio Security Dividend

Strict mTLS with automated ambient mesh enrollment meant every message passing through our system was encrypted and authenticated. We didn't have to implement custom authentication logic. We didn't have to rotate API keys or manage service credentials manually.

The trade-off? Complexity. Debugging network issues became harder. Understanding traffic flow required deep knowledge of both DAPR and Istio. Our observability stack had to evolve to handle distributed tracing across multiple proxy layers.

But for a system handling sensitive transactions at massive scale, the security posture was non-negotiable. We passed compliance audits that would have been nightmares without mTLS everywhere.

## Lessons from the Trenches

**Abstractions have costs.** DAPR added latency and resource overhead. We accepted this trade-off for portability and velocity, but teams must understand what they're giving up.

**Migration strategies matter more than technology.** DAPR didn't make the migration easy—it made it possible. The real work was planning, communication, and incremental rollout.

**Observability is everything.** With multiple layers of abstraction, understanding system behavior required massive investment in tracing, metrics, and logging. We couldn't debug by reading code anymore—we had to read distributed traces. The upfront cost of building comprehensive observability was high, but it became our most valuable asset during the migration.

**Automation saves lives.** Manually configuring Istio and DAPR for 200+ services would have been impossible. GitOps and infrastructure-as-code weren't optional—they were survival tools.

**Performance testing is non-negotiable.** We load-tested every migration phase in staging environments that mirrored production. Without this, we would have discovered our capacity limits at the worst possible time.

## Was It Worth It?

Six months post-migration, we're running entirely on Pulsar with DAPR as our pub-sub abstraction layer, handling 550K+ TPS during peak loads. Developer satisfaction is up because they're not wrestling with client libraries. Operational burden is down because we have one messaging system to maintain instead of two.

The trade-offs were real—performance overhead, complexity, learning curve, significant observability investment. But the strategic value of decoupling our application layer from our messaging infrastructure turned out to be the best architectural decision we made all year.

Our observability stack now gives us visibility that was impossible with direct Kafka clients. When issues occur, we diagnose them in minutes instead of hours. Our on-call engineers actually sleep through the night.

Sometimes the right abstraction at the right time is worth its weight in latency milliseconds. And sometimes, investing heavily in observability is what separates successful migrations from catastrophic failures.

---

*Considering a similar migration? Think hard about your trade-offs and observability strategy before diving in.*
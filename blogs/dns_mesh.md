# Cross-Region Multi-Account AWS DNS Mesh

## What Do You Mean by DNS Mesh? And Why Should We Care?

DNS mesh is a term I coined—it's not part of mainstream infrastructure literature as far as I know. But it solves a real problem: **centralized DNS resolution and management in distributed systems**, whether you're spanning multiple AWS accounts, regions, or both.

This article focuses heavily on AWS and how to implement the "mesh" using Route53 services. Here's what you get:

- **Centralized DNS management** across your entire AWS footprint
- **Unified auditing and security controls** for all DNS operations
- **Infrastructure as Code everywhere**—no manual configuration allowed
- **Centralized flow logs and resolver query logs** for compliance and debugging
- **Universal access**—any AWS account in any region can resolve private hosted zones tied to VPCs in your centralized account

If you're running a multi-account AWS organization and DNS management feels like herding cats, keep reading.

## Constraints and Requirements

When we designed this, we had some non-negotiable requirements:

- **Single point of DNS management** across the entire organization
- **Centralized logging** for DNS queries and VPC flow logs
- **Centralized DDoS protection** and security audit controls
- **Handle 10K QPS per resolver endpoint per region** without breaking a sweat
- **Consumer accounts create records** in the centralized account without jumping through hoops
- **100% Terraform-managed**—zero manual components tolerated
- **Built-in alerting and monitoring** to catch issues before they become incidents
- **Two degrees of redundancy**—Availability Zone level and region level

## The Solution

![architecture design diagram](assets/dns_mesh.svg)

## Let's Talk Implementation

The concept is elegant and relies on AWS Route53's underlying magic that lets you manage private hosted zones without explicit DNS delegations.

### Route53 Delegation Magic

When you create a private hosted zone (PHZ), these name servers are assigned:

- ns-0.awsdns-00.com
- ns-512.awsdns-00.net
- ns-1024.awsdns-00.org
- ns-1536.awsdns-00.co.uk

These name servers exist because DNS protocol requires every hosted zone to have an NS record set. They're reserved and never used by Route53 public hosted zones. You can only query private zones via Route53 Resolver in a VPC that's associated with the hosted zone through an inbound endpoint.

Here's the clever part: **all private hosted zones share the same nameservers**. Setting up delegation would be pointless—you'd be delegating to servers your zone already uses. Effectively, delegation is already configured the moment you create a private zone.

### The Complexity of Private Hosted Zones

Here's the catch: private hosted zones need at least one VPC to exist. Even though Route53 is a global service, your PHZs are essentially region-bound at creation time.

Let's understand this with an example using two PHZs:

- `us.xyz.com`
- `eu.xyz.com`

Using common sense, you might host `us.xyz.com` in an Oregon VPC (`us`) and `eu.xyz.com` in an eu-west-1 VPC (`eu`). Good news: these VPCs can be bare-bones—no security groups or complex firewall rules needed initially. We'll tackle networking complexity later.

Let's focus on VPC `us` and PHZ `us.xyz.com` as our reference implementation. The setup for `eu` would be identical.

### Handling PHZ I/O

Working with VPC `us` and PHZ `us.xyz.com`, let's solve the input side first.

**Requirements:**
- 2 subnets in the `us` VPC (different AZs for redundancy)
- Resolver inbound endpoint
- Resolver rule

**Setup:**
- Create a resolver inbound endpoint using IPs from both subnets
- Create a resolver rule for domain `us.xyz.com` pointing to the inbound IP addresses
- Share the resolver rule with the entire AWS organization

That's it for inbound traffic.

**For outbound traffic:**
- Create a resolver outbound endpoint using 2 IPs from the same subnets
- Attach the outbound endpoint to the PHZ

Done. With proper security group configuration (or wide-open settings for testing), this PHZ is now ready for bidirectional operations.

Repeat these steps for all PHZs. Remember: **you only need one VPC per region**. The same VPC, subnets, and endpoints can serve all PHZs hosted in that region.

### Making the Mesh

Now for the interesting question: **"What if my workload in US needs to access DNS records in EU?"**

This hits several roadblocks:

- **Resolver rules are region-specific**—how do you create a resolver for an EU domain in US when regions have different VPCs and endpoint IPs?
- **Endpoints are region-specific**—do you replicate IPs across regions or enable cross-region IP communication?
- **VPCs are region-specific**—and since PHZs depend on VPCs, we've inherited this constraint

*This sounds like an IP mess... the solution is a mesh*

Wait—PHZs inherit VPC regional constraints, right? Wrong. **PHZs are still global Route53 services.**

Here's the solution: **associate all your PHZs with DNS VPCs from all regions**. If you have 15 regions, associate all 15 VPCs (that host resolver endpoints) with every single PHZ. This enables communication between each PHZ and each region.

*But how do I reach these VPCs cross-region? Don't I need region-specific resolver rules?*

Exactly. Create resolver rules for **all PHZs in all regions**—not just the home region. You need inbound and outbound IPs in all regions, which you already have from the earlier setup. Reuse them.

Because all DNS VPCs are associated with all PHZs, once you're in any region's VPC, you can access all PHZs. It doesn't matter if you hit EU or US IPs—you get access to everything. In our implementation, this meant 12 regions and roughly 50 PHZs, all interconnected.

### What About Security?

The security model is surprisingly simple because everything lives in one centralized account:

**Requirements:**
- 1 S3 bucket
- Security groups on all resolver endpoints
- AWS Shield

**DNS Query Logs:** Enable query logging in all host VPCs targeting `S3_bucket_name/dnsQueryLogs`

**VPC Flow Logs:** Target the same bucket at `S3_bucket_name/vpcFlowLogs`

AWS automatically organizes logs by region, VPC ID, and datetime. Stick with the default structure unless you have specific audit requirements.

**Security Groups:** Keep it simple—allow egress on all ports and IP ranges. Restrict ingress to your VPN IP range and Cloud WAN (CWAN) IP range if applicable.

**AWS Shield:** The basic plan should suffice, assuming you have network-level firewalls intercepting internet traffic before it reaches your internal network and queries these private zones.

## The Result

We now have a fully automated, centralized DNS infrastructure spanning multiple AWS accounts and regions. Services in any account can create records in the centralized account. DNS resolution works seamlessly across regions. All logs flow to a single S3 bucket for audit and compliance.

The mesh handles 10K+ QPS per endpoint without breaking stride, provides AZ and region-level redundancy, and everything is defined in Terraform.

Sometimes the best solution is the one that turns a mess into a mesh.

---

*This design was inspired by [AWS's blog on simplifying DNS management in multi-account environments](https://aws.amazon.com/blogs/security/simplify-dns-management-in-a-multiaccount-environment-with-route-53-resolver/).*
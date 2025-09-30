# Cross-region Multi-account AWS DNS Mesh

## What do you mean by DNS mesh ? and why should we bother with it ?

- DNS mesh is a term I coined, its not part of any popular literature as far as I know
- DNS mesh enables centralized DNS resolution and management in any distributed system (multi-account or multi-region or both)
- This article is heavily focused on AWS and how to implement the "mesh" in AWS using route53 services
- The mesh enables centralized auditing and security controls pertaining to DNS
- The mesh enables easier and centralized IAC for all DNS components in an architecture
- The mesh would further allow for centralized flow-logs and DNS resolver queries
- The mesh can receive and respond to resolution requests from any AWS account and region to private hosted zones tied to any AWS VPC in any region within the centralised AWS account

## Constraints and requirements

- Should be a single point of DNS management across the firm
- Should have centralized flow-logs and DNS resolution query logs
- Should have centralized DDOS protection and security audit controls
- Should be able to handle 10K QPS per resolver endpoint per region
- Services running in any consumer AWS account must have the ability to create records in the centralized account
- Should be hosted using Terraform ... NO manual components at all
- Should have alerting and monitoring systems baked in to prevent unwarranted failures
- Should have 2 degress of redundancy (AZ level and region level)

## The solution

![architecture design diagram](assets/dns_mesh.svg)

## Lets talk implementation

The idea is quite simple and relies on the underlying AWS magic that allows you to manage private hosted zones without setting up explicit DNS delegations.

### _route53 delegation magic_

When you create a private hosted zone, the following name servers are used: -

- ns-0.awsdns-00.com
- ns-512.awsdns-00.net
- ns-1024.awsdns-00.org
- ns-1536.awsdns-00.co.uk

These name servers are used because the DNS protocol requires that every hosted zone must have an NS record set. These name servers are reserved and never used by Route 53 public hosted zones. You can only query those zones via Route 53 Resolver in a VPC that has been associated to the hosted zone by using an inbound endpoint connected to the VPCs specified in the private hosted zone

Because all the private hosted zones have the same set of nameservers, setting up delegation would be pointless, you'd be delegating to the same set of servers across all zones. On top of this, your current zone already has these entries running as active servers, effectively, delegation is already set-up whenever you create a new private zone.

### _Complexity of private hosted zones_

Well, here's the biggest problem - private hosted zones need atleast one VPC to host themselves. This means, even though Route53 is a global service, your PHZs are essentially bound to a region at the time of creation. Its best to understand this inter-connectivity using an example with 2 PHZs: -

- us.xyz.com
- eu.xyz.com

Now, lets imagine that you use common-sense and determine that the Oregon region would make a good VPC host (VPC named `us`) for PHZ `us.xyz.com` and `eu-west-1` would make a good host VPC (VPC named `eu`) for PHZ `eu.xyz.com`.
Just so you know, you don't have to setup any seurity groups and firewalls while creating these VPCs ... a bare-bones setup is good enough ... we will dive into networking complexities later on.
Now, lets focus on VPC `us` and PHZ `us.xyz.com` as our benchmark example, the setup for `eu` would be isometric.

### _Handling PHZ I/O_
A reminder that we are working with VPC `us` and PHZ `us.xyz.com`
Lets solve the input side of this equation first, here's the list of items required: -

- 2 subnets in `us` VPC (preferrably different AZs if you want AZ level redundancy)
- Resolver inbound endpoint
- Resolver rule

Here's the simplest explanation around how to wire all of this together: -

- Create a resolver inbound endpoint using IPs from the 2 subnets
- Create a resolver rule for domain `us.xyz.com` with the 2 inbound IP addresses as targets for these resolver rules
- Share the resolver rule with the entire organization

And ... you are done ... as far as the input is concerned

Now lets solve the output side of this equation, here's the list of items required: -

- Resolver outbound endpoint
- The same 2 subnets you created above

And here comes the simplest explanation of outbound setup possible: -

- Create a resolver outbound endpoint using 2 IPs from the same subnets you created enough
- Attach the resolver outbound endpoint to the same PHZ

You've completed the setup for VPC `us` and PHZ `us.xyz.com`. Given that your security group configurations are either wide open or properly configured, this PHZ is now available for both inbound and outbound operations

Now repeat the same steps for all PHZs ... remember that you only need one VPC per regoin to do this ... the same VPC, subnets, inbound and outbound endpoints can be used for all PHZs hosted out of the same region

### _Making the mesh_ ###
Now that you know how to create region scpecific PHZs, the natural evolution is to ask a very simple question i.e. "what if my workload is in US but wants to access the DNS records of EU ?" ... this is an interesting question for multiple reasons because it hits the following roadblocks: -

- Resolver rules are region specific : how do you create a resolver for EU domain in US, because both regions have different VPCs and endpoint IPs ?
- Resolver inbound and outbound endpoints are region specific : either you replicate the IPs in resolver rules from different regions or somehow allow the IPs to communicate cross-region
- VPCs are region specific, because our PHZs are hosted on top of VPCs - we have inheritted this constraint

_This suddenly sounds complicated ... an IP mess_
_solution is mess --> mesh_

PHZs must have inheritted VPC scoped regional constraints, this is not true ... PHZs are still very a much a global Route53 service

The idea is to associate all of your PHZs with DNS VPC from all regions, this is copletely A-ok ... so if you have 15 regions, go ahead and associate all 15 VPCs (from which you created resolver inbound/outbound) to all of your PHZs ... this solves communication between each PHZ and each region

_But how do I reach these VPCs cross-region ... wouldn't that require a region specific reoslver rule ??_
You are bang-on-the-money ... the idea is to create resolver rules for all of your PHZs in all regions, not just the one region where host VPC resides, but ALL regions. To do this, you need inbound and outbound IPs in all regions ... does that ring a bell ? ... you already have these endpoints in all regions ... simply reuse them.
Because all DNS VPCs are associated to all PHZs, once you are in any region's VPC you will have access to all PHZs ... so it doesn't matter if you hit IPs from EU or US ... you'll end up having the best of both worlds ... in my case, this was 12 regions and roughly 50 PHZs

### What about Security ?###

Here's all you need: -

- 1 S3 bucket
- Security groups on all i/o endpoints
- AWS shield

Yeah, thats all you need ... because all of your DNS infrastructure is in one account

- DNS Query logs : Turn on DNS query logging in all host VPCs to target S3_bucket_name/dnsQueryLogs
- VPC flowlogs : Target the same bucket S3_bucket_name/vpcFlowLogs

AWS is pretty smart, it has a default log structure to sort the events based on distinct regoins, VPC IDs and date-time files ... I'd recommend sticking to this default format unless you have some weird audit controls ... ofcourse you can modify them if you feel like it

- Security groups : well, this can get complex, but the idea is to allow egress on all ports and IP ranges ... limit the ingress to your VPN IP range and CWAN IP range if you have one
- AWS shield : basic plan should be good enough ... assuming you have network level firewalls in place to intercept traffic from the internet much before it enters your internal network and starts querying these private zones

[Here's the AWS document which inspired this mesh design](https://aws.amazon.com/blogs/security/simplify-dns-management-in-a-multiaccount-environment-with-route-53-resolver/) 
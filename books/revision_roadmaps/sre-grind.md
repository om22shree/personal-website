# SRE grind-set

## 1. Release gates — interview ready when

- [ ] **OA:** solve 2 medium Python questions in a 75-minute mock in ≥4/5 recent attempts; complexity + edge cases explained.
- [ ] **Python:** build and test 8 of the 12 operational exercises below; remaining four stay on the backlog. At least two under live 60-minute time limits.
- [ ] **Terraform:** complete T1–T4; predict plans, provider scope, and state effects without hand-waving.
- [ ] **Troubleshooting:** present 3 evidence-led incidents: symptom → hypothesis → metric/command → isolation → mitigation → prevention.
- [ ] **HLD/LLD:** pass 3 timed Python LLDs and 4 platform HLD whiteboards; cover SLO, failures, scaling, security, cost.
- [ ] **Production ownership:** prepare 6 honest STAR stories with YOUR actions, metrics, trade-offs, and failures.

## 2. DSA — OA only

**Priority order:** arrays/hash → sliding windows/prefix → binary search → sorting/intervals → stacks/queues/heaps → trees/graphs → greedy/basic DP → design.  
**How to solve:** Python; first independent solution; identify invariant; state time/space; edge tests; repeat at +1 and +7 days.  
**Mandatory bank:** the **161** lines tagged `MUST` in Appendix A. The remaining questions are extension work, not the admission ticket.

**Most platform-adjacent MUST items:** `#135` Time Based Key-Value Store, `#209` Design Hit Counter, `#304` Task Scheduler, `#367–368` Course Schedule, `#506` LRU Cache, `#507` LFU Cache, `#511` Logger Rate Limiter, `#512` In-Memory File System, `#519` Authentication Manager.

## 3. Python engineering — know it cold

- [ ] **Core:** `dict/set/deque/heapq`, iterators/generators, dataclasses, protocols, exceptions, typing, datetime/timezones, pathlib, JSON/YAML/CSV/regex, `ipaddress`.
- [ ] **OS:** `subprocess` without shell injection, process groups, timeouts, signals, tempfiles, atomic writes, file permissions, exit codes.
- [ ] **Concurrency:** `asyncio` / `TaskGroup`, semaphore, bounded queues, cancellation, deadlines, backpressure, graceful shutdown. Know threads vs processes vs async and GIL constraints.
- [ ] **HTTP/API:** FastAPI, Pydantic, `httpx`, pagination, auth, connection pools, idempotency, 429/retry-after, bounded exponential backoff + jitter.
- [ ] **Ops:** `boto3` pagination/assume-role, Kubernetes Python client with least-privilege RBAC; structured logs, metrics, cardinality, SLOs; no hardcoded secrets.
- [ ] **Quality:** `uv sync/run`, lockfile, `pytest` fixtures/mocks/parametrization, Ruff, CI, deterministic failures and tests.

### Mandatory hands-on Python exercises

**Definition of done for each:** CLI/API, README, sample sanitized fixtures, tests including one injected failure, bounded resource usage, useful errors, dry-run for changes.

- [ ] **P01 — Streaming logs:** multi-GB JSONL, bounded memory, time-window 5xx counts, malformed/out-of-order lines.
- [ ] **P02 — Async health checker:** concurrency bounds, timeout, pooling, jitter, cancel/SIGINT, JSON + Prometheus output.
- [ ] **P03 — AWS auditor:** `boto3` paginator, multi-region inventory, unattached resources/tag drift, assume-role and read-only IAM.
- [ ] **P04 — K8s triage:** collect events, pod state, restart reason, previous logs; distinguish OOM vs probe/config failure.
- [ ] **P05 — Config drift engine:** semantic YAML/JSON diff, secret redaction, deterministic output, safe plan-only remediation.
- [ ] **P06 — TTL/LRU cache:** bounds, injected clock, expiration, concurrency policy, deterministic tests.
- [ ] **P07 — Rate limiter:** deque/sliding window, per-key expiry, burst/boundary tests, distributed Redis trade-off.
- [ ] **P08 — Process watchdog:** capture/stream stdout + stderr safely, timeout, kill process group, propagate exit codes.
- [ ] **P09 — Bulk API reconciler:** plan/apply, idempotency, retries/429, checkpoint, resume and audit trail.
- [ ] **P10 — SLO calculator:** multiwindow burn rates, rolling windows, missing data semantics.
- [ ] **P11 — DNS/TLS probe:** separate DNS/connect/TLS/HTTP latency, errors and certificate expiry.
- [ ] **P12 — Deployment policy scanner:** detect security, probe, limits and rollout issues; minimize false positives.

## 4. Terraform — whiteboard and terminal

- [ ] **T1 Aliases:** two AWS provider aliases, cross-account roles, explicit child-module `providers` mapping. Explain provider identity and state.
- [ ] **T2 Logic:** typed variables/validation, `condition ? a : b`, `for` expressions, `for_each`, `count`, dynamic blocks. Demonstrate index churn vs stable keys.
- [ ] **T3 Refactor:** modules, remote state/locking, `moved` blocks, import, safe non-destructive plan.
- [ ] **T4 Drift:** detect and resolve out-of-band edits without surprise replacement; CI validate/plan, IAM and secrets.

**Must explain:** dependency graph, plan vs apply, state risks, workspace limits, destructive diff, rollback/recovery, `sensitive` not protecting state data.

## 5. SRE troubleshooting — use existing arsenal

**Existing baseline:** `Auditor_commands.sh` already covers `ss`, `tcpdump`, `ethtool`, `pidstat`, `/proc`, `dmesg`, `iostat`, `strace`, `perf`, `nsenter`, `crictl`. No memorization laps.

- [ ] **Linux:** file descriptors, zombie/process signals, OOM/cgroups v2, PSI, CPU throttling, inode exhaustion, filesystem stalls.
- [ ] **Network:** TCP retransmits/conntrack/ephemeral ports, DNS cache/resolution, NAT/SNAT, TLS/SNI, MTU, routing and packet drops.
- [ ] **K8s advanced:** CNI/service path, CoreDNS failures, IP exhaustion, HPA/scheduler/eviction, stuck PVC, controller reconciliation, etcd/control-plane risks, upgrades, RBAC/admission.
- [ ] **Observability:** RED/USE, histograms vs summaries, cardinality, trace correlation, alert burn rates, missing telemetry.
- [ ] **Incident drill A:** high p99, low CPU: separate app waits, queueing, downstream, throttling, network, storage.
- [ ] **Incident drill B:** CrashLoop/OOM/restarts: identify cause before increasing limits/restarting.
- [ ] **Incident drill C:** Kafka lag or DNS timeouts: evidence-based narrowing and mitigation.

**Operator writing = optional stretch:** Python `kopf` (community framework) + official Kubernetes Python client; desired/observed reconciliation, retries, finalizers, idempotency, duplicate events. **Do not pretend deploying CRDs equals authoring controllers.**

## 6. AWS + Bash — what interviewers can probe

- [ ] **AWS:** VPC routing, IGW/NAT/endpoints, SG/NACL, IAM trust vs permissions, STS, KMS, ALB/NLB, ASG, EKS/VPC CNI, EBS/EFS/S3, RDS, multi-AZ failover, CloudWatch, capacity + cost.
- [ ] **Bash:** quoting/word splitting, arrays, `find -print0`, `xargs -0`, `jq`, pipes/exit statuses, traps, `set -euo pipefail` pitfalls, safe cleanup, idempotent scripts.
- [ ] **On-call:** SLI/SLO/error budgets, runbooks, incident command, escalation, postmortems, rollout/rollback, toil reduction. Ask employers about 24h coverage, staffing and compensatory recovery.

## 7. Python LLD and platform HLD

**LLD — timed implementation + tests:**

- [ ] L1 TTL/LRU cache: expiration, eviction, concurrency.
- [ ] L2 rate limiter: sliding window, bounded retention, Redis-backed variant.
- [ ] L3 task scheduler: backoff, deadlines, cancellation, graceful drain.
- [ ] L4 log aggregator: ingest, filters, pagination, storage bounds.

**HLD — 45-minute whiteboards:**

- [ ] H1 internal developer deployment platform: tenant isolation, paved roads, rollback.
- [ ] H2 highly available multi-AZ Kubernetes application: SLO, capacity, failure boundaries.
- [ ] H3 metrics/logs/traces pipeline: ingestion, sampling, cardinality, retention.
- [ ] H4 distributed rate limiter: hot keys, Redis failure, correctness trade-offs.
- [ ] H5 secure CI/CD: provenance, promotion, canary, rollback.
- [ ] H6 async job platform: idempotency, queue, DLQ, backpressure, retries.

**HLD answer skeleton:** requirements → rough numbers → APIs/data flow → failure domains → security → scaling → SLO/observability → cost → trade-offs. If you can't explain a failure, you don't own the design.

## 8. Build only what you can defend

- [ ] **C1 FastAPI platform control API:** authenticated deployment intent, idempotent workflow, async worker, fake K8s adapter, metrics/traces, failure tests.
- [ ] **C2 Python incident investigator:** read-only K8s client, dedupe/classify incidents, bounded collection, deterministic report, `kind` integration.
- [ ] **C3 Terraform AWS blueprint:** aliases, multi-account IAM, networking, remote-state design, CI plan, cost notes, no careless provisioning.

**Every repo:** `uv` lockfile if Python, reproducible run/test commands, CI, Mermaid diagram, threat/failure model, README and one tested recovery path. Prioritize **one finished, defensible repo** over three half-built ones.

---

# APPENDIX A — FULL 520-PROBLEM BANK

`[x]` preserves original solved status. **MUST** denotes the 161 selected OA targets. `F/B/T` are curriculum stages, **not** LeetCode difficulty. All original links retained.

### 01 — True Fundamentals: Traversal, Indexing & Simple State

- [x] 001 [Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/) · F · **MUST**

- [x] 002 [Richest Customer Wealth](https://leetcode.com/problems/richest-customer-wealth/) · F

- [x] 003 [Build Array from Permutation](https://leetcode.com/problems/build-array-from-permutation/) · F

- [x] 004 [Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/) · F

- [x] 005 [Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/) · F

- [x] 006 [Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/) · F

- [x] 007 [Final Value of Variable After Performing Operations](https://leetcode.com/problems/final-value-of-variable-after-performing-operations/) · B

- [x] 008 [Maximum Number of Words Found in Sentences](https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/) · B

- [x] 009 [Plus One](https://leetcode.com/problems/plus-one/) · F · **MUST**

### 02 — Hash Maps, Sets & Frequency Tables

- [ ] 010 [Majority Element](https://leetcode.com/problems/majority-element/) · B

- [ ] 011 [Two Sum](https://leetcode.com/problems/two-sum/) · F · **MUST**

- [ ] 012 [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) · F · **MUST**

- [ ] 013 [Valid Anagram](https://leetcode.com/problems/valid-anagram/) · F · **MUST**

- [ ] 014 [Ransom Note](https://leetcode.com/problems/ransom-note/) · F

- [ ] 015 [First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/) · F · **MUST**

- [ ] 016 [Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/) · F

- [ ] 017 [Word Pattern](https://leetcode.com/problems/word-pattern/) · B

- [ ] 018 [Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/) · B

- [ ] 019 [Happy Number](https://leetcode.com/problems/happy-number/) · B

- [ ] 020 [Group Anagrams](https://leetcode.com/problems/group-anagrams/) · B · **MUST**

- [ ] 021 [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) · B · **MUST**

- [ ] 022 [Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/) · B · **MUST**

- [ ] 023 [4Sum II](https://leetcode.com/problems/4sum-ii/) · B

- [ ] 024 [Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/) · B

- [ ] 025 [Determine if Two Strings Are Close](https://leetcode.com/problems/determine-if-two-strings-are-close/) · T

- [ ] 026 [Unique Number of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/) · T

- [ ] 027 [Find Common Characters](https://leetcode.com/problems/find-common-characters/) · T

- [ ] 028 [Jewels and Stones](https://leetcode.com/problems/jewels-and-stones/) · T

- [ ] 029 [Find the Difference](https://leetcode.com/problems/find-the-difference/) · T

- [ ] 030 [Longest Palindrome](https://leetcode.com/problems/longest-palindrome/) · T

### 03 — Prefix Sums, Prefix/Suffix State & Difference Arrays

- [ ] 031 [Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) · B · **MUST**

- [ ] 032 [Find Pivot Index](https://leetcode.com/problems/find-pivot-index/) · F · **MUST**

- [ ] 033 [Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/) · F · **MUST**

- [ ] 034 [Left and Right Sum Differences](https://leetcode.com/problems/left-and-right-sum-differences/) · F

- [ ] 035 [Find the Highest Altitude](https://leetcode.com/problems/find-the-highest-altitude/) · F

- [ ] 036 [Minimum Value to Get Positive Step by Step Sum](https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/) · F

- [ ] 037 [Continuous Subarray Sum](https://leetcode.com/problems/continuous-subarray-sum/) · F · **MUST**

- [ ] 038 [Contiguous Array](https://leetcode.com/problems/contiguous-array/) · B · **MUST**

- [ ] 039 [Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum/) · B · **MUST**

- [ ] 040 [Corporate Flight Bookings](https://leetcode.com/problems/corporate-flight-bookings/) · B · **MUST**

- [ ] 041 [Range Addition](https://leetcode.com/problems/range-addition/) · B

- [ ] 042 [Matrix Block Sum](https://leetcode.com/problems/matrix-block-sum/) · B

- [ ] 043 [Range Sum Query 2D - Immutable](https://leetcode.com/problems/range-sum-query-2d-immutable/) · B

- [ ] 044 [Number of Ways to Split Array](https://leetcode.com/problems/number-of-ways-to-split-array/) · B

- [ ] 045 [Sum of Absolute Differences in a Sorted Array](https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/) · B

- [ ] 046 [Maximum Population Year](https://leetcode.com/problems/maximum-population-year/) · T

- [ ] 047 [Plates Between Candles](https://leetcode.com/problems/plates-between-candles/) · T

- [ ] 048 [Ways to Make a Fair Array](https://leetcode.com/problems/ways-to-make-a-fair-array/) · T

- [ ] 049 [K Radius Subarray Averages](https://leetcode.com/problems/k-radius-subarray-averages/) · T

- [ ] 050 [XOR Queries of a Subarray](https://leetcode.com/problems/xor-queries-of-a-subarray/) · T

- [ ] 051 [Shifting Letters II](https://leetcode.com/problems/shifting-letters-ii/) · T

### 04 — Two Pointers

- [x] 052 [Move Zeroes](https://leetcode.com/problems/move-zeroes/) · F · **MUST**

- [x] 053 [Remove Element](https://leetcode.com/problems/remove-element/) · F

- [x] 054 [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) · F

- [x] 055 [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) · F · **MUST**

- [ ] 056 [Rotate Array](https://leetcode.com/problems/rotate-array/) · B · **MUST**

- [ ] 057 [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) · F · **MUST**

- [ ] 058 [Reverse String](https://leetcode.com/problems/reverse-string/) · F

- [ ] 059 [Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/) · F · **MUST**

- [ ] 060 [Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) · F · **MUST**

- [ ] 061 [Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/) · F

- [ ] 062 [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) · F · **MUST**

- [ ] 063 [3Sum](https://leetcode.com/problems/3sum/) · B · **MUST**

- [ ] 064 [3Sum Closest](https://leetcode.com/problems/3sum-closest/) · B

- [ ] 065 [4Sum](https://leetcode.com/problems/4sum/) · B

- [ ] 066 [Sort Colors](https://leetcode.com/problems/sort-colors/) · B · **MUST**

- [ ] 067 [Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/) · B

- [ ] 068 [Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/) · B

- [ ] 069 [String Compression](https://leetcode.com/problems/string-compression/) · B

- [ ] 070 [Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/) · B

- [ ] 071 [Append Characters to String to Make Subsequence](https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/) · T

- [ ] 072 [Minimum Length of String After Deleting Similar Ends](https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/) · T

- [ ] 073 [Boats to Save People](https://leetcode.com/problems/boats-to-save-people/) · T

- [ ] 074 [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) · T · **MUST**

- [ ] 075 [Minimum Number of Moves to Make Palindrome](https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/) · T

- [ ] 076 [Next Permutation](https://leetcode.com/problems/next-permutation/) · T · **MUST**

### 05 — Sliding Window

- [ ] 077 [Maximum Average Subarray I](https://leetcode.com/problems/maximum-average-subarray-i/) · F

- [ ] 078 [Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/) · F

- [ ] 079 [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) · F · **MUST**

- [ ] 080 [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/) · F · **MUST**

- [ ] 081 [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/) · F · **MUST**

- [ ] 082 [Permutation in String](https://leetcode.com/problems/permutation-in-string/) · F · **MUST**

- [ ] 083 [Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/) · B · **MUST**

- [ ] 084 [Fruit Into Baskets](https://leetcode.com/problems/fruit-into-baskets/) · B

- [ ] 085 [Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/) · B · **MUST**

- [ ] 086 [Grumpy Bookstore Owner](https://leetcode.com/problems/grumpy-bookstore-owner/) · B

- [ ] 087 [Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/) · B

- [ ] 088 [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) · B · **MUST**

- [ ] 089 [Longest Subarray of 1's After Deleting One Element](https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/) · B

- [ ] 090 [Get Equal Substrings Within Budget](https://leetcode.com/problems/get-equal-substrings-within-budget/) · T

- [ ] 091 [Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/) · T

- [ ] 092 [Count Number of Nice Subarrays](https://leetcode.com/problems/count-number-of-nice-subarrays/) · T

- [ ] 093 [Replace the Substring for Balanced String](https://leetcode.com/problems/replace-the-substring-for-balanced-string/) · T

- [ ] 094 [Maximum Points You Can Obtain from Cards](https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/) · T · **MUST**

- [ ] 095 [Take K of Each Character From Left and Right](https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/) · T

### 06 — Sorting, Buckets & Ordering

- [ ] 096 [Sort Array By Parity](https://leetcode.com/problems/sort-array-by-parity/) · F

- [ ] 097 [Sort Array by Increasing Frequency](https://leetcode.com/problems/sort-array-by-increasing-frequency/) · F

- [ ] 098 [Relative Sort Array](https://leetcode.com/problems/relative-sort-array/) · F · **MUST**

- [ ] 099 [Height Checker](https://leetcode.com/problems/height-checker/) · F

- [ ] 100 [Array Partition](https://leetcode.com/problems/array-partition/) · F

- [ ] 101 [Minimum Number of Moves to Seat Everyone](https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/) · F

- [ ] 102 [Minimum Difference Between Highest and Lowest of K Scores](https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/) · B

- [ ] 103 [Maximum Product Difference Between Two Pairs](https://leetcode.com/problems/maximum-product-difference-between-two-pairs/) · B

- [ ] 104 [Relative Ranks](https://leetcode.com/problems/relative-ranks/) · B

- [ ] 105 [Third Maximum Number](https://leetcode.com/problems/third-maximum-number/) · B

- [ ] 106 [Custom Sort String](https://leetcode.com/problems/custom-sort-string/) · B

- [ ] 107 [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/) · B

- [ ] 108 [Rank Transform of an Array](https://leetcode.com/problems/rank-transform-of-an-array/) · B

- [ ] 109 [Sort the People](https://leetcode.com/problems/sort-the-people/) · B

- [ ] 110 [H-Index](https://leetcode.com/problems/h-index/) · T · **MUST**

- [ ] 111 [Largest Number](https://leetcode.com/problems/largest-number/) · T

- [ ] 112 [Sort an Array](https://leetcode.com/problems/sort-an-array/) · T · **MUST**

- [ ] 113 [Wiggle Sort II](https://leetcode.com/problems/wiggle-sort-ii/) · T

- [ ] 114 [Pancake Sorting](https://leetcode.com/problems/pancake-sorting/) · T

- [ ] 115 [Maximum Gap](https://leetcode.com/problems/maximum-gap/) · T

### 07 — Binary Search

- [ ] 116 [Binary Search](https://leetcode.com/problems/binary-search/) · F · **MUST**

- [ ] 117 [Search Insert Position](https://leetcode.com/problems/search-insert-position/) · F · **MUST**

- [ ] 118 [Guess Number Higher or Lower](https://leetcode.com/problems/guess-number-higher-or-lower/) · F

- [ ] 119 [First Bad Version](https://leetcode.com/problems/first-bad-version/) · F · **MUST**

- [ ] 120 [Sqrt(x)](https://leetcode.com/problems/sqrtx/) · F

- [ ] 121 [Valid Perfect Square](https://leetcode.com/problems/valid-perfect-square/) · F

- [ ] 122 [Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) · B · **MUST**

- [ ] 123 [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) · B · **MUST**

- [ ] 124 [Search in Rotated Sorted Array II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/) · B

- [ ] 125 [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) · B · **MUST**

- [ ] 126 [Find Peak Element](https://leetcode.com/problems/find-peak-element/) · B · **MUST**

- [ ] 127 [Peak Index in a Mountain Array](https://leetcode.com/problems/peak-index-in-a-mountain-array/) · B

- [ ] 128 [Find K Closest Elements](https://leetcode.com/problems/find-k-closest-elements/) · B

- [ ] 129 [Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/) · B · **MUST**

- [ ] 130 [Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/) · T · **MUST**

- [ ] 131 [Capacity To Ship Packages Within D Days](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/) · T · **MUST**

- [ ] 132 [Minimum Speed to Arrive on Time](https://leetcode.com/problems/minimum-speed-to-arrive-on-time/) · T

- [ ] 133 [Split Array Largest Sum](https://leetcode.com/problems/split-array-largest-sum/) · T

- [ ] 134 [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) · T

- [ ] 135 [Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/) · T · **MUST**

### 08 — Intervals & Sweep-Line Thinking

- [ ] 136 [Summary Ranges](https://leetcode.com/problems/summary-ranges/) · F

- [ ] 137 [Merge Intervals](https://leetcode.com/problems/merge-intervals/) · F · **MUST**

- [ ] 138 [Insert Interval](https://leetcode.com/problems/insert-interval/) · F · **MUST**

- [ ] 139 [Meeting Rooms](https://leetcode.com/problems/meeting-rooms/) · F · **MUST**

- [ ] 140 [Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/) · F · **MUST**

- [ ] 141 [Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) · F · **MUST**

- [ ] 142 [Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/) · B · **MUST**

- [ ] 143 [Interval List Intersections](https://leetcode.com/problems/interval-list-intersections/) · B · **MUST**

- [ ] 144 [Remove Covered Intervals](https://leetcode.com/problems/remove-covered-intervals/) · B

- [ ] 145 [Meeting Scheduler](https://leetcode.com/problems/meeting-scheduler/) · B

- [ ] 146 [My Calendar I](https://leetcode.com/problems/my-calendar-i/) · B · **MUST**

- [ ] 147 [My Calendar II](https://leetcode.com/problems/my-calendar-ii/) · B · **MUST**

- [ ] 148 [Employee Free Time](https://leetcode.com/problems/employee-free-time/) · B

- [ ] 149 [Video Stitching](https://leetcode.com/problems/video-stitching/) · B

- [ ] 150 [Minimum Interval to Include Each Query](https://leetcode.com/problems/minimum-interval-to-include-each-query/) · T · **MUST**

- [ ] 151 [Data Stream as Disjoint Intervals](https://leetcode.com/problems/data-stream-as-disjoint-intervals/) · T

- [ ] 152 [Amount of New Area Painted Each Day](https://leetcode.com/problems/amount-of-new-area-painted-each-day/) · T

- [ ] 153 [Divide Intervals Into Minimum Number of Groups](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/) · T

- [ ] 154 [Points That Intersect With Cars](https://leetcode.com/problems/points-that-intersect-with-cars/) · T

- [ ] 155 [Maximum Number of Events That Can Be Attended](https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/) · T

### 09 — Stack Fundamentals & Parsing

- [ ] 156 [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) · F · **MUST**

- [ ] 157 [Baseball Game](https://leetcode.com/problems/baseball-game/) · F

- [ ] 158 [Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/) · F

- [ ] 159 [Make The String Great](https://leetcode.com/problems/make-the-string-great/) · F

- [ ] 160 [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/) · F

- [ ] 161 [Min Stack](https://leetcode.com/problems/min-stack/) · F · **MUST**

- [ ] 162 [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/) · B · **MUST**

- [ ] 163 [Simplify Path](https://leetcode.com/problems/simplify-path/) · B · **MUST**

- [ ] 164 [Score of Parentheses](https://leetcode.com/problems/score-of-parentheses/) · B

- [ ] 165 [Decode String](https://leetcode.com/problems/decode-string/) · B · **MUST**

- [ ] 166 [Asteroid Collision](https://leetcode.com/problems/asteroid-collision/) · B

- [ ] 167 [Validate Stack Sequences](https://leetcode.com/problems/validate-stack-sequences/) · B

- [ ] 168 [Removing Stars From a String](https://leetcode.com/problems/removing-stars-from-a-string/) · B

- [ ] 169 [Minimum Remove to Make Valid Parentheses](https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/) · B

- [ ] 170 [Basic Calculator II](https://leetcode.com/problems/basic-calculator-ii/) · T · **MUST**

- [ ] 171 [Basic Calculator](https://leetcode.com/problems/basic-calculator/) · T

- [ ] 172 [Exclusive Time of Functions](https://leetcode.com/problems/exclusive-time-of-functions/) · T · **MUST**

- [ ] 173 [Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/) · T

- [ ] 174 [Check If Word Is Valid After Substitutions](https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/) · T

- [ ] 175 [Parse Lisp Expression](https://leetcode.com/problems/parse-lisp-expression/) · T

### 10 — Monotonic Stack

- [ ] 176 [Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/) · F · **MUST**

- [ ] 177 [Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/) · F

- [ ] 178 [Daily Temperatures](https://leetcode.com/problems/daily-temperatures/) · F · **MUST**

- [ ] 179 [Final Prices With a Special Discount in a Shop](https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/) · F

- [ ] 180 [Online Stock Span](https://leetcode.com/problems/online-stock-span/) · F · **MUST**

- [ ] 181 [Next Greater Node In Linked List](https://leetcode.com/problems/next-greater-node-in-linked-list/) · F

- [ ] 182 [132 Pattern](https://leetcode.com/problems/132-pattern/) · B

- [ ] 183 [Remove K Digits](https://leetcode.com/problems/remove-k-digits/) · B

- [ ] 184 [Car Fleet](https://leetcode.com/problems/car-fleet/) · B

- [ ] 185 [Shortest Unsorted Continuous Subarray](https://leetcode.com/problems/shortest-unsorted-continuous-subarray/) · B

- [ ] 186 [Maximum Width Ramp](https://leetcode.com/problems/maximum-width-ramp/) · B

- [ ] 187 [Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/) · B · **MUST**

- [ ] 188 [Maximal Rectangle](https://leetcode.com/problems/maximal-rectangle/) · B

- [ ] 189 [Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/) · B

- [ ] 190 [Sum of Subarray Ranges](https://leetcode.com/problems/sum-of-subarray-ranges/) · T

- [ ] 191 [Number of Visible People in a Queue](https://leetcode.com/problems/number-of-visible-people-in-a-queue/) · T

- [ ] 192 [Steps to Make Array Non-decreasing](https://leetcode.com/problems/steps-to-make-array-non-decreasing/) · T

- [ ] 193 [Minimum Cost Tree From Leaf Values](https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/) · T

- [ ] 194 [Beautiful Towers II](https://leetcode.com/problems/beautiful-towers-ii/) · T

- [ ] 195 [Total Strength of Wizards](https://leetcode.com/problems/total-strength-of-wizards/) · T

### 11 — Queues, Deques & Streaming Windows

- [ ] 196 [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/) · B · **MUST**

- [ ] 197 [Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/) · F · **MUST**

- [ ] 198 [Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues/) · F

- [ ] 199 [Number of Recent Calls](https://leetcode.com/problems/number-of-recent-calls/) · F · **MUST**

- [ ] 200 [Design Circular Queue](https://leetcode.com/problems/design-circular-queue/) · F

- [ ] 201 [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/) · F · **MUST**

- [ ] 202 [Number of Students Unable to Eat Lunch](https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/) · F

- [ ] 203 [Time Needed to Buy Tickets](https://leetcode.com/problems/time-needed-to-buy-tickets/) · B

- [ ] 204 [Find the Winner of the Circular Game](https://leetcode.com/problems/find-the-winner-of-the-circular-game/) · B

- [ ] 205 [Dota2 Senate](https://leetcode.com/problems/dota2-senate/) · B

- [ ] 206 [Reveal Cards In Increasing Order](https://leetcode.com/problems/reveal-cards-in-increasing-order/) · B

- [ ] 207 [Design Front Middle Back Queue](https://leetcode.com/problems/design-front-middle-back-queue/) · B

- [ ] 208 [First Unique Number](https://leetcode.com/problems/first-unique-number/) · B · **MUST**

- [ ] 209 [Design Hit Counter](https://leetcode.com/problems/design-hit-counter/) · B · **MUST**

- [ ] 210 [Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit](https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/) · B · **MUST**

- [ ] 211 [Continuous Subarrays](https://leetcode.com/problems/continuous-subarrays/) · T

- [ ] 212 [Maximum Number of Robots Within Budget](https://leetcode.com/problems/maximum-number-of-robots-within-budget/) · T

- [ ] 213 [Shortest Subarray with Sum at Least K](https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/) · T · **MUST**

- [ ] 214 [Constrained Subsequence Sum](https://leetcode.com/problems/constrained-subsequence-sum/) · T

- [ ] 215 [Jump Game VI](https://leetcode.com/problems/jump-game-vi/) · T

- [ ] 216 [Max Value of Equation](https://leetcode.com/problems/max-value-of-equation/) · T

### 12 — Linked Lists

- [ ] 217 [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) · F

- [ ] 218 [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) · F · **MUST**

- [ ] 219 [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) · F · **MUST**

- [ ] 220 [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) · F · **MUST**

- [ ] 221 [Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/) · F

- [ ] 222 [Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/) · F

- [ ] 223 [Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/) · B

- [ ] 224 [Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/) · B

- [ ] 225 [Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/) · B

- [ ] 226 [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/) · B · **MUST**

- [ ] 227 [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) · B · **MUST**

- [ ] 228 [Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/) · B

- [ ] 229 [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) · B · **MUST**

- [ ] 230 [Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/) · B

- [ ] 231 [Partition List](https://leetcode.com/problems/partition-list/) · T

- [ ] 232 [Reorder List](https://leetcode.com/problems/reorder-list/) · T · **MUST**

- [ ] 233 [Sort List](https://leetcode.com/problems/sort-list/) · T

- [ ] 234 [Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/) · T · **MUST**

- [ ] 235 [Flatten a Multilevel Doubly Linked List](https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/) · T

- [ ] 236 [Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/) · T

### 13 — Recursion & Backtracking

- [ ] 237 [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/) · F

- [ ] 238 [Pow(x, n)](https://leetcode.com/problems/powx-n/) · F

- [ ] 239 [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) · F · **MUST**

- [ ] 240 [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) · F · **MUST**

- [ ] 241 [Subsets](https://leetcode.com/problems/subsets/) · F · **MUST**

- [ ] 242 [Subsets II](https://leetcode.com/problems/subsets-ii/) · F

- [ ] 243 [Permutations](https://leetcode.com/problems/permutations/) · B · **MUST**

- [ ] 244 [Permutations II](https://leetcode.com/problems/permutations-ii/) · B

- [ ] 245 [Combinations](https://leetcode.com/problems/combinations/) · B

- [ ] 246 [Combination Sum](https://leetcode.com/problems/combination-sum/) · B · **MUST**

- [ ] 247 [Combination Sum II](https://leetcode.com/problems/combination-sum-ii/) · B

- [ ] 248 [Combination Sum III](https://leetcode.com/problems/combination-sum-iii/) · B

- [ ] 249 [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/) · B

- [ ] 250 [Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/) · B

- [ ] 251 [Word Search](https://leetcode.com/problems/word-search/) · T

- [ ] 252 [Beautiful Arrangement](https://leetcode.com/problems/beautiful-arrangement/) · T

- [ ] 253 [Matchsticks to Square](https://leetcode.com/problems/matchsticks-to-square/) · T

- [ ] 254 [N-Queens](https://leetcode.com/problems/n-queens/) · T

- [ ] 255 [N-Queens II](https://leetcode.com/problems/n-queens-ii/) · T

- [ ] 256 [Sudoku Solver](https://leetcode.com/problems/sudoku-solver/) · T

### 14 — Binary Trees & BST Fundamentals

- [ ] 257 [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) · F · **MUST**

- [ ] 258 [Same Tree](https://leetcode.com/problems/same-tree/) · F

- [ ] 259 [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) · F

- [ ] 260 [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) · F

- [ ] 261 [Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/) · F

- [ ] 262 [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/) · F · **MUST**

- [ ] 263 [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/) · B · **MUST**

- [ ] 264 [Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/) · B

- [ ] 265 [Path Sum](https://leetcode.com/problems/path-sum/) · B · **MUST**

- [ ] 266 [Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/) · B

- [ ] 267 [Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/) · B

- [ ] 268 [Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/) · B

- [ ] 269 [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) · B · **MUST**

- [ ] 270 [Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/) · B

- [ ] 271 [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/) · T · **MUST**

- [ ] 272 [Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) · T

- [ ] 273 [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) · T · **MUST**

- [ ] 274 [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/) · T

- [ ] 275 [Average of Levels in Binary Tree](https://leetcode.com/problems/average-of-levels-in-binary-tree/) · T

- [ ] 276 [Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/) · T

### 15 — Advanced Trees, Paths & Reconstruction

- [ ] 277 [Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/) · F

- [ ] 278 [Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/) · F

- [ ] 279 [Path Sum II](https://leetcode.com/problems/path-sum-ii/) · F

- [ ] 280 [Path Sum III](https://leetcode.com/problems/path-sum-iii/) · F

- [ ] 281 [Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) · F · **MUST**

- [ ] 282 [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) · F

- [ ] 283 [Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) · B

- [ ] 284 [Flatten Binary Tree to Linked List](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/) · B

- [ ] 285 [Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/) · B

- [ ] 286 [Binary Search Tree Iterator](https://leetcode.com/problems/binary-search-tree-iterator/) · B

- [ ] 287 [Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/) · B

- [ ] 288 [Recover Binary Search Tree](https://leetcode.com/problems/recover-binary-search-tree/) · B

- [ ] 289 [House Robber III](https://leetcode.com/problems/house-robber-iii/) · B

- [ ] 290 [All Nodes Distance K in Binary Tree](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/) · B

- [ ] 291 [Smallest Subtree with all the Deepest Nodes](https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/) · T

- [ ] 292 [Vertical Order Traversal of a Binary Tree](https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/) · T

- [ ] 293 [Boundary of Binary Tree](https://leetcode.com/problems/boundary-of-binary-tree/) · T

- [ ] 294 [Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) · T · **MUST**

- [ ] 295 [Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) · T · **MUST**

- [ ] 296 [Binary Tree Cameras](https://leetcode.com/problems/binary-tree-cameras/) · T

### 16 — Heaps & Priority Queues

- [ ] 297 [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/) · F

- [ ] 298 [Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) · F · **MUST**

- [ ] 299 [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) · F · **MUST**

- [ ] 300 [K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin/) · F · **MUST**

- [ ] 301 [Top K Frequent Words](https://leetcode.com/problems/top-k-frequent-words/) · F

- [ ] 302 [Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/) · F

- [ ] 303 [Smallest Number in Infinite Set](https://leetcode.com/problems/smallest-number-in-infinite-set/) · B

- [ ] 304 [Task Scheduler](https://leetcode.com/problems/task-scheduler/) · B · **MUST**

- [ ] 305 [Reorganize String](https://leetcode.com/problems/reorganize-string/) · B

- [ ] 306 [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) · B · **MUST**

- [ ] 307 [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) · B · **MUST**

- [ ] 308 [Ugly Number II](https://leetcode.com/problems/ugly-number-ii/) · B

- [ ] 309 [Furthest Building You Can Reach](https://leetcode.com/problems/furthest-building-you-can-reach/) · B

- [ ] 310 [Single-Threaded CPU](https://leetcode.com/problems/single-threaded-cpu/) · B · **MUST**

- [ ] 311 [Meeting Rooms III](https://leetcode.com/problems/meeting-rooms-iii/) · T

- [ ] 312 [Total Cost to Hire K Workers](https://leetcode.com/problems/total-cost-to-hire-k-workers/) · T

- [ ] 313 [Maximum Subsequence Score](https://leetcode.com/problems/maximum-subsequence-score/) · T

- [ ] 314 [IPO](https://leetcode.com/problems/ipo/) · T

- [ ] 315 [Smallest Range Covering Elements from K Lists](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/) · T

- [ ] 316 [Design Twitter](https://leetcode.com/problems/design-twitter/) · T

### 17 — Greedy

- [x] 317 [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) · F · **MUST**

- [ ] 318 [Assign Cookies](https://leetcode.com/problems/assign-cookies/) · F

- [ ] 319 [Lemonade Change](https://leetcode.com/problems/lemonade-change/) · F

- [ ] 320 [Can Place Flowers](https://leetcode.com/problems/can-place-flowers/) · F

- [ ] 321 [Maximum Units on a Truck](https://leetcode.com/problems/maximum-units-on-a-truck/) · F

- [ ] 322 [Minimum Cost to Move Chips to The Same Position](https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/) · F

- [ ] 323 [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) · F

- [ ] 324 [Jump Game](https://leetcode.com/problems/jump-game/) · B · **MUST**

- [ ] 325 [Jump Game II](https://leetcode.com/problems/jump-game-ii/) · B · **MUST**

- [ ] 326 [Gas Station](https://leetcode.com/problems/gas-station/) · B · **MUST**

- [ ] 327 [Partition Labels](https://leetcode.com/problems/partition-labels/) · B

- [ ] 328 [Hand of Straights](https://leetcode.com/problems/hand-of-straights/) · B

- [ ] 329 [Merge Triplets to Form Target Triplet](https://leetcode.com/problems/merge-triplets-to-form-target-triplet/) · B

- [ ] 330 [Valid Parenthesis String](https://leetcode.com/problems/valid-parenthesis-string/) · B · **MUST**

- [ ] 331 [Bag of Tokens](https://leetcode.com/problems/bag-of-tokens/) · B

- [ ] 332 [Eliminate Maximum Number of Monsters](https://leetcode.com/problems/eliminate-maximum-number-of-monsters/) · T

- [ ] 333 [Two City Scheduling](https://leetcode.com/problems/two-city-scheduling/) · T

- [ ] 334 [Broken Calculator](https://leetcode.com/problems/broken-calculator/) · T

- [ ] 335 [Maximum Swap](https://leetcode.com/problems/maximum-swap/) · T

- [ ] 336 [Monotone Increasing Digits](https://leetcode.com/problems/monotone-increasing-digits/) · T

- [ ] 337 [Earliest Possible Day of Full Bloom](https://leetcode.com/problems/earliest-possible-day-of-full-bloom/) · T

### 18 — Graphs: BFS & DFS

- [ ] 338 [Flood Fill](https://leetcode.com/problems/flood-fill/) · F · **MUST**

- [ ] 339 [Island Perimeter](https://leetcode.com/problems/island-perimeter/) · F

- [ ] 340 [Find if Path Exists in Graph](https://leetcode.com/problems/find-if-path-exists-in-graph/) · F · **MUST**

- [ ] 341 [Keys and Rooms](https://leetcode.com/problems/keys-and-rooms/) · F

- [ ] 342 [Number of Provinces](https://leetcode.com/problems/number-of-provinces/) · F · **MUST**

- [ ] 343 [Number of Islands](https://leetcode.com/problems/number-of-islands/) · F · **MUST**

- [ ] 344 [Max Area of Island](https://leetcode.com/problems/max-area-of-island/) · B

- [ ] 345 [Surrounded Regions](https://leetcode.com/problems/surrounded-regions/) · B

- [ ] 346 [Clone Graph](https://leetcode.com/problems/clone-graph/) · B · **MUST**

- [ ] 347 [All Paths From Source to Target](https://leetcode.com/problems/all-paths-from-source-to-target/) · B

- [ ] 348 [Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/) · B

- [ ] 349 [Number of Enclaves](https://leetcode.com/problems/number-of-enclaves/) · B

- [ ] 350 [Count Sub Islands](https://leetcode.com/problems/count-sub-islands/) · B

- [ ] 351 [01 Matrix](https://leetcode.com/problems/01-matrix/) · B · **MUST**

- [ ] 352 [As Far from Land as Possible](https://leetcode.com/problems/as-far-from-land-as-possible/) · T

- [ ] 353 [Nearest Exit from Entrance in Maze](https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/) · T

- [ ] 354 [Shortest Path in Binary Matrix](https://leetcode.com/problems/shortest-path-in-binary-matrix/) · T · **MUST**

- [ ] 355 [Minimum Genetic Mutation](https://leetcode.com/problems/minimum-genetic-mutation/) · T

- [ ] 356 [Word Ladder](https://leetcode.com/problems/word-ladder/) · T · **MUST**

- [ ] 357 [Shortest Bridge](https://leetcode.com/problems/shortest-bridge/) · T

### 19 — Union-Find, DAGs & Topological Sort

- [ ] 358 [Redundant Connection](https://leetcode.com/problems/redundant-connection/) · F · **MUST**

- [ ] 359 [Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/) · F · **MUST**

- [ ] 360 [Number of Connected Components in an Undirected Graph](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/) · F · **MUST**

- [ ] 361 [Accounts Merge](https://leetcode.com/problems/accounts-merge/) · F

- [ ] 362 [Satisfiability of Equality Equations](https://leetcode.com/problems/satisfiability-of-equality-equations/) · F

- [ ] 363 [Most Stones Removed with Same Row or Column](https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/) · F

- [ ] 364 [Regions Cut By Slashes](https://leetcode.com/problems/regions-cut-by-slashes/) · B

- [ ] 365 [Smallest String With Swaps](https://leetcode.com/problems/smallest-string-with-swaps/) · B

- [ ] 366 [Similar String Groups](https://leetcode.com/problems/similar-string-groups/) · B

- [ ] 367 [Course Schedule](https://leetcode.com/problems/course-schedule/) · B · **MUST**

- [ ] 368 [Course Schedule II](https://leetcode.com/problems/course-schedule-ii/) · B · **MUST**

- [ ] 369 [Find Eventual Safe States](https://leetcode.com/problems/find-eventual-safe-states/) · B

- [ ] 370 [Minimum Height Trees](https://leetcode.com/problems/minimum-height-trees/) · B

- [ ] 371 [Alien Dictionary](https://leetcode.com/problems/alien-dictionary/) · B · **MUST**

- [ ] 372 [Parallel Courses](https://leetcode.com/problems/parallel-courses/) · T

- [ ] 373 [Sequence Reconstruction](https://leetcode.com/problems/sequence-reconstruction/) · T

- [ ] 374 [Find All Possible Recipes from Given Supplies](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/) · T · **MUST**

- [ ] 375 [Sort Items by Groups Respecting Dependencies](https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/) · T

- [ ] 376 [Largest Color Value in a Directed Graph](https://leetcode.com/problems/largest-color-value-in-a-directed-graph/) · T

### 20 — Shortest Paths, Weighted Graphs & MST

- [ ] 377 [Min Cost to Connect All Points](https://leetcode.com/problems/min-cost-to-connect-all-points/) · F · **MUST**

- [ ] 378 [Network Delay Time](https://leetcode.com/problems/network-delay-time/) · F · **MUST**

- [ ] 379 [Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/) · F · **MUST**

- [ ] 380 [Path With Minimum Effort](https://leetcode.com/problems/path-with-minimum-effort/) · F · **MUST**

- [ ] 381 [Swim in Rising Water](https://leetcode.com/problems/swim-in-rising-water/) · F

- [ ] 382 [Path with Maximum Probability](https://leetcode.com/problems/path-with-maximum-probability/) · F

- [ ] 383 [Minimum Cost to Make at Least One Valid Path in a Grid](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/) · F

- [ ] 384 [Minimum Obstacle Removal to Reach Corner](https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/) · B

- [ ] 385 [Reachable Nodes In Subdivided Graph](https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/) · B

- [ ] 386 [Number of Ways to Arrive at Destination](https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/) · B

- [ ] 387 [Find the City With the Smallest Number of Neighbors at a Threshold Distance](https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) · B

- [ ] 388 [Evaluate Division](https://leetcode.com/problems/evaluate-division/) · B

- [ ] 389 [The Maze II](https://leetcode.com/problems/the-maze-ii/) · B

- [ ] 390 [The Maze III](https://leetcode.com/problems/the-maze-iii/) · B

- [ ] 391 [Path With Maximum Minimum Value](https://leetcode.com/problems/path-with-maximum-minimum-value/) · B

- [ ] 392 [Minimum Cost to Reach City With Discounts](https://leetcode.com/problems/minimum-cost-to-reach-city-with-discounts/) · T

- [ ] 393 [Connecting Cities With Minimum Cost](https://leetcode.com/problems/connecting-cities-with-minimum-cost/) · T

- [ ] 394 [Optimize Water Distribution in a Village](https://leetcode.com/problems/optimize-water-distribution-in-a-village/) · T

- [ ] 395 [Critical Connections in a Network](https://leetcode.com/problems/critical-connections-in-a-network/) · T · **MUST**

- [ ] 396 [Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/) · T

- [ ] 397 [Shortest Path Visiting All Nodes](https://leetcode.com/problems/shortest-path-visiting-all-nodes/) · T

### 21 — Dynamic Programming I: 1D State

- [x] 398 [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) · F · **MUST**

- [ ] 399 [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) · F · **MUST**

- [ ] 400 [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/) · F

- [ ] 401 [N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/) · F

- [ ] 402 [House Robber](https://leetcode.com/problems/house-robber/) · F · **MUST**

- [ ] 403 [House Robber II](https://leetcode.com/problems/house-robber-ii/) · F

- [ ] 404 [Delete and Earn](https://leetcode.com/problems/delete-and-earn/) · F

- [ ] 405 [Perfect Squares](https://leetcode.com/problems/perfect-squares/) · B

- [ ] 406 [Decode Ways](https://leetcode.com/problems/decode-ways/) · B · **MUST**

- [ ] 407 [Word Break](https://leetcode.com/problems/word-break/) · B · **MUST**

- [ ] 408 [Integer Break](https://leetcode.com/problems/integer-break/) · B

- [ ] 409 [Coin Change](https://leetcode.com/problems/coin-change/) · B · **MUST**

- [ ] 410 [Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/) · B

- [ ] 411 [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/) · B · **MUST**

- [ ] 412 [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) · B · **MUST**

- [ ] 413 [Number of Longest Increasing Subsequence](https://leetcode.com/problems/number-of-longest-increasing-subsequence/) · T

- [ ] 414 [Wiggle Subsequence](https://leetcode.com/problems/wiggle-subsequence/) · T

- [ ] 415 [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) · T

- [ ] 416 [Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) · T

- [ ] 417 [Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/) · T

- [ ] 418 [Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/) · T · **MUST**

### 22 — Dynamic Programming II: Grids & Knapsack

- [ ] 419 [Unique Paths](https://leetcode.com/problems/unique-paths/) · F · **MUST**

- [ ] 420 [Unique Paths II](https://leetcode.com/problems/unique-paths-ii/) · F

- [ ] 421 [Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/) · F · **MUST**

- [ ] 422 [Triangle](https://leetcode.com/problems/triangle/) · F

- [ ] 423 [Minimum Falling Path Sum](https://leetcode.com/problems/minimum-falling-path-sum/) · F

- [ ] 424 [Minimum Falling Path Sum II](https://leetcode.com/problems/minimum-falling-path-sum-ii/) · F

- [ ] 425 [Dungeon Game](https://leetcode.com/problems/dungeon-game/) · B

- [ ] 426 [Maximal Square](https://leetcode.com/problems/maximal-square/) · B

- [ ] 427 [Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) · B

- [ ] 428 [Out of Boundary Paths](https://leetcode.com/problems/out-of-boundary-paths/) · B

- [ ] 429 [Knight Probability in Chessboard](https://leetcode.com/problems/knight-probability-in-chessboard/) · B

- [ ] 430 [Coin Change II](https://leetcode.com/problems/coin-change-ii/) · B

- [ ] 431 [Target Sum](https://leetcode.com/problems/target-sum/) · B

- [ ] 432 [Ones and Zeroes](https://leetcode.com/problems/ones-and-zeroes/) · B

- [ ] 433 [Last Stone Weight II](https://leetcode.com/problems/last-stone-weight-ii/) · T

- [ ] 434 [Profitable Schemes](https://leetcode.com/problems/profitable-schemes/) · T

- [ ] 435 [Paint House](https://leetcode.com/problems/paint-house/) · T

- [ ] 436 [Paint House II](https://leetcode.com/problems/paint-house-ii/) · T

- [ ] 437 [Cherry Pickup II](https://leetcode.com/problems/cherry-pickup-ii/) · T

- [ ] 438 [Cherry Pickup](https://leetcode.com/problems/cherry-pickup/) · T

### 23 — Dynamic Programming III: Strings, Sequences & Games

- [ ] 439 [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/) · F · **MUST**

- [ ] 440 [Delete Operation for Two Strings](https://leetcode.com/problems/delete-operation-for-two-strings/) · F

- [ ] 441 [Uncrossed Lines](https://leetcode.com/problems/uncrossed-lines/) · F

- [ ] 442 [Minimum ASCII Delete Sum for Two Strings](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/) · F

- [ ] 443 [Edit Distance](https://leetcode.com/problems/edit-distance/) · F · **MUST**

- [ ] 444 [Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/) · F

- [ ] 445 [Interleaving String](https://leetcode.com/problems/interleaving-string/) · B

- [ ] 446 [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/) · B

- [ ] 447 [Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/) · B

- [ ] 448 [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) · B

- [ ] 449 [Minimum Insertion Steps to Make a String Palindrome](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/) · B

- [ ] 450 [Shortest Common Supersequence](https://leetcode.com/problems/shortest-common-supersequence/) · B

- [ ] 451 [Palindrome Partitioning II](https://leetcode.com/problems/palindrome-partitioning-ii/) · B

- [ ] 452 [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/) · B

- [ ] 453 [Wildcard Matching](https://leetcode.com/problems/wildcard-matching/) · T

- [ ] 454 [Scramble String](https://leetcode.com/problems/scramble-string/) · T

- [ ] 455 [Predict the Winner](https://leetcode.com/problems/predict-the-winner/) · T

- [ ] 456 [Stone Game](https://leetcode.com/problems/stone-game/) · T

- [ ] 457 [Stone Game II](https://leetcode.com/problems/stone-game-ii/) · T

- [ ] 458 [Burst Balloons](https://leetcode.com/problems/burst-balloons/) · T

### 24 — Tries & Bit Manipulation

- [ ] 459 [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) · F · **MUST**

- [ ] 460 [Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/) · F

- [ ] 461 [Replace Words](https://leetcode.com/problems/replace-words/) · F

- [ ] 462 [Map Sum Pairs](https://leetcode.com/problems/map-sum-pairs/) · F

- [ ] 463 [Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/) · F

- [ ] 464 [Search Suggestions System](https://leetcode.com/problems/search-suggestions-system/) · F

- [ ] 465 [Word Search II](https://leetcode.com/problems/word-search-ii/) · B

- [ ] 466 [Concatenated Words](https://leetcode.com/problems/concatenated-words/) · B

- [ ] 467 [Stream of Characters](https://leetcode.com/problems/stream-of-characters/) · B

- [ ] 468 [Single Number](https://leetcode.com/problems/single-number/) · B · **MUST**

- [ ] 469 [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/) · B

- [ ] 470 [Counting Bits](https://leetcode.com/problems/counting-bits/) · B

- [ ] 471 [Reverse Bits](https://leetcode.com/problems/reverse-bits/) · B

- [ ] 472 [Missing Number](https://leetcode.com/problems/missing-number/) · B · **MUST**

- [ ] 473 [Power of Two](https://leetcode.com/problems/power-of-two/) · T

- [ ] 474 [Single Number II](https://leetcode.com/problems/single-number-ii/) · T

- [ ] 475 [Single Number III](https://leetcode.com/problems/single-number-iii/) · T

- [ ] 476 [Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/) · T

- [ ] 477 [Bitwise AND of Numbers Range](https://leetcode.com/problems/bitwise-and-of-numbers-range/) · T

- [ ] 478 [Maximum XOR of Two Numbers in an Array](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/) · T

### 25 — Math, Geometry & Matrix Manipulation

- [ ] 479 [Spiral Matrix](https://leetcode.com/problems/spiral-matrix/) · B · **MUST**

- [ ] 480 [Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/) · B · **MUST**

- [ ] 481 [Palindrome Number](https://leetcode.com/problems/palindrome-number/) · F

- [ ] 482 [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) · F

- [ ] 483 [Integer to Roman](https://leetcode.com/problems/integer-to-roman/) · F

- [ ] 484 [Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/) · F

- [ ] 485 [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/) · F

- [ ] 486 [Factorial Trailing Zeroes](https://leetcode.com/problems/factorial-trailing-zeroes/) · F

- [ ] 487 [Add Digits](https://leetcode.com/problems/add-digits/) · B

- [ ] 488 [Power of Three](https://leetcode.com/problems/power-of-three/) · B

- [ ] 489 [Ugly Number](https://leetcode.com/problems/ugly-number/) · B

- [ ] 490 [Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/) · B

- [ ] 491 [Count Primes](https://leetcode.com/problems/count-primes/) · B

- [ ] 492 [Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/) · B

- [ ] 493 [Rotate Image](https://leetcode.com/problems/rotate-image/) · B

- [ ] 494 [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/) · B · **MUST**

- [ ] 495 [Rectangle Overlap](https://leetcode.com/problems/rectangle-overlap/) · T

- [ ] 496 [Rectangle Area](https://leetcode.com/problems/rectangle-area/) · T

- [ ] 497 [Max Points on a Line](https://leetcode.com/problems/max-points-on-a-line/) · T

- [ ] 498 [Detect Squares](https://leetcode.com/problems/detect-squares/) · T

- [ ] 499 [Random Pick with Weight](https://leetcode.com/problems/random-pick-with-weight/) · T · **MUST**

- [ ] 500 [Random Point in Non-overlapping Rectangles](https://leetcode.com/problems/random-point-in-non-overlapping-rectangles/) · T

### 26 — Design Problems & OA Hardening

- [ ] 501 [Design HashSet](https://leetcode.com/problems/design-hashset/) · F

- [ ] 502 [Design HashMap](https://leetcode.com/problems/design-hashmap/) · F

- [ ] 503 [Design Linked List](https://leetcode.com/problems/design-linked-list/) · F

- [ ] 504 [Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/) · F · **MUST**

- [ ] 505 [Insert Delete GetRandom O(1) - Duplicates allowed](https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/) · F

- [ ] 506 [LRU Cache](https://leetcode.com/problems/lru-cache/) · F · **MUST**

- [ ] 507 [LFU Cache](https://leetcode.com/problems/lfu-cache/) · B · **MUST**

- [ ] 508 [Snapshot Array](https://leetcode.com/problems/snapshot-array/) · B · **MUST**

- [ ] 509 [Design Browser History](https://leetcode.com/problems/design-browser-history/) · B

- [ ] 510 [Design Underground System](https://leetcode.com/problems/design-underground-system/) · B

- [ ] 511 [Logger Rate Limiter](https://leetcode.com/problems/logger-rate-limiter/) · B · **MUST**

- [ ] 512 [Design In-Memory File System](https://leetcode.com/problems/design-in-memory-file-system/) · B · **MUST**

- [ ] 513 [Design Tic-Tac-Toe](https://leetcode.com/problems/design-tic-tac-toe/) · B

- [ ] 514 [Design Snake Game](https://leetcode.com/problems/design-snake-game/) · B

- [ ] 515 [Design Excel Sum Formula](https://leetcode.com/problems/design-excel-sum-formula/) · T

- [ ] 516 [Design Search Autocomplete System](https://leetcode.com/problems/design-search-autocomplete-system/) · T

- [ ] 517 [Encode and Decode TinyURL](https://leetcode.com/problems/encode-and-decode-tinyurl/) · T · **MUST**

- [ ] 518 [All O`one Data Structure](https://leetcode.com/problems/all-oone-data-structure/) · T

- [ ] 519 [Design Authentication Manager](https://leetcode.com/problems/design-authentication-manager/) · T · **MUST**

- [ ] 520 [Design a Food Rating System](https://leetcode.com/problems/design-a-food-rating-system/) · T

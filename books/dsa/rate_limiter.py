from collections import deque, defaultdict
import math
from typing import Dict


class SlidingWindowLogRateLimiter:
    # Pattern: per-key sliding window log.
    # Invariant: each deque contains only timestamps inside the active window.
    # Complexity: O(1) amortized per request, O(limit * keys) space.
    # Interview line: keep exact recent request timestamps when precision matters.
    def __init__(self, limit: int, window_seconds: float):
        self.limit = limit
        self.window_seconds = window_seconds
        self.requests: Dict[str, deque] = defaultdict(deque)

    # Pattern: evict old timestamps, then count recent requests.
    # Invariant: if deque length is below limit, the request is allowed.
    # Complexity: O(1) amortized time, O(limit) space per active key.
    # Interview line: old requests expire from the left because timestamps arrive in order.
    def allow(self, key: str, now: float) -> bool:
        q = self.requests[key]
        cutoff = now - self.window_seconds

        while q and q[0] <= cutoff:
            q.popleft()

        if len(q) >= self.limit:
            return False

        q.append(now)
        return True


class TokenBucketRateLimiter:
    # Pattern: token bucket with lazy refill.
    # Invariant: tokens never exceed capacity and refill based on elapsed time.
    # Complexity: O(1) time per request, O(keys) space.
    # Interview line: token bucket permits bursts up to capacity while enforcing long-term rate.
    def __init__(self, capacity: int, refill_rate_per_second: float):
        self.capacity = capacity
        self.refill_rate = refill_rate_per_second
        self.tokens = defaultdict(lambda: float(capacity))
        self.last_seen = defaultdict(float)

    # Pattern: lazy refill then spend tokens.
    # Invariant: a request is allowed only when enough tokens are available for its cost.
    # Complexity: O(1) time, O(1) extra space per key touched.
    # Interview line: refill on demand instead of running a background timer.
    def allow(self, key: str, now: float, cost: float = 1.0) -> bool:
        elapsed = max(0.0, now - self.last_seen[key])
        self.tokens[key] = min(self.capacity, self.tokens[key] + elapsed * self.refill_rate)
        self.last_seen[key] = now

        if self.tokens[key] < cost:
            return False

        self.tokens[key] -= cost
        return True


class FixedWindowRateLimiter:
    # Pattern: fixed window counter by key and window id.
    # Invariant: each bucket counts requests for one key within one fixed time window.
    # Complexity: O(1) time per request, O(keys * windows) space unless cleaned up.
    # Interview line: fixed windows are simple but can allow boundary bursts.
    def __init__(self, limit: int, window_seconds: int):
        self.limit = limit
        self.window_seconds = window_seconds
        self.counts = defaultdict(int)

    # Pattern: compute window bucket, increment, compare with limit.
    # Invariant: count for the current bucket is the only state needed for the decision.
    # Complexity: O(1) time, O(1) extra space per touched bucket.
    # Interview line: divide time by window size to map requests into counters.
    def allow(self, key: str, now: float) -> bool:
        window_id = math.floor(now / self.window_seconds)
        bucket = (key, window_id)
        self.counts[bucket] += 1
        return self.counts[bucket] <= self.limit


class HealthCheckWindow:
    # Pattern: fixed-size rolling health window.
    # Invariant: deque stores the last window_size health events per service.
    # Complexity: O(1) time per event, O(window_size * services) space.
    # Interview line: keep a bounded failure window to avoid reacting to stale incidents.
    def __init__(self, window_size: int, max_failures: int):
        self.window_size = window_size
        self.max_failures = max_failures
        self.events = defaultdict(deque)
        self.failures = defaultdict(int)

    # Pattern: append newest health result and evict oldest if needed.
    # Invariant: failures[service] matches the number of False values in the deque.
    # Complexity: O(1) time, O(window_size) space per service.
    # Interview line: maintain a running failure count so each health decision is constant time.
    def record(self, service: str, ok: bool) -> bool:
        q = self.events[service]
        q.append(ok)
        if not ok:
            self.failures[service] += 1

        if len(q) > self.window_size:
            old = q.popleft()
            if not old:
                self.failures[service] -= 1

        return self.failures[service] <= self.max_failures


if __name__ == "__main__":
    limiter = SlidingWindowLogRateLimiter(limit=2, window_seconds=10)
    assert limiter.allow("api", 0.0) is True
    assert limiter.allow("api", 1.0) is True
    assert limiter.allow("api", 2.0) is False

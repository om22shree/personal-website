package dsa

import (
	"math"
)

// Pattern: per-key sliding window log.
// Invariant: each deque contains only timestamps inside the active window.
// Complexity: O(1) amortized per request, O(limit * keys) space.
// Interview line: keep exact recent request timestamps when precision matters.
type SlidingWindowLogRateLimiter struct {
	limit         int
	windowSeconds float64
	requests      map[string][]float64
}

func NewSlidingWindowLogRateLimiter(limit int, windowSeconds float64) *SlidingWindowLogRateLimiter {
	return &SlidingWindowLogRateLimiter{
		limit:         limit,
		windowSeconds: windowSeconds,
		requests:      make(map[string][]float64),
	}
}

// Pattern: evict old timestamps, then count recent requests.
// Invariant: if deque length is below limit, the request is allowed.
// Complexity: O(1) amortized time, O(limit) space per active key.
// Interview line: old requests expire from the left because timestamps arrive in order.
func (this *SlidingWindowLogRateLimiter) Allow(key string, now float64) bool {
	q := this.requests[key]
	cutoff := now - this.windowSeconds

	for len(q) > 0 && q[0] <= cutoff {
		q = q[1:]
	}

	if len(q) >= this.limit {
		this.requests[key] = q
		return false
	}

	q = append(q, now)
	this.requests[key] = q
	return true
}

// Pattern: token bucket with lazy refill.
// Invariant: tokens never exceed capacity and refill based on elapsed time.
// Complexity: O(1) time per request, O(keys) space.
// Interview line: token bucket permits bursts up to capacity while enforcing long-term rate.
type TokenBucketRateLimiter struct {
	capacity   float64
	refillRate float64
	tokens     map[string]float64
	lastSeen   map[string]float64
}

func NewTokenBucketRateLimiter(capacity int, refillRatePerSecond float64) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		capacity:   float64(capacity),
		refillRate: refillRatePerSecond,
		tokens:     make(map[string]float64),
		lastSeen:   make(map[string]float64),
	}
}

// Pattern: lazy refill then spend tokens.
// Invariant: a request is allowed only when enough tokens are available for its cost.
// Complexity: O(1) time, O(1) extra space per key touched.
// Interview line: refill on demand instead of running a background timer.
func (this *TokenBucketRateLimiter) Allow(key string, now float64, cost float64) bool {
	if _, ok := this.tokens[key]; !ok {
		this.tokens[key] = this.capacity
	}

	elapsed := now - this.lastSeen[key]
	if elapsed < 0 {
		elapsed = 0.0
	}

	refilled := this.tokens[key] + elapsed*this.refillRate
	if refilled > this.capacity {
		this.tokens[key] = this.capacity
	} else {
		this.tokens[key] = refilled
	}
	this.lastSeen[key] = now

	if this.tokens[key] < cost {
		return false
	}

	this.tokens[key] -= cost
	return true
}

type windowKey struct {
	key      string
	windowId int64
}

// Pattern: fixed window counter by key and window id.
// Invariant: each bucket counts requests for one key within one fixed time window.
// Complexity: O(1) time per request, O(keys * windows) space unless cleaned up.
// Interview line: fixed windows are simple but can allow boundary bursts.
type FixedWindowRateLimiter struct {
	limit         int
	windowSeconds float64
	counts        map[windowKey]int
}

func NewFixedWindowRateLimiter(limit int, windowSeconds float64) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		limit:         limit,
		windowSeconds: windowSeconds,
		counts:        make(map[windowKey]int),
	}
}

// Pattern: compute window bucket, increment, compare with limit.
// Invariant: count for the current bucket is the only state needed for the decision.
// Complexity: O(1) time, O(1) extra space per touched bucket.
// Interview line: divide time by window size to map requests into counters.
func (this *FixedWindowRateLimiter) Allow(key string, now float64) bool {
	windowId := int64(math.Floor(now / this.windowSeconds))
	wk := windowKey{key: key, windowId: windowId}
	this.counts[wk]++
	return this.counts[wk] <= this.limit
}

// Pattern: fixed-size rolling health window.
// Invariant: deque stores the last window_size health events per service.
// Complexity: O(1) time per event, O(window_size * services) space.
// Interview line: keep a bounded failure window to avoid reacting to stale incidents.
type HealthCheckWindow struct {
	windowSize  int
	maxFailures int
	events      map[string][]bool
	failures    map[string]int
}

func NewHealthCheckWindow(windowSize int, maxFailures int) *HealthCheckWindow {
	return &HealthCheckWindow{
		windowSize:  windowSize,
		maxFailures: maxFailures,
		events:      make(map[string][]bool),
		failures:    make(map[string]int),
	}
}

// Pattern: append newest health result and evict oldest if needed.
// Invariant: failures[service] matches the number of False values in the deque.
// Complexity: O(1) time, O(window_size) space per service.
// Interview line: maintain a running failure count so each health decision is constant time.
func (this *HealthCheckWindow) Record(service string, ok bool) bool {
	q := this.events[service]
	q = append(q, ok)
	if !ok {
		this.failures[service]++
	}

	if len(q) > this.windowSize {
		old := q[0]
		q = q[1:]
		if !old {
			this.failures[service]--
		}
	}

	this.events[service] = q
	return this.failures[service] <= this.maxFailures
}

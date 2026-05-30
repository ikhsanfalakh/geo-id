package middleware

import (
	"sync"
	"time"
)

// RateLimiter is an in-memory, thread-safe Sliding Window rate limiter.
// It is designed to be easily replaced by a Redis-backed implementation
// in the future — just swap out the store behind the same interface.
type RateLimiter struct {
	mu      sync.Mutex
	windows map[string]*windowEntry
	limit   int           // max requests per window
	window  time.Duration // window duration
}

// windowEntry holds the sliding window state for one identifier.
type windowEntry struct {
	timestamps []time.Time
}

// LimitResult contains the outcome of a rate limit check.
type LimitResult struct {
	Allowed   bool
	Limit     int
	Remaining int
	ResetAt   time.Time // Unix epoch of when the oldest token exits the window
}

// NewRateLimiter creates a RateLimiter with the given limit and window duration.
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		windows: make(map[string]*windowEntry),
		limit:   limit,
		window:  window,
	}

	// Background goroutine to periodically evict idle entries and prevent
	// unbounded memory growth.
	go rl.cleanupLoop()

	return rl
}

// Check evaluates whether the identifier is within its rate limit.
// It records the current request timestamp and returns a LimitResult.
// Thread-safe.
func (rl *RateLimiter) Check(identifier string) LimitResult {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)

	entry, exists := rl.windows[identifier]
	if !exists {
		entry = &windowEntry{}
		rl.windows[identifier] = entry
	}

	// Evict timestamps older than the window boundary (sliding window).
	entry.timestamps = filterAfter(entry.timestamps, cutoff)

	count := len(entry.timestamps)

	if count >= rl.limit {
		// Limit exceeded — compute when the oldest request will expire.
		resetAt := entry.timestamps[0].Add(rl.window)
		return LimitResult{
			Allowed:   false,
			Limit:     rl.limit,
			Remaining: 0,
			ResetAt:   resetAt,
		}
	}

	// Within limit — record this request.
	entry.timestamps = append(entry.timestamps, now)
	remaining := rl.limit - len(entry.timestamps)

	// ResetAt: when the oldest request in the window will expire.
	resetAt := now.Add(rl.window)
	if len(entry.timestamps) > 0 {
		resetAt = entry.timestamps[0].Add(rl.window)
	}

	return LimitResult{
		Allowed:   true,
		Limit:     rl.limit,
		Remaining: remaining,
		ResetAt:   resetAt,
	}
}

// filterAfter returns only timestamps strictly after the cutoff (in-place reuse of slice).
func filterAfter(ts []time.Time, cutoff time.Time) []time.Time {
	i := 0
	for _, t := range ts {
		if t.After(cutoff) {
			ts[i] = t
			i++
		}
	}
	return ts[:i]
}

// cleanupLoop removes identifiers whose windows have gone completely idle.
// Runs every minute to keep memory usage bounded.
func (rl *RateLimiter) cleanupLoop() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		rl.mu.Lock()
		cutoff := time.Now().Add(-rl.window)
		for id, entry := range rl.windows {
			entry.timestamps = filterAfter(entry.timestamps, cutoff)
			if len(entry.timestamps) == 0 {
				delete(rl.windows, id)
			}
		}
		rl.mu.Unlock()
	}
}

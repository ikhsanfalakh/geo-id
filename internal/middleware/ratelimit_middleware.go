package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	// DefaultLimitAnonymous is the fallback cap per minute for unauthenticated clients.
	DefaultLimitAnonymous = 60
	// DefaultLimitAPIKey is the fallback cap per minute for authenticated API key clients.
	DefaultLimitAPIKey = 1000

	// windowDuration is the sliding window interval.
	windowDuration = time.Minute

	// headerAPIKey is the name of the request header carrying the API key.
	headerAPIKey = "X-API-KEY"
)

// excludedPrefixes lists URL path prefixes that are exempt from rate limiting.
var excludedPrefixes = []string{
	"/apidocs",
	"/swagger",
	"/actuator/health",
	"/assets",
}

// RateLimitConfig holds the dependencies for the rate limit middleware.
type RateLimitConfig struct {
	APIKeyService    *APIKeyService
	AnonymousLimiter *RateLimiter
	APIKeyLimiter    *RateLimiter
}

// NewRateLimitConfig constructs the config with the given per-minute limits.
// limitAnonymous and limitAPIKey are typically read from environment variables
// (RATE_LIMIT_ANONYMOUS, RATE_LIMIT_API_KEY) and fall back to the defaults
// (DefaultLimitAnonymous, DefaultLimitAPIKey) when unset.
func NewRateLimitConfig(apiKeySvc *APIKeyService, limitAnonymous, limitAPIKey int) *RateLimitConfig {
	return &RateLimitConfig{
		APIKeyService:    apiKeySvc,
		AnonymousLimiter: NewRateLimiter(limitAnonymous, windowDuration),
		APIKeyLimiter:    NewRateLimiter(limitAPIKey, windowDuration),
	}
}

// RateLimitMiddleware returns a Fiber handler that:
//  1. Skips excluded paths (docs, health, static assets).
//  2. Validates X-API-KEY header if present.
//  3. Applies the correct rate limiter (anonymous or API-key tier).
//  4. Injects X-RateLimit-* response headers.
//  5. Returns 429 when the limit is exceeded, 401 for invalid API keys.
func RateLimitMiddleware(cfg *RateLimitConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()

		// Skip rate limiting for excluded paths.
		for _, prefix := range excludedPrefixes {
			if strings.HasPrefix(path, prefix) {
				return c.Next()
			}
		}

		rawKey := c.Get(headerAPIKey)

		// --- API Key branch ---
		if rawKey != "" {
			if !cfg.APIKeyService.IsValid(rawKey) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"success": false,
					"error": fiber.Map{
						"code":    "INVALID_API_KEY",
						"message": "Invalid API key",
					},
				})
			}

			result := cfg.APIKeyLimiter.Check(rawKey)
			setRateLimitHeaders(c, result)

			if !result.Allowed {
				return rateLimitExceededResponse(c)
			}

			return c.Next()
		}

		// --- Anonymous branch ---
		clientIP := getClientIP(c)
		result := cfg.AnonymousLimiter.Check(clientIP)
		setRateLimitHeaders(c, result)

		if !result.Allowed {
			return rateLimitExceededResponse(c)
		}

		return c.Next()
	}
}

// setRateLimitHeaders writes the standard rate limit headers to the response.
func setRateLimitHeaders(c *fiber.Ctx, r LimitResult) {
	c.Set("X-RateLimit-Limit", itoa(r.Limit))
	c.Set("X-RateLimit-Remaining", itoa(r.Remaining))
	c.Set("X-RateLimit-Reset", itoa(int(r.ResetAt.Unix())))
}

// rateLimitExceededResponse returns the standardised 429 error response.
func rateLimitExceededResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    "RATE_LIMIT_EXCEEDED",
			"message": "Too many requests",
		},
	})
}

// getClientIP extracts the real client IP, honouring common proxy headers.
func getClientIP(c *fiber.Ctx) string {
	// Respect X-Forwarded-For when behind a reverse proxy.
	if xff := c.Get("X-Forwarded-For"); xff != "" {
		// X-Forwarded-For may be a comma-separated list; take the first entry.
		if idx := strings.IndexByte(xff, ','); idx != -1 {
			return strings.TrimSpace(xff[:idx])
		}
		return strings.TrimSpace(xff)
	}
	if xri := c.Get("X-Real-IP"); xri != "" {
		return xri
	}
	return c.IP()
}

// itoa converts an int to its string representation without importing strconv.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [20]byte
	pos := len(buf)
	for n >= 10 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	pos--
	buf[pos] = byte('0' + n)
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}

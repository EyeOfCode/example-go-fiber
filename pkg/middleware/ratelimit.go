package middleware

import (
	"go-fiber-api/pkg/utils"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RateLimiter struct {
	rate     int
	interval time.Duration
	mu       sync.Mutex
	tokens   map[string][]time.Time
}

func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:     rate,
		interval: interval,
		tokens:   make(map[string][]time.Time),
	}
}

func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.interval)

	if _, exists := rl.tokens[key]; !exists {
		rl.tokens[key] = []time.Time{now}
		return true
	}

	var validTokens []time.Time
	for _, t := range rl.tokens[key] {
		if t.After(windowStart) {
			validTokens = append(validTokens, t)
		}
	}

	if len(validTokens) < rl.rate {
		validTokens = append(validTokens, now)
		rl.tokens[key] = validTokens
		return true
	}

	rl.tokens[key] = validTokens
	return false
}

func RateLimit(rate int, interval time.Duration) fiber.Handler {
	limiter := NewRateLimiter(rate, interval)

	return func(c *fiber.Ctx) error {
		// use ip
		key := c.IP()

		// or use user id on jwt
		// if user, ok := c.Locals("user").(string); ok {
		// 		key = user
		// }

		if !limiter.Allow(key) {
			return utils.SendError(c, http.StatusTooManyRequests, "Rate limit exceeded. Please try again later.")
		}

		return c.Next()
	}
}

package ratelimiter

import (
	"time"
)

func StartRateLimiter(ch chan struct{}, capacity int, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		select {
		case ch <- struct{}{}:
		default:
		}
	}
}

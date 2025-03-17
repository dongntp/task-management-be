package httpclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"task-management-be/internal/pkg/env"
	"time"

	"golang.org/x/time/rate"
)

type RateLimitedClient struct {
	Client      *http.Client
	Timeout     time.Duration
	RateLimiter *rate.Limiter
}

func NewRateLimitedClient(limit envlib.HTTPLimit) *RateLimitedClient {
	rateLimiter := rate.NewLimiter(rate.Limit(limit.MaxQueryRate), limit.MaxQueryRate)
	return &RateLimitedClient{
		Client:      http.DefaultClient,
		Timeout:     limit.Timeout,
		RateLimiter: rateLimiter,
	}
}

// Implement interface HttpRequestDoer
func (c *RateLimitedClient) Do(req *http.Request) (*http.Response, error) {
	if c.RateLimiter == nil {
		return nil, errors.New("rate limiter not set")
	}
	ctxWithTimeout, cancel := context.WithTimeout(req.Context(), c.Timeout)
	defer cancel()

	err := c.RateLimiter.Wait(ctxWithTimeout)
	if err != nil {
		return nil, fmt.Errorf("rate limiter throttled with error: %w", err)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		if ctxWithTimeout.Err() != nil {
			err = errors.Join(ctxWithTimeout.Err(), err)
		}
		return nil, err
	}
	return resp, nil
}

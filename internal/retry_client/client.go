package retry_client

import (
	"context"
	"errors"
	"net/http"
	"time"
)

var (
	timeSleepDuration = []time.Duration{
		1 * time.Second,
		2 * time.Second,
		3 * time.Second,
		5 * time.Second,
		8 * time.Second,
	}
)

func New(client *http.Client) *RetryClient {
	return &RetryClient{
		client: client,
	}
}

type RetryClient struct {
	client *http.Client
}

func (c *RetryClient) Do(_ context.Context, req *http.Request) (*http.Response, error) {
	var attempts int

	resp, err := c.client.Do(req)
	for c.needRetry(resp, err) {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}

		time.Sleep(timeSleepDuration[min(attempts, len(timeSleepDuration)-1)])
		resp, err = c.client.Do(req)
		attempts++
	}

	return resp, err
}

func (c *RetryClient) needRetry(resp *http.Response, err error) bool {
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return false
	}

	if err != nil {
		return true
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return true
	}

	if resp.StatusCode >= 500 {
		return true
	}

	return false
}

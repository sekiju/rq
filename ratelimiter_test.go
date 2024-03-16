package rq

import (
	"testing"
	"time"
)

var margin = time.Millisecond * 100

func TestRateLimiter_RPS(t *testing.T) {
	rateLimiter := NewRateLimiter(SetRPS(5))
	startTime := time.Now()

	for i := 0; i < 5; i++ {
		rateLimiter.Wait()
	}

	elapsedTime := time.Since(startTime)
	expectedDuration := time.Second

	if elapsedTime < expectedDuration-margin || elapsedTime > expectedDuration+margin {
		t.Errorf("Elapsed time %s is not within expected range [%s, %s]", elapsedTime, expectedDuration-margin, expectedDuration+margin)
	}
}

func TestRateLimiter_RPM(t *testing.T) {
	rateLimiter := NewRateLimiter(SetRPM(240))
	startTime := time.Now()

	for i := 0; i < 20; i++ {
		rateLimiter.Wait()
	}

	elapsedTime := time.Since(startTime)
	expectedDuration := 5 * time.Second

	if elapsedTime < expectedDuration-margin || elapsedTime > expectedDuration+margin {
		t.Errorf("Elapsed time %s is not within expected range [%s, %s]", elapsedTime, expectedDuration-margin, expectedDuration+margin)
	}
}

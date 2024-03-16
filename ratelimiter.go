package rq

import "time"

// RateLimiter manages rate limiting based on specified options.
type RateLimiter struct {
	Options         RateLimiterOpts
	lastTokenUpdate time.Time
	tokens          int
}

// RateLimiterOpts contains options for rate limiting.
type RateLimiterOpts struct {
	RPS *int
	RPM *int
}

// RateLimiterOptsFn is a function type for setting RateLimiterOpts.
type RateLimiterOptsFn func(*RateLimiterOpts)

// SetRPS sets the requests per second value in RateLimiterOpts.
func SetRPS(value int) RateLimiterOptsFn {
	return func(o *RateLimiterOpts) {
		o.RPS = &value
	}
}

// SetRPM sets the requests per minute value in RateLimiterOpts.
func SetRPM(value int) RateLimiterOptsFn {
	return func(o *RateLimiterOpts) {
		o.RPM = &value
	}
}

// NewRateLimiter creates a new RateLimiter instance with specified options.
func NewRateLimiter(opts ...RateLimiterOptsFn) *RateLimiter {
	cfg := RateLimiterOpts{}
	for _, fn := range opts {
		fn(&cfg)
	}
	return &RateLimiter{
		Options: cfg,
	}
}

// Wait blocks until it's allowed to proceed based on the rate limit.
func (rl *RateLimiter) Wait() {
	currentTime := time.Now()
	var interval time.Duration

	if rl.Options.RPS != nil && rl.Options.RPM != nil {
		rpsInterval := time.Second / time.Duration(*rl.Options.RPS)
		rpmInterval := time.Minute / time.Duration(*rl.Options.RPM)
		interval = lcm(rpsInterval, rpmInterval)
	} else if rl.Options.RPS != nil {
		interval = time.Second / time.Duration(*rl.Options.RPS)
	} else if rl.Options.RPM != nil {
		interval = time.Minute / time.Duration(*rl.Options.RPM)
	} else {
		return
	}

	rl.updateTokens(currentTime, interval)
	rl.waitForTokens(1, interval)
}

func lcm(a, b time.Duration) time.Duration {
	return (a * b) / gcd(a, b)
}

func gcd(a, b time.Duration) time.Duration {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (rl *RateLimiter) updateTokens(currentTime time.Time, interval time.Duration) {
	if rl.lastTokenUpdate.IsZero() {
		rl.lastTokenUpdate = currentTime
	}
	elapsed := currentTime.Sub(rl.lastTokenUpdate)
	tokensToAdd := int(elapsed / interval)
	rl.tokens += tokensToAdd
	if rl.tokens < 0 {
		rl.tokens = 0
	}
	rl.lastTokenUpdate = rl.lastTokenUpdate.Add(time.Duration(tokensToAdd) * interval)
}

func (rl *RateLimiter) waitForTokens(n int, interval time.Duration) {
	requiredTokens := n
	sleepDuration := (time.Duration(requiredTokens) * interval) - time.Since(rl.lastTokenUpdate)

	if sleepDuration > 0 {
		time.Sleep(sleepDuration)
		rl.updateTokens(time.Now(), interval)
	}
}

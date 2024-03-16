package rq

type Opts struct {
	Headers     map[string]string
	Body        interface{}
	RateLimiter *RateLimiter
}

type OptsFn func(*Opts)

func defaultOpts() Opts {
	return Opts{}
}

func SetHeader(key, value string) OptsFn {
	return func(o *Opts) {
		o.Headers[key] = value
	}
}

func SetBody(v interface{}) OptsFn {
	return func(o *Opts) {
		o.Body = v
	}
}

func SetRateLimiter(rateLimiter *RateLimiter) OptsFn {
	return func(o *Opts) {
		o.RateLimiter = rateLimiter
	}
}

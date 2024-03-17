package rq

type Opts struct {
	Headers map[string]string
	Body    interface{}
}

type OptsFn func(*Opts)

func defaultOpts() Opts {
	return Opts{
		Headers: make(map[string]string),
	}
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

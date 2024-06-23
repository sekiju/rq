package rq

type BodyType string

const (
	RawBodyType  BodyType = "raw"
	JsonBodyType BodyType = "json"
)

type Opts struct {
	Headers  map[string]string
	Body     interface{}
	BodyType BodyType
}

type OptsFn func(*Opts)

func defaultOpts() Opts {
	return Opts{
		Headers:  map[string]string{},
		BodyType: JsonBodyType,
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

func SetBodyType(bodyType BodyType) OptsFn {
	return func(o *Opts) {
		o.BodyType = bodyType
	}
}

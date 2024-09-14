package rq

type (
	BodyEncoding string

	Config struct {
		Headers      map[string]string
		Body         interface{}
		BodyEncoding BodyEncoding
	}

	Opt func(*Config)
)

const (
	RawEncoding  BodyEncoding = "raw"
	JsonEncoding BodyEncoding = "json"
)

func (c *Config) Set(opts ...Opt) {
	for _, fn := range opts {
		fn(c)
	}
}

func getDefaultOptions() Config {
	return Config{
		Headers:      map[string]string{},
		BodyEncoding: JsonEncoding,
	}
}

func SetHeader(key, value string) Opt {
	return func(c *Config) {
		c.Headers[key] = value
	}
}

func SetBody(v interface{}) Opt {
	return func(c *Config) {
		c.Body = v
	}
}

func SetBodyEncoding(encoding BodyEncoding) Opt {
	return func(c *Config) {
		c.BodyEncoding = encoding
	}
}

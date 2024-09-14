package rq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	Config *Config
}

func New(opts ...Opt) *Request {
	config := getDefaultOptions()
	for _, fn := range opts {
		fn(&config)
	}
	return &Request{Config: &config}
}

func (r *Request) Execute(method, URL string) (*Response, error) {
	var body io.Reader
	if r.Config.Body != nil {
		switch r.Config.BodyEncoding {
		case JsonEncoding:
			serialized, err := json.Marshal(r.Config.Body)
			if err != nil {
				return nil, err
			}

			body = bytes.NewBuffer(serialized)
		}
	}

	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}

	switch r.Config.BodyEncoding {
	case JsonEncoding:
		req.Header.Set("content-type", "application/json")
	}

	for key, value := range r.Config.Headers {
		req.Header.Set(key, value)
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{res}, nil
}

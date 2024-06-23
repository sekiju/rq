package rq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	opts *Opts
}

// New creates a new Request with the provided Opts.
func New(opts ...OptsFn) *Request {
	catOpts := defaultOpts()
	for _, fn := range opts {
		fn(&catOpts)
	}

	return &Request{opts: &catOpts}
}

func (r *Request) Get(URL string, args ...any) (*Response, error) {
	return r.constructor("GET", URL, args...)
}

func (r *Request) Post(URL string, args ...any) (*Response, error) {
	return r.constructor("POST", URL, args...)
}

func (r *Request) Put(URL string, args ...any) (*Response, error) {
	return r.constructor("PUT", URL, args...)
}

func (r *Request) Patch(URL string, args ...any) (*Response, error) {
	return r.constructor("PATCH", URL, args...)
}

func (r *Request) Delete(URL string, args ...any) (*Response, error) {
	return r.constructor("DELETE", URL, args...)
}

func (r *Request) constructor(method, URL string, args ...any) (*Response, error) {
	URL = fmt.Sprintf(URL, args...)

	var body io.Reader
	if r.opts.Body != nil {
		switch r.opts.BodyType {
		case JsonBodyType:
			serialized, err := json.Marshal(r.opts.Body)
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

	switch r.opts.BodyType {
	case JsonBodyType:
		req.Header.Set("Content-Type", "application/json")
	}

	if r.opts.Headers != nil {
		for key, value := range r.opts.Headers {
			req.Header.Set(key, value)
		}
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	return NewResponse(res, r.opts), nil
}

package rq

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type postMethod struct{}

// Post is a repository for POST methods.
var Post = postMethod{}

// JSON encodes the interface into JSON format, and adds "application/json" header.
//
// Deprecated: use rq.New().Post() instead.
func (postMethod) JSON(url string, opts ...OptsFn) (*Response, error) {
	cfg := defaultOpts()
	for _, fn := range opts {
		fn(&cfg)
	}

	serialized, err := json.Marshal(cfg.Body)
	if err != nil {
		return nil, err
	}

	req, err := newRequest("POST", url, bytes.NewBuffer(serialized), cfg)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	return NewResponse(res, &cfg), nil
}

// Bytes send []byte.
//
// Deprecated: use rq.New(rq.SetBodyType(RawBodyType)).Post() instead.
func (postMethod) Bytes(url string, opts ...OptsFn) (*Response, error) {
	cfg := defaultOpts()
	for _, fn := range opts {
		fn(&cfg)
	}

	buf, ok := cfg.Body.([]byte)
	if !ok {
		return nil, fmt.Errorf("body should initially be []byte")
	}

	req, err := newRequest("POST", url, bytes.NewBuffer(buf), cfg)
	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	return NewResponse(res, &cfg), nil
}

package rq

// Get is GET method.
func Get(url string, opts ...OptsFn) (*Response, error) {
	cfg := defaultOpts()
	for _, fn := range opts {
		fn(&cfg)
	}

	req, err := newRequest("GET", url, nil, cfg)
	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	return NewResponse(res), nil
}

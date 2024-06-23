package rq

import (
	"io"
	"net/http"
)

// Deprecated: used only for old requests.
func newRequest(method, url string, body io.Reader, opts Opts) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	if opts.Headers != nil {
		for key, value := range opts.Headers {
			req.Header.Set(key, value)
		}
	}

	return req, nil
}

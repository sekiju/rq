package round_tripper

import "net/http"

type HeaderRoundTripper struct {
	transport http.RoundTripper
	headers   map[string]string
}

func NewHeaderRoundTripper(transport http.RoundTripper, headers map[string]string) http.RoundTripper {
	if transport == nil {
		transport = http.DefaultTransport
	}

	return &HeaderRoundTripper{transport: transport, headers: headers}
}

func (t *HeaderRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	for header, value := range t.headers {
		if _, ok := r.Header[header]; !ok {
			r.Header.Set(header, value)
		}
	}

	return t.transport.RoundTrip(r)
}

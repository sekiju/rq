package rq

import (
	"crypto/tls"
	"github.com/sekiju/rq/pkg/util/round_tripper"
	"net/http"
	"time"
)

// Client is the global HTTP client used by the package.
var Client = DefaultClient()

// DefaultClient creates a new HTTP client with custom transport settings.
func DefaultClient() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256, tls.CurveP384, tls.CurveP521, tls.X25519},
		},
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := http.DefaultClient
	client.Transport = round_tripper.NewHeaderRoundTripper(
		transport,
		map[string]string{
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			"Accept-Language": "en-US,en;q=0.5",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
		},
	)

	return client
}

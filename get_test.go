package rq

import (
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://httpbin.org/get"

	t.Run("Basic GET Request", func(t *testing.T) {
		res, err := Get(url)
		if err != nil {
			t.Fatalf("Error making GET request: %v", err)
		}

		if !res.Ok {
			t.Error("Response failed, status not ok")
		}

		var v HttpBinGetResponse
		if err := res.JSON(&v); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		if v.Url != url {
			t.Errorf("Got unexpected URL: %s", v.Url)
		}
	})

	t.Run("GET Request with Header", func(t *testing.T) {
		res, err := Get(url, SetHeader("X", "y"))
		if err != nil {
			t.Fatalf("Error making GET request with header: %v", err)
		}

		if !res.Ok {
			t.Error("Response failed, status not ok")
		}

		var v HttpBinGetResponse
		if err := res.JSON(&v); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		header, exists := v.Headers["X"]
		if !exists {
			t.Error("Expected header does not exist")
		}

		if header != "y" {
			t.Errorf("Got unexpected header value: %s", header)
		}
	})
}

type HttpBinGetResponse struct {
	Args struct {
	} `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

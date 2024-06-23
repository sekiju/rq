package rq

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRequest(t *testing.T) {
	url := "https://httpbin.org/get"

	t.Run("GET request", func(t *testing.T) {
		res, err := New().Get(url)
		if err != nil {
			t.Fatalf("Error making GET request: %v", err)
		}

		if !res.Ok {
			t.Error("Response failed, status not ok")
		}

		var v HttpBinGetResponse
		if err = res.JSON(&v); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		if v.Url != url {
			t.Errorf("Got unexpected URL: %s", v.Url)
		}
	})

	t.Run("GET request with URL and query params", func(t *testing.T) {
		res, err := New().Get("https://httpbin.org/get?param=%s", "value")
		if err != nil {
			t.Fatalf("Error making GET request: %v", err)
		}

		if !res.Ok {
			t.Errorf("Expected status code 200, got %d", res.RawResponse.StatusCode)
		}

		var v HttpBinGetResponse
		if err = res.JSON(&v); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		query, exists := v.Args["param"]
		if !exists {
			t.Error("Expected query param does not exist")
		}

		if query != "value" {
			t.Errorf("Got unexpected query param: %s", v.Url)
		}
	})

	t.Run("GET request with Header", func(t *testing.T) {
		res, err := New(SetHeader("X", "y")).Get(url)
		if err != nil {
			t.Fatalf("Error making GET request with header: %v", err)
		}

		if !res.Ok {
			t.Errorf("Expected status code 200, got %d", res.RawResponse.StatusCode)
		}

		var v HttpBinGetResponse
		if err = res.JSON(&v); err != nil {
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

	t.Run("POST request with JSON data", func(t *testing.T) {
		testData := map[string]interface{}{
			"name":     "sekiju/rq",
			"birthday": 1710018066000,
			"is_child": true,
		}

		res, err := New(SetBody(testData)).Post("https://httpbin.org/post")
		if err != nil {
			t.Fatalf("Error making POST request: %v", err)
		}

		if !res.Ok {
			t.Errorf("Expected status code 200, got %d", res.RawResponse.StatusCode)
		}

		var v HttpBinPostResponse
		if err = res.JSON(&v); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		testDataBytes, err := json.Marshal(testData)
		if err != nil {
			t.Fatalf("Error marshalling test data: %v", err)
		}

		expectedData := make(map[string]interface{})
		if err = json.Unmarshal(testDataBytes, &expectedData); err != nil {
			t.Fatalf("Error unmarshalling test data: %v", err)
		}

		if !reflect.DeepEqual(v.Json, expectedData) {
			t.Error("Response data mismatch")
		}
	})
}

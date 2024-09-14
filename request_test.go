package rq

import (
	"encoding/json"
	"reflect"
	"testing"
)

type (
	HttpBinGetResponse struct {
		Args    map[string]string `json:"args"`
		Headers map[string]string `json:"headers"`
		Origin  string            `json:"origin"`
		Url     string            `json:"url"`
	}

	HttpBinPostResponse struct {
		Args    map[string]string      `json:"args"`
		Data    string                 `json:"data"`
		Files   map[string]string      `json:"files"`
		Form    map[string]string      `json:"form"`
		Headers map[string]string      `json:"headers"`
		Json    map[string]interface{} `json:"json"`
		Origin  string                 `json:"origin"`
		Url     string                 `json:"url"`
	}
)

func TestRequest(t *testing.T) {
	t.Run("GET request", func(t *testing.T) {
		res, err := New().Get("https://httpbin.org/get")
		if err != nil {
			t.Fatalf("Error making GET request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
		}

		var v HttpBinGetResponse
		if err = res.JSON(&v); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		if v.Url != "https://httpbin.org/get" {
			t.Errorf("Got unexpected URL: %s", v.Url)
		}
	})

	t.Run("GET request with Getf method", func(t *testing.T) {
		res, err := New().Getf("https://httpbin.org/get?param=%s", "value")
		if err != nil {
			t.Fatalf("Error making GET request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
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
		res, err := New(SetHeader("X", "y")).Get("https://httpbin.org/get")
		if err != nil {
			t.Fatalf("Error making GET request with header: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
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

	t.Run("GET request with Header, but dynamic options", func(t *testing.T) {
		req := New()
		req.Config.Set(SetHeader("X", "y"))

		res, err := req.Get("https://httpbin.org/get")
		if err != nil {
			t.Fatalf("Error making GET request with header: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
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

	testData := map[string]interface{}{"is_fiber": true}

	t.Run("POST request with JSON data", func(t *testing.T) {
		res, err := New(SetBody(testData)).Post("https://httpbin.org/post")
		if err != nil {
			t.Fatalf("Error making POST request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
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

	t.Run("POST request with Postf method", func(t *testing.T) {
		res, err := New().Postf("https://httpbin.org/post?param=%s", "value")
		if err != nil {
			t.Fatalf("Error making DELETE request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
		}

		var v HttpBinPostResponse
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

	t.Run("PUT request with JSON data", func(t *testing.T) {
		res, err := New(SetBody(testData)).Put("https://httpbin.org/put")
		if err != nil {
			t.Fatalf("Error making PUT request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
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

	t.Run("PUT request with Putf method", func(t *testing.T) {
		res, err := New().Putf("https://httpbin.org/put?param=%s", "value")
		if err != nil {
			t.Fatalf("Error making DELETE request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
		}

		var v HttpBinPostResponse
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

	t.Run("PATCH request with JSON data", func(t *testing.T) {
		res, err := New(SetBody(testData)).Patch("https://httpbin.org/patch")
		if err != nil {
			t.Fatalf("Error making PUT request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
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

	t.Run("PATCH request with Patchf method", func(t *testing.T) {
		res, err := New().Patchf("https://httpbin.org/patch?param=%s", "value")
		if err != nil {
			t.Fatalf("Error making DELETE request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
		}

		var v HttpBinPostResponse
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

	t.Run("DELETE request", func(t *testing.T) {
		res, err := New().Delete("https://httpbin.org/delete")
		if err != nil {
			t.Fatalf("Error making DELETE request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
		}
	})

	t.Run("DELETE request with Deletef method", func(t *testing.T) {
		res, err := New().Deletef("https://httpbin.org/delete?param=%s", "value")
		if err != nil {
			t.Fatalf("Error making DELETE request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
		}

		var v HttpBinPostResponse
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

	t.Run("OPTIONS request with Execute method", func(t *testing.T) {
		res, err := New().Execute("OPTIONS", "https://httpbin.org/get")
		if err != nil {
			t.Fatalf("Error making DELETE request: %v", err)
		}

		if res.StatusCode != 200 {
			t.Error("Response failed, invalid status: " + res.Status)
		}
	})
}

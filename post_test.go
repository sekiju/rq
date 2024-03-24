package rq

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPost(t *testing.T) {
	t.Run("Basic POST Request with JSON Data", func(t *testing.T) {
		testData := map[string]interface{}{
			"name":     "Konstantin",
			"age":      20,
			"isAuthor": true,
		}

		res, err := Post.JSON("https://httpbin.org/post", SetBody(testData))
		if err != nil {
			t.Fatalf("Error making POST request: %v", err)
		}

		if !res.Ok {
			t.Error("Response failed, status not ok")
		}

		var v HttpBinPostResponse
		if err := res.JSON(&v); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		testDataBytes, err := json.Marshal(testData)
		if err != nil {
			t.Fatalf("Error marshalling test data: %v", err)
		}

		expectedData := make(map[string]interface{})
		if err := json.Unmarshal(testDataBytes, &expectedData); err != nil {
			t.Fatalf("Error unmarshalling test data: %v", err)
		}

		if !reflect.DeepEqual(v.Json, expectedData) {
			t.Error("Response data mismatch")
		}
	})
}

type HttpBinPostResponse struct {
	Args    map[string]string      `json:"args"`
	Data    string                 `json:"data"`
	Files   map[string]string      `json:"files"`
	Form    map[string]string      `json:"form"`
	Headers map[string]string      `json:"headers"`
	Json    map[string]interface{} `json:"json"`
	Origin  string                 `json:"origin"`
	Url     string                 `json:"url"`
}

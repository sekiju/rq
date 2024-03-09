package rq

import (
	"reflect"
	"testing"
)

func TestPostMethod_JSON(t *testing.T) {
	testData := map[string]interface{}{
		"foo": "bar",
		"one": 1,
		"two": true,
	}

	res, err := Post.JSON("https://httpbin.org/post", SetBody(testData))
	if err != nil {
		t.Error(err)
	}

	if !res.Ok {
		t.Error("response failed, status not ok")
	}

	var v HttpBinPostResponse
	if err = res.JSON(&v); err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(v.Json, testData) {
		t.Error("invalid response data")
	}
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

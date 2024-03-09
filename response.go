package rq

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Response struct {
	RawResponse *http.Response
	Ok          bool
}

func NewResponse(rawResponse *http.Response) *Response {
	newResponse := &Response{RawResponse: rawResponse}

	if rawResponse.StatusCode >= 200 && rawResponse.StatusCode < 300 {
		newResponse.Ok = true
	}

	return newResponse
}

func (r Response) JSON(v interface{}) error {
	body, err := io.ReadAll(r.RawResponse.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

func (r Response) Text() (string, error) {
	body, err := io.ReadAll(r.RawResponse.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (r Response) Bytes() ([]byte, error) {
	return io.ReadAll(r.RawResponse.Body)
}

func (r Response) UnmarshallAuto(v interface{}) error {
	if r.RawResponse == nil {
		return errors.New("no body")
	}

	switch r.RawResponse.Header.Get("content-type") {
	case "application/json", "application/json; charset=utf-8", "application/json; charset=UTF-8":
		return r.JSON(v)
	}

	return nil
}

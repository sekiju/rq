package rq

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

type Response struct {
	*http.Response
}

func (r *Response) JSON(v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

func (r *Response) XML(v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return xml.Unmarshal(body, &v)
}

func (r *Response) Text() (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (r *Response) Bytes() ([]byte, error) {
	return io.ReadAll(r.Body)
}

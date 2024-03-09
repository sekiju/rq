package rq

import (
	"testing"
)

func TestGet(t *testing.T) {
	res, err := Get("https://httpbin.org/get")
	if err != nil {
		t.Error(err)
	}

	if !res.Ok {
		t.Error("response failed, status not ok")
	}

	var v HttpBinGetResponse
	if err = res.JSON(&v); err != nil {
		t.Error(err)
	}

	if v.Url != "https://httpbin.org/get" {
		t.Error("unmarshall error")
	}
}

type HttpBinGetResponse struct {
	Args struct {
	} `json:"args"`
	Headers struct {
		Accept                  string `json:"Accept"`
		AcceptEncoding          string `json:"Accept-Encoding"`
		AcceptLanguage          string `json:"Accept-Language"`
		Host                    string `json:"Host"`
		SecChUa                 string `json:"Sec-Ch-Ua"`
		SecChUaMobile           string `json:"Sec-Ch-Ua-Mobile"`
		SecChUaPlatform         string `json:"Sec-Ch-Ua-Platform"`
		SecFetchDest            string `json:"Sec-Fetch-Dest"`
		SecFetchMode            string `json:"Sec-Fetch-Mode"`
		SecFetchSite            string `json:"Sec-Fetch-Site"`
		SecFetchUser            string `json:"Sec-Fetch-User"`
		UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests"`
		UserAgent               string `json:"User-Agent"`
		XAmznTraceId            string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

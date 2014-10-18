package hydrator

import (
	"net/http"
	"net/url"
)

func Proxy(req *http.Request) (Response, error) {
	upstreamRequest := createRequest(req)
	response, err := http.DefaultClient.Do(upstreamRequest)
	if err != nil {
		return nil, err
	}
	readResponse := NewNormalResponse(response)
	hydrateField := response.Header.Get("X-Hydrate")
	if len(hydrateField) == 0 {
		return readResponse, nil
	}
	return NewHydrateResponse(readResponse, hydrateField)
}

func createRequest(req *http.Request) *http.Request {
	return &http.Request{
		Close:      false,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Body:       req.Body,
		Header:     req.Header,
		URL: &url.URL{
			Scheme:   "http",
			Host:     "127.0.0.1:4005",
			Path:     req.URL.Path,
			RawQuery: req.URL.RawQuery,
		},
	}
}

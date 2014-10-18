package hydrator

import (
	"bytes"
	"io"
	"net/http"
)

type Response interface {
	Header() http.Header
	Status() int
	Body() []byte
}

type NormalResponse struct {
	status int
	header http.Header
	body   []byte
}

// The body is read into a []byte because it's assumed
// that most of responses are going to be cached.
// However, this obviously won't always be the case.
// We should add an optimization to stream the bytes
// when we aren't going to cache (no cache control, post, put, ...)
func NewNormalResponse(response *http.Response) Response {
	var body []byte
	length := response.ContentLength
	if length > 0 {
		body = make([]byte, length)
		io.ReadFull(response.Body, body)
	} else if length == -1 {
		buffer := bytes.NewBuffer(make([]byte, 0, 16384))
		io.Copy(buffer, response.Body)
		body = buffer.Bytes()
		// if we're going to cache this request
		// we should consider trimming the buffer
	}
	response.Body.Close()
	return &NormalResponse{
		status: response.StatusCode,
		header: response.Header,
		body:   body,
	}
}

func (res *NormalResponse) Header() http.Header {
	return res.header
}

func (res *NormalResponse) Status() int {
	return res.status
}

func (res *NormalResponse) Body() []byte {
	return res.body
}

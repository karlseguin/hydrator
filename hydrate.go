package hydrator

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HydrateResponse struct {
	status int
	header http.Header
	parts  []Part
}

func NewHydrateResponse(res Response, fieldName string) (Response, error) {
	position := 0
	body := res.Body()
	needle := []byte("\"" + fieldName)
	parts := make([]Part, 0, 10)
	for {
		index := bytes.Index(body, needle)
		if index == -1 {
			parts = append(parts, LiteralPart(body[position:]))
			break
		}
		parts = append(parts, LiteralPart(body[position:index]))
		body = body[index:]
		start := bytes.IndexRune(body, '{')
		if start == -1 {
			//should probably at least log an error
			continue
		}
		end := bytes.IndexRune(body, '}')
		if end == -1 {
			//should probably at least log an error
			continue
		}
		end += 1

		var ref map[string]string
		if err := json.Unmarshal(body[start:end], &ref); err != nil {
			return nil, err
		}
		parts = append(parts, &ReferencePart{ref["id"], ref["type"]})
		body = body[end:]
	}

	return &HydrateResponse{
		status: res.Status(),
		header: res.Header(),
		parts:  parts,
	}, nil
}

func (res *HydrateResponse) Header() http.Header {
	return res.header
}

func (res *HydrateResponse) Status() int {
	return res.status
}

func (res *HydrateResponse) Body() []byte {
	// should use a pre-allocated buffer pool
	// to minimize the amount of allocations
	buffer := bytes.NewBuffer(make([]byte, 0, 16384))
	for _, part := range res.parts {
		buffer.Write(part.Render())
	}
	return buffer.Bytes()
}

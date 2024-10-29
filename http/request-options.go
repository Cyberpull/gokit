package http

import (
	"io"
	"net/http"
)

type RequestOptions struct {
	ExpectsJSON bool
	Header      http.Header
	Body        io.Reader
}

func (o *RequestOptions) mergeTo(req *http.Request) {
	if req == nil {
		return
	}

	// Merge Headers
	if o.Header != nil {
		for k, v := range o.Header {
			req.Header[k] = v
		}
	}

	// Merge Headers
	// Merge Headers
	// Merge Headers
	// Merge Headers
	// Merge Headers
	// Merge Headers
}

// ===================

func defaultRequestOptions(opts ...*RequestOptions) *RequestOptions {
	if len(opts) > 0 && opts[0] != nil {
		return opts[0]
	}

	return &RequestOptions{}
}

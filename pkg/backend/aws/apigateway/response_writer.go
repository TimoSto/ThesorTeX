package apigateway

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// struct that implements the api response writer
type responseWriter struct {
	statusCode int
	header     http.Header
	body       *bytes.Buffer

	wroteHeader bool
	snapHeader  http.Header // snapshot of HeaderMap at first Write
}

func (rw *responseWriter) response() (*events.APIGatewayProxyResponse, error) {
	response := &events.APIGatewayProxyResponse{}

	if rw.snapHeader != nil && len(rw.snapHeader) > 0 {
		response.MultiValueHeaders = rw.snapHeader
	}

	ct := rw.header.Get("Content-Type")
	// to fix font and image types
	if strings.HasPrefix(ct, "font/") || strings.HasPrefix(ct, "image/") {
		response.IsBase64Encoded = true
	}

	if rw.body.Len() > 0 {
		b, err := ioutil.ReadAll(rw.body)
		if err != nil {
			return nil, err
		}
		if response.IsBase64Encoded {
			enc := base64.StdEncoding.EncodeToString(b)
			response.Body = enc
		} else {
			response.Body = string(b)
		}
	}

	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}

	response.StatusCode = rw.statusCode

	return response, nil
}

// Header necessary to implement the responsewriter interface
func (rw *responseWriter) Header() http.Header {
	return rw.header
}

// Write necessary to implement the responsewriter interface
func (rw *responseWriter) Write(buf []byte) (int, error) {
	rw.writeHeader(buf)
	if rw.body != nil {
		rw.body.Write(buf)
	}
	return len(buf), nil
}

// writes header if it hasn't been called already and trys to detect the content-type if it is not set explicitly
func (rw *responseWriter) writeHeader(b []byte) {
	if rw.wroteHeader {
		return
	}

	m := rw.Header()
	_, hasType := m["Content-Type"]
	hasTE := m.Get("Transfer-Encoding") != ""
	if !hasType && !hasTE {
		m.Set("Content-Type", http.DetectContentType(b))
	}

	rw.WriteHeader(http.StatusOK)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	if rw.wroteHeader {
		return
	}
	rw.statusCode = statusCode
	rw.wroteHeader = true
	rw.snapHeader = cloneHeader(rw.header)
}

func cloneHeader(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, v := range h {
		v2 := make([]string, len(v))
		copy(v2, v)
		h2[k] = v2
	}
	return h2
}

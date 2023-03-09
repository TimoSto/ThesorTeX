package lambda

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//TODO: write differently

// implements https://godoc.org/net/http#ResponseWriter
type responseWriter struct {
	statusCode int
	header     http.Header
	body       *bytes.Buffer

	wroteHeader bool
	snapHeader  http.Header // snapshot of HeaderMap at first Write
}

func Start(handler http.Handler) {
	lambda.Start(ApplyHandlerToLambda(handler))
}

func ApplyHandlerToLambda(handler http.Handler) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fn := func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Info(fmt.Sprintf("Received APIGatewayRequest '%v'", request.RequestContext.RequestID))
		respw := &responseWriter{header: http.Header{}, body: &bytes.Buffer{}}
		req, err := newRequest(&request)
		if err != nil {
			log.Error(fmt.Sprintf("could not convert api gw request to http request: %v", err))
			resp := events.APIGatewayProxyResponse{Body: http.StatusText(http.StatusInternalServerError), StatusCode: http.StatusInternalServerError}
			return resp, nil
		}
		handler.ServeHTTP(respw, req.WithContext(ctx))
		resp, err := respw.response()
		if err != nil {
			log.Error(fmt.Sprintf("could not write http response to api gw: %v", err))
			resp := events.APIGatewayProxyResponse{Body: http.StatusText(http.StatusInternalServerError), StatusCode: http.StatusInternalServerError}
			return resp, nil
		}
		return *resp, nil
	}
	return fn
}

func newRequest(evt *events.APIGatewayProxyRequest) (*http.Request, error) {
	req := &http.Request{
		Method: mapMethod(evt),
		URL:    mapURL(evt),
		Header: *mapHeader(evt),
	}

	if req.URL.RawQuery != "" {
		req.RequestURI = req.URL.Path + "?" + req.URL.RawQuery
	} else {
		req.RequestURI = req.URL.Path
	}

	if evt.IsBase64Encoded {
		decodedString, err := base64.StdEncoding.DecodeString(evt.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Decoding of base64 body failed: %v", err))
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(decodedString))
	} else {
		req.Body = ioutil.NopCloser(strings.NewReader(evt.Body))
	}
	return req, nil
}

func mapMethod(e *events.APIGatewayProxyRequest) string {
	switch strings.ToUpper(e.HTTPMethod) {
	case "GET":
		return http.MethodGet
	case "POST":
		return http.MethodPost
	case "HEAD":
		return http.MethodHead
	case "PUT":
		return http.MethodPut
	case "PATCH":
		return http.MethodPatch
	case "DELETE":
		return http.MethodDelete
	case "CONNECT":
		return http.MethodConnect
	case "OPTIONS":
		return http.MethodOptions
	case "TRACE":
		return http.MethodTrace
	default:
		return ""
	}
}

func mapURL(e *events.APIGatewayProxyRequest) *url.URL {
	u := &url.URL{Path: e.Path}
	if e.QueryStringParameters != nil {
		values := make(url.Values)
		for key, value := range e.QueryStringParameters {
			values.Add(key, value)
		}
		u.RawQuery = values.Encode()
	}
	return u
}

func mapHeader(e *events.APIGatewayProxyRequest) *http.Header {
	result := &http.Header{}
	for k, v := range e.Headers {
		result.Add(k, v)
	}
	return result
}

func (rw *responseWriter) response() (*events.APIGatewayProxyResponse, error) {
	response := &events.APIGatewayProxyResponse{}

	if rw.snapHeader != nil && len(rw.snapHeader) > 0 {
		response.MultiValueHeaders = rw.snapHeader
	}

	ct := rw.header.Get("Content-Type")
	// to fix font types
	if strings.HasPrefix(ct, "font/") {
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

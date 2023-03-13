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

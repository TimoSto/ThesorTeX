package apigateway

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

func StartLambda(handler http.Handler) {
	lambda.Start(MapHandler(handler))
}

func MapHandler(handler http.Handler) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fn := func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Info(fmt.Sprintf("Received APIGatewayRequest '%v'", request.RequestContext.RequestID))

		respw := &responseWriter{header: http.Header{}, body: &bytes.Buffer{}}

		// convert api gw request to http request
		req, err := newRequest(&request)
		if err != nil {
			log.Error(fmt.Sprintf("could not convert api gw request to http request: %v", err))
			resp := events.APIGatewayProxyResponse{Body: http.StatusText(http.StatusInternalServerError), StatusCode: http.StatusInternalServerError}
			return resp, nil
		}

		// handle the http request
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
		Method: evt.HTTPMethod,
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

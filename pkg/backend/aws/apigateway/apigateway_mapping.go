package apigateway

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
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
	req := mapRequest(evt)

	if req.URL.RawQuery != "" {
		req.RequestURI = req.URL.Path + "?" + req.URL.RawQuery
	} else {
		req.RequestURI = req.URL.Path
	}

	fmt.Println("EVENT", evt.Body)

	// necessary for bin files like images and zips
	if evt.IsBase64Encoded {
		decodedString, err := base64.StdEncoding.DecodeString(evt.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Decoding of base64 body failed: %v", err))
		}
		req.Body = io.NopCloser(bytes.NewReader(decodedString))
	} else {
		req.Body = io.NopCloser(strings.NewReader(evt.Body))
	}
	return req, nil
}

func mapRequest(evt *events.APIGatewayProxyRequest) *http.Request {
	u := &url.URL{Path: evt.Path}
	if evt.QueryStringParameters != nil {
		queries := make(url.Values)
		for key, value := range evt.QueryStringParameters {
			queries.Add(key, value)
		}
		u.RawQuery = queries.Encode()
	}

	h := http.Header{}
	for k, v := range evt.Headers {
		h.Add(k, v)
	}

	return &http.Request{
		Method: evt.HTTPMethod,
		URL:    u,
		Header: h,
	}
}

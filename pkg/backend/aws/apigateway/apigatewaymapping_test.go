package apigateway

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestMapRequest(t *testing.T) {
	for _, c := range []struct {
		title string
		evt   events.APIGatewayProxyRequest
		exp   http.Request
	}{
		{
			title: "get only url",
			evt: events.APIGatewayProxyRequest{
				HTTPMethod: "GET",
				Path:       "/test/route1",
			},
			exp: http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/test/route1",
				},
			},
		},
		{
			title: "post one query param",
			evt: events.APIGatewayProxyRequest{
				HTTPMethod: "POST",
				Path:       "/test/route1",
				QueryStringParameters: map[string]string{
					"foo": "bar",
				},
			},
			exp: http.Request{
				Method: "POST",
				URL: &url.URL{
					Path:     "/test/route1",
					RawQuery: "foo=bar",
				},
			},
		},
		{
			title: "put three query params",
			evt: events.APIGatewayProxyRequest{
				HTTPMethod: "PUT",
				Path:       "/test/route1",
				QueryStringParameters: map[string]string{
					"foo":   "bar",
					"test":  "value",
					"query": "string",
				},
			},
			exp: http.Request{
				Method: "PUT",
				URL: &url.URL{
					Path:     "/test/route1",
					RawQuery: "foo=bar&query=string&test=value",
				},
			},
		},
		{
			title: "DELETE three query params with special chars",
			evt: events.APIGatewayProxyRequest{
				HTTPMethod: "DELETE",
				Path:       "/test/route1",
				QueryStringParameters: map[string]string{
					"foo":   "bar",
					"test":  "value",
					"query": "string value",
				},
			},
			exp: http.Request{
				Method: "DELETE",
				URL: &url.URL{
					Path:     "/test/route1",
					RawQuery: "foo=bar&query=string+value&test=value",
				},
			},
		},
	} {
		t.Run(c.title, func(t *testing.T) {
			got := mapRequest(&c.evt)

			if got.Method != c.exp.Method {
				t.Errorf("expected method %s but got %s", c.exp.Method, got.Method)
			}

			if diff := cmp.Diff(*got.URL, *c.exp.URL); diff != "" {
				t.Errorf("unexpected diff on url: %s", diff)
			}
		})
	}
}

func TestNewRequest(t *testing.T) {
	for _, c := range []struct {
		title   string
		evt     events.APIGatewayProxyRequest
		exp     http.Request
		expBody string
	}{
		{
			title: "get url without query params",
			evt: events.APIGatewayProxyRequest{
				HTTPMethod: http.MethodGet,
				Path:       "/test/route1",
			},
			exp: http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Path: "/test/route1",
				},
			},
		},
		{
			title: "get with one query param",
			evt: events.APIGatewayProxyRequest{
				HTTPMethod: http.MethodGet,
				Path:       "/test/route1",
				QueryStringParameters: map[string]string{
					"foo": "bar",
				},
			},
			exp: http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Path:     "/test/route1",
					RawQuery: "foo=bar",
				},
			},
		},
		{
			title: "post with three query params",
			evt: events.APIGatewayProxyRequest{
				HTTPMethod: http.MethodPost,
				Path:       "/test/route1",
				QueryStringParameters: map[string]string{
					"foo":   "bar",
					"test":  "value",
					"query": "string",
				},
				Body: "{\"foo\":\"bar\"}",
			},
			exp: http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path:     "/test/route1",
					RawQuery: "foo=bar&query=string&test=value",
				},
			},
			expBody: "{\"foo\":\"bar\"}",
		},
		{
			title: "post with three query params and special chars in body",
			evt: events.APIGatewayProxyRequest{
				HTTPMethod: http.MethodPost,
				Path:       "/test/route1",
				QueryStringParameters: map[string]string{
					"foo":   "bar",
					"test":  "value",
					"query": "string",
				},
				Body: "{\"f\\_oo\":\"bar\"}",
			},
			exp: http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path:     "/test/route1",
					RawQuery: "foo=bar&query=string&test=value",
				},
			},
			expBody: "{\"f\\_oo\":\"bar\"}",
		},
	} {
		t.Run(c.title, func(t *testing.T) {
			got, err := newRequest(&c.evt)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(got.URL, c.exp.URL); diff != "" {
				t.Errorf("unexpected diff in url: %s", diff)
			}

			if got.Method != c.exp.Method {
				t.Errorf("expected method %s but got %s", c.exp.Method, got.Method)
			}

			body, err := io.ReadAll(got.Body)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if string(body) != c.expBody {
				t.Errorf("expected body %s but got %s", c.expBody, string(body))
			}
		})
	}
}

package apigateway

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"
	"net/url"
	"testing"
)

func TestMapHeaders(t *testing.T) {
	agwReq := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"foo": "bar",
			"bar": "foo",
		},
	}

	headers := mapHeader(&agwReq)

	if headers.Get("foo") != "bar" {
		t.Errorf("expected value 'bar' for header 'foo'")
	}

	if headers.Get("bar") != "foo" {
		t.Errorf("expected value 'foo' for header 'bar'")
	}
}

func TestMapUrl(t *testing.T) {
	for _, c := range []struct {
		title string
		evt   events.APIGatewayProxyRequest
		exp   url.URL
	}{
		{
			title: "only url",
			evt: events.APIGatewayProxyRequest{
				Path: "/test/route1",
			},
			exp: url.URL{
				Path: "/test/route1",
			},
		},
		{
			title: "one query param",
			evt: events.APIGatewayProxyRequest{
				Path: "/test/route1",
				QueryStringParameters: map[string]string{
					"foo": "bar",
				},
			},
			exp: url.URL{
				Path:     "/test/route1",
				RawQuery: "foo=bar",
			},
		},
		{
			title: "three query params",
			evt: events.APIGatewayProxyRequest{
				Path: "/test/route1",
				QueryStringParameters: map[string]string{
					"foo":   "bar",
					"test":  "value",
					"query": "string",
				},
			},
			exp: url.URL{
				Path:     "/test/route1",
				RawQuery: "foo=bar&query=string&test=value",
			},
		},
		{
			title: "three query params with special chars",
			evt: events.APIGatewayProxyRequest{
				Path: "/test/route1",
				QueryStringParameters: map[string]string{
					"foo":   "bar",
					"test":  "value",
					"query": "string value",
				},
			},
			exp: url.URL{
				Path:     "/test/route1",
				RawQuery: "foo=bar&query=string+value&test=value",
			},
		},
	} {
		t.Run(c.title, func(t *testing.T) {
			got := mapURL(&c.evt)

			if diff := cmp.Diff(*got, c.exp); diff != "" {
				t.Errorf("unexpected diff: %s", diff)
			}
		})
	}
}

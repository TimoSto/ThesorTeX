package config_test

import (
	"errors"
	"github.com/TimoSto/ThesorTeX/services/router/internal/config"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetMappings(t *testing.T) {
	for _, tc := range []struct {
		name   string
		envs   map[string]string
		exp    []config.RouteMapping
		expErr error
	}{
		{
			name:   "nothing set",
			expErr: config.ErrNoRouteMappingFound,
		},
		{
			name: "apps set but not values",
			envs: map[string]string{
				"ROUTER_APPS": "APP-1, APP-2",
			},
			expErr: config.ErrNoRouteMappingFound,
		},
		{
			name: "apps set but only values for second app",
			envs: map[string]string{
				"ROUTER_APPS":              "APP-1, APP-2",
				"ROUTER_TARGET_URL__APP-2": "https://some-url.com",
				"ROUTER_ROUTE__APP-2":      "/app-2/",
			},
			exp: []config.RouteMapping{
				{
					TargetURL: "https://some-url.com",
					Route:     "/app-2/",
				},
			},
		},
		{
			name: "multiple apps set",
			envs: map[string]string{
				"ROUTER_APPS":              "APP-1, APP-2,APP-3 ",
				"ROUTER_TARGET_URL__APP-1": "https://some-url-1.com",
				"ROUTER_ROUTE__APP-1":      "/app-1/",
				"ROUTER_TARGET_URL__APP-2": "https://some-url-2.com",
				"ROUTER_ROUTE__APP-2":      "/app-2/",
				"ROUTER_TARGET_URL__APP-3": "https://some-url-3.com",
				"ROUTER_ROUTE__APP-3":      "/app-3/",
			},
			exp: []config.RouteMapping{
				{
					TargetURL: "https://some-url-1.com",
					Route:     "/app-1/",
				},
				{
					TargetURL: "https://some-url-2.com",
					Route:     "/app-2/",
				},
				{
					TargetURL: "https://some-url-3.com",
					Route:     "/app-3/",
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.envs {
				t.Setenv(k, v)
			}

			got, err := config.GetMappings()
			if !errors.Is(err, tc.expErr) {
				t.Errorf("expected error %v but got %v", tc.expErr, err)
			}

			if diff := cmp.Diff(got, tc.exp); diff != "" {
				t.Errorf("unexpected diff: %s", diff)
			}
		})
	}
}

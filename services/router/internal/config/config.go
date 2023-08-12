package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrNoRouteMappingFound = errors.New("no route mapping was found")
)

type RouteMapping struct {
	Route     string
	TargetURL string
}

func GetMappings() ([]RouteMapping, error) {
	appList := os.Getenv("ROUTER_APPS")
	apps := strings.Split(appList, ",")

	var mappings []RouteMapping

	for _, a := range apps {
		a = strings.TrimSpace(a)
		targetURL := os.Getenv(fmt.Sprintf("ROUTER_TARGET_URL__%s", a))
		route := os.Getenv(fmt.Sprintf("ROUTER_ROUTE__%s", a))
		if targetURL != "" && route != "" {
			mappings = append(mappings, RouteMapping{
				Route:     route,
				TargetURL: targetURL,
			})
		}
	}

	if len(mappings) == 0 {
		return nil, ErrNoRouteMappingFound
	}

	return mappings, nil
}

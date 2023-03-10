package parser

import (
	"testing"
)

func TestGetVersions(t *testing.T) {
	file := "APP=0.0.2\nTHESIS_TEMPLATE=1.0.1\nCV_TEMPLATE=1.0.0\n"

	version := GetVersions([]byte(file), "APP")

	expected := "0.0.2"

	if version != expected {
		t.Errorf("expected %v, got %v", expected, version)
	}
}

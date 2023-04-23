package semver

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseString(t *testing.T) {
	for _, tc := range []struct {
		str string
		ver Version
	}{
		{
			str: "v1.0.0",
			ver: Version{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
		},
		{
			str: "1.0.0",
			ver: Version{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
		},
		{
			str: "9.6.7",
			ver: Version{
				Major: 9,
				Minor: 6,
				Patch: 7,
			},
		},
	} {
		t.Run(tc.str, func(t *testing.T) {
			got, err := ParseString(tc.str)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(got, tc.ver); diff != "" {
				t.Errorf("unexpected diff: %s", diff)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	for _, tc := range []struct {
		title string
		v1    Version
		v2    Version
		exp   int
	}{
		{
			title: "v2 is larger",
			v1: Version{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
			v2: Version{
				Major: 1,
				Minor: 1,
				Patch: 0,
			},
			exp: -1,
		},
		{
			title: "v1 is larger",
			v1: Version{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			v2: Version{
				Major: 1,
				Minor: 1,
				Patch: 0,
			},
			exp: 1,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			if res := Compare(tc.v1, tc.v2); res != tc.exp {
				t.Errorf("expected %v, but got %v", tc.exp, res)
			}
		})
	}
}

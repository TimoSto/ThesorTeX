package parser

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAnalyseLine(t *testing.T) {
	tcs := []struct {
		line string
		exp  analyseLineResult
	}{
		{
			line: "",
			exp: analyseLineResult{
				Type:    TypeEmpty,
				Content: "",
			},
		},
		{
			line: "\n",
			exp: analyseLineResult{
				Type:    TypeEmpty,
				Content: "",
			},
		},
		{
			line: "foo bar",
			exp: analyseLineResult{
				Type:    TypeText,
				Content: "foo bar",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.line, func(t *testing.T) {
			result := analyseLine(tc.line)

			if diff := cmp.Diff(result, tc.exp); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestSplitLineIntoElements(t *testing.T) {
	tcs := []struct {
		line string
		exp  []element
	}{
		{
			line: "foo bar developer",
			exp: []element{
				{
					Style:   StylePlain,
					Content: "foo bar developer",
				},
			},
		},
		{
			line: "foo *bar* developer",
			exp: []element{
				{
					Style:   StylePlain,
					Content: "foo ",
				},
				{
					Style:   StyleItalic,
					Content: "bar",
				},
				{
					Style:   StylePlain,
					Content: " developer",
				},
			},
		},
		{
			line: "**foo** bar developer",
			exp: []element{
				{
					Style:   StyleBold,
					Content: "foo",
				},
				{
					Style:   StylePlain,
					Content: " bar developer",
				},
			},
		},
		{
			line: "**foo** hallo *bar* developer",
			exp: []element{
				{
					Style:   StyleBold,
					Content: "foo",
				},
				{
					Style:   StylePlain,
					Content: " hallo ",
				},
				{
					Style:   StyleItalic,
					Content: "bar",
				},
				{
					Style:   StylePlain,
					Content: " developer",
				},
			},
		},
		{
			line: "**foo** *bar* developer",
			exp: []element{
				{
					Style:   StyleBold,
					Content: "foo",
				},
				{
					Style:   StylePlain,
					Content: " ",
				},
				{
					Style:   StyleItalic,
					Content: "bar",
				},
				{
					Style:   StylePlain,
					Content: " developer",
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.line, func(t *testing.T) {
			result := splitLineIntoElements(tc.line)

			if diff := cmp.Diff(tc.exp, result); diff != "" {
				fmt.Print(result)
				t.Errorf("%s", diff)
			}
		})
	}
}

var simpleExpectedBody = []DocBody{
	{
		Groups: []group{
			{
				Type: "TEXT",
				Elements: []element{
					{
						Style:   "PLAIN",
						Content: "Some content",
					},
				},
			},
		},
	},
	{
		Groups: []group{
			{
				Type: "TEXT",
				Elements: []element{
					{
						Style:   "PLAIN",
						Content: "Simple content ",
					},
					{
						Style:   "ITALIC",
						Content: "with",
					},
					{
						Style:   "PLAIN",
						Content: " ",
					},
					{
						Style:   "BOLD",
						Content: "some",
					},
					{
						Style:   "PLAIN",
						Content: " ",
					},
					{
						Style:   "ITALIC-BOLD",
						Content: "styling",
					},
				},
			},
		},
	},
	{
		Groups: []group{
			{
				Type: "TEXT",
				Elements: []element{
					{
						Style:   "PLAIN",
						Content: "Simple",
					},
				},
			},
			{
				Type: "TEXT",
				Elements: []element{
					{
						Style:   "PLAIN",
						Content: "content ",
					},
					{
						Style:   "ITALIC",
						Content: "with",
					},
					{
						Style:   "PLAIN",
						Content: " ",
					},
					{
						Style:   "BOLD",
						Content: "some",
					},
					{
						Style:   "PLAIN",
						Content: " ",
					},
					{
						Style:   "ITALIC-BOLD",
						Content: "styling",
					},
				},
			},
		},
	},
}

func testParseDocBody(t *testing.T) {
	for i, s := range simpleExpected {
		t.Run(s.Title, func(t *testing.T) {
			result := parseDocBody(s)

			expected := simpleExpectedBody[i]

			if diff := cmp.Diff(result, expected); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

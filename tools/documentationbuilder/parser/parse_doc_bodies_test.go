package parser

import (
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
		title string
		line  string
		exp   []element
	}{
		{
			title: "just plain",
			line:  "foo bar developer",
			exp: []element{
				{
					Style:   StylePlain,
					Content: "foo bar developer",
				},
			},
		},
		{
			title: "italic in middle",
			line:  "foo *bar* developer",
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
			title: "bold in front",
			line:  "**foo** bar developer",
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
			title: "italic and bold texts separated by word",
			line:  "**foo** hallo *bar* developer",
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
			title: "italic and bold separated by space",
			line:  "**foo** *bar* developer",
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
		{
			title: "italic, bold and italic-bold",
			line:  "**foo** *bar* ***de***veloper",
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
					Content: " ",
				},
				{
					Style:   StyleItalicAndBold,
					Content: "de",
				},
				{
					Style:   StylePlain,
					Content: "veloper",
				},
			},
		},
		{
			title: "italic, bold and italic-bold twice at end",
			line:  "**foo** *bar* ***de***velop***er***",
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
					Content: " ",
				},
				{
					Style:   StyleItalicAndBold,
					Content: "de",
				},
				{
					Style:   StylePlain,
					Content: "velop",
				},
				{
					Style:   StyleItalicAndBold,
					Content: "er",
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.line, func(t *testing.T) {
			result := splitLineIntoElements(tc.line)

			if diff := cmp.Diff(tc.exp, result); diff != "" {
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

func TestParseDocBody(t *testing.T) {
	for i, s := range simpleExpected {
		t.Run(s.Title, func(t *testing.T) {
			result := parseDocBody(s)

			expected := simpleExpectedBody[i]

			if diff := cmp.Diff(expected, result); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestParseDocBodies(t *testing.T) {

	result := ParseDocBodies(simpleExpected)

	if diff := cmp.Diff(simpleExpectedBody, result); diff != "" {
		t.Errorf("%s", diff)
	}
}

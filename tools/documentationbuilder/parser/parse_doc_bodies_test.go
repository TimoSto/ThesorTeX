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
		{
			line: "![some img](./some/images.png) ",
			exp: analyseLineResult{
				Type:    TypeImage,
				Content: "![some img](./some/images.png)",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.line, func(t *testing.T) {
			result := analyseLine(tc.line, false)

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
		{
			title: "with link",
			line:  "hallo [An Internal Link](/guides/content/editing-an-existing-page)",
			exp: []element{
				{
					Style:   StylePlain,
					Content: "hallo ",
				},
				{
					Style:   LinkTitle,
					Content: "An Internal Link",
				},
				{
					Style:   LinkHref,
					Content: "/guides/content/editing-an-existing-page",
				},
			},
		},
		{
			title: "with link and bold and italic",
			line:  "hallo *hallo* [An Internal Link](/guides/content/editing-an-existing-page) **hallo**",
			exp: []element{
				{
					Style:   StylePlain,
					Content: "hallo ",
				},
				{
					Style:   StyleItalic,
					Content: "hallo",
				},
				{
					Style:   StylePlain,
					Content: " ",
				},
				{
					Style:   LinkTitle,
					Content: "An Internal Link",
				},
				{
					Style:   LinkHref,
					Content: "/guides/content/editing-an-existing-page",
				},
				{
					Style:   StylePlain,
					Content: " ",
				},
				{
					Style:   StyleBold,
					Content: "hallo",
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
		Title: "Doc 1",
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
		Title: "Doc 2 styled",
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
		Title: "Extra spacing",
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

var withCodeExpectedBody = []DocBody{
	{
		Title: "Doc 1",
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
			{
				Type: "CODE",
				Elements: []element{
					{
						Style:   "",
						Content: "\\testcommdand{}",
					},
					{
						Style:   "",
						Content: "hallo",
					},
				},
			},
			{
				Type: "TEXT",
				Elements: []element{
					{
						Style:   "PLAIN",
						Content: "test",
					},
				},
			},
		},
	},
}

var withImageExpectedBody = []DocBody{
	{
		Title: "Doc 1",
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
			{
				Type: "IMAGE",
				Elements: []element{
					{
						Style:   "",
						Content: "./images/img1.png",
					},
					{
						Style:   "",
						Content: "test image",
					},
				},
			},
			{
				Type: "TEXT",
				Elements: []element{
					{
						Style:   "PLAIN",
						Content: "test",
					},
				},
			},
		},
	},
}

var withLinkExpectedBody = []DocBody{
	{
		Title: "Doc 1",
		Groups: []group{
			{
				Type: "TEXT",
				Elements: []element{
					{
						Style:   "PLAIN",
						Content: "Some content ",
					},
					{
						Content: "link",
						Style:   "LINK_TITLE",
					},
					{
						Content: "https://funny.address",
						Style:   "LINK_HREF",
					},
					{
						Style:   "PLAIN",
						Content: " test",
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

func TestParseDocBodyWithCode(t *testing.T) {
	for i, s := range withCodeExpected {
		t.Run(s.Title, func(t *testing.T) {
			result := parseDocBody(s)

			expected := withCodeExpectedBody[i]

			if diff := cmp.Diff(expected, result); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestParseDocBodyWithImage(t *testing.T) {
	for i, s := range withImageExpected {
		t.Run(s.Title, func(t *testing.T) {
			result := parseDocBody(s)

			expected := withImageExpectedBody[i]

			if diff := cmp.Diff(expected, result); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestParseDocBodyWithLink(t *testing.T) {
	for i, s := range withLinkExpected {
		t.Run(s.Title, func(t *testing.T) {
			result := parseDocBody(s)

			expected := withLinkExpectedBody[i]

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

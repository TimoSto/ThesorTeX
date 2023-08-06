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
		{
			line: "[^11]:hallo",
			exp: analyseLineResult{
				Type:    TypeFootnoteRef,
				Content: "hallo",
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
			title: "just plain with footnote",
			line:  "foo bar[^1] developer",
			exp: []element{
				{
					Style:   StylePlain,
					Content: "foo bar",
				},
				{
					Style:   Footnote,
					Content: "1",
				},
				{
					Style:   StylePlain,
					Content: " developer",
				},
			},
		},
		{
			title: "just plain with footnote at end",
			line:  "foo bar developer[^12].",
			exp: []element{
				{
					Style:   StylePlain,
					Content: "foo bar developer",
				},
				{
					Style:   Footnote,
					Content: "12",
				},
				{
					Style:   StylePlain,
					Content: ".",
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.title, func(t *testing.T) {
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

var withLinksExpectedBody = []DocBody{
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
						Style:   "LINK-TITLE",
						Content: "link 1",
					},
					{
						Style:   "LINK-HREF",
						Content: "https://localhost.test1",
					},
					{
						Style:   "PLAIN",
						Content: " and ",
					},
					{
						Style:   "LINK-TITLE",
						Content: "https://localhost.test2",
					},
					{
						Style:   "LINK-HREF",
						Content: "https://localhost.test2",
					},
					{
						Style:   "PLAIN",
						Content: ".",
					},
				},
			},
		},
	},
}

var withFootnoteExpectedBody = []DocBody{
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
					{
						Style:   "PLAIN",
						Content: " and footnote.",
					},
					{
						Style:   "FOOTNOTE",
						Content: "1",
					},
				},
			},
		},
		Footnotes: map[int][]element{
			1: []element{
				{
					Style:   "PLAIN",
					Content: "footnote content",
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

func TestParseDocBodyWithLinks(t *testing.T) {
	for i, s := range withLinksExpected {
		t.Run(s.Title, func(t *testing.T) {
			result := parseDocBody(s)

			expected := withLinksExpectedBody[i]

			if diff := cmp.Diff(expected, result); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestParseDocBodyWithFootnote(t *testing.T) {
	for i, s := range withFootnoteExpected {
		t.Run(s.Title, func(t *testing.T) {
			result := parseDocBody(s)

			expected := withFootnoteExpectedBody[i]

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

package bib_categories

import (
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func TestGeneratePrintCommands(t *testing.T) {
	cases := []struct {
		title         string
		categories    []database.BibCategory
		expectedBib   string
		expectedCites string
	}{
		{
			title: "one empty category",
			categories: []database.BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields:         nil,
					CiteFields:     nil,
				},
			},
			expectedBib: `\newcommand{\printtest}[0]{%
	\hangindent=\bibparindent%
	\parindent 0pt%
	\hangafter=1%
	%
	\\%
	\vspace{-12pt}%

}
`,
			expectedCites: `\newcommand{\citetest}[0]{%
	%
}`,
		},
		{
			title: "one category with one bib and citefield",
			categories: []database.BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []database.Field{
						{
							Field:            "f1",
							Italic:           true,
							Prefix:           "in: ",
							Suffix:           ". ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
					CiteFields: []database.Field{
						{
							Field:            "f2",
							Italic:           true,
							Prefix:           "",
							Suffix:           ", ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
				},
			},
			expectedBib: `\newcommand{\printtest}[0]{%
	\hangindent=\bibparindent%
	\parindent 0pt%
	\hangafter=1%
	in: \textit{\argI}. %
	\\%
	\vspace{-12pt}%

}
`,
			expectedCites: `\newcommand{\citetest}[0]{%
	\textit{\argII}, %
}`,
		},
		{
			title: "one category with the same bib and citefield",
			categories: []database.BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []database.Field{
						{
							Field:            "f1",
							Italic:           true,
							Prefix:           "in: ",
							Suffix:           ". ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
					CiteFields: []database.Field{
						{
							Field:            "f1",
							Italic:           true,
							Prefix:           "",
							Suffix:           ", ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
				},
			},
			expectedBib: `\newcommand{\printtest}[0]{%
	\hangindent=\bibparindent%
	\parindent 0pt%
	\hangafter=1%
	in: \textit{\argI}. %
	\\%
	\vspace{-12pt}%

}
`,
			expectedCites: `\newcommand{\citetest}[0]{%
	\textit{\argI}, %
}`,
		},
		{
			title: "one category with overlapping bib and citefields but different styles",
			categories: []database.BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []database.Field{
						{
							Field:            "f1",
							Italic:           true,
							Prefix:           "in: ",
							Suffix:           ". ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
						{
							Field:            "f2",
							Italic:           true,
							Prefix:           "(",
							Suffix:           ")",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
					CiteFields: []database.Field{
						{
							Field:            "f3",
							Italic:           true,
							Prefix:           "",
							Suffix:           ", ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
						{
							Field:            "f2",
							Italic:           true,
							Prefix:           "Jahr: ",
							Suffix:           "",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
				},
			},
			expectedBib: `\newcommand{\printtest}[0]{%
	\hangindent=\bibparindent%
	\parindent 0pt%
	\hangafter=1%
	in: \textit{\argI}. (\textit{\argII})%
	\\%
	\vspace{-12pt}%

}
`,
			expectedCites: `\newcommand{\citetest}[0]{%
	\textit{\argIII}, Jahr: \textit{\argII}%
}`,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			bib, cites := GeneratePrintCommands(c.categories)
			if bib != c.expectedBib {
				t.Errorf("got bib prints %s, want: %v", bib, c.expectedBib)
			}
			if cites != c.expectedCites {
				t.Errorf("got cite prints %s, want: %v", cites, c.expectedCites)
			}
		})
	}
}

func TestGenerateCommand(t *testing.T) {
	cases := []struct {
		title        string
		fields       []database.Field
		searchfields []database.Field
		expected     string
	}{
		{
			title: "one plain, one with pre- and suffix, one italic with pre- and suffix",
			fields: []database.Field{
				{
					Field:            "f1",
					Italic:           false,
					Prefix:           "",
					Suffix:           "",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f2",
					Italic:           false,
					Prefix:           "(",
					Suffix:           "). ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f3",
					Italic:           true,
					Prefix:           "in: ",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: `\argI(\argII). in: \textit{\argIII}, `,
		},
		{
			title: "one italic with suffix, one with pre- and suffix, one italic",
			fields: []database.Field{
				{
					Field:            "f1",
					Italic:           true,
					Prefix:           "",
					Suffix:           " - ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f2",
					Italic:           false,
					Prefix:           "(",
					Suffix:           "). ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f3",
					Italic:           true,
					Prefix:           "",
					Suffix:           "",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: `\textit{\argI} - (\argII). \textit{\argIII}`,
		},
		{
			title: "one plain, one with pre- and suffix, one italic with pre- and suffix with searchfields (less than fields)",
			fields: []database.Field{
				{
					Field:            "f1",
					Italic:           false,
					Prefix:           "",
					Suffix:           "",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f2",
					Italic:           false,
					Prefix:           "(",
					Suffix:           "). ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f3",
					Italic:           true,
					Prefix:           "in: ",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			searchfields: []database.Field{
				{
					Field:            "f2",
					Italic:           true,
					Prefix:           "",
					Suffix:           " - ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: `\argII(\argI). in: \textit{\argIII}, `,
		},
		{
			title: "one plain, one with pre- and suffix, one italic with pre- and suffix with searchfields (equal num than fields)",
			fields: []database.Field{
				{
					Field:            "f1",
					Italic:           false,
					Prefix:           "",
					Suffix:           "",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f2",
					Italic:           false,
					Prefix:           "(",
					Suffix:           "). ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f4",
					Italic:           true,
					Prefix:           "in: ",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			searchfields: []database.Field{
				{
					Field:            "f1",
					Italic:           true,
					Prefix:           "",
					Suffix:           " - ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f2",
					Italic:           true,
					Prefix:           "",
					Suffix:           " - ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f3",
					Italic:           true,
					Prefix:           "",
					Suffix:           " - ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: `\argI(\argII). in: \textit{\argIV}, `,
		},
		{
			title: "one plain, one with pre- and suffix, one italic with pre- and suffix with searchfields (less than fields)",
			fields: []database.Field{
				{
					Field:            "f1",
					Italic:           false,
					Prefix:           "",
					Suffix:           " ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f4",
					Italic:           true,
					Prefix:           "in: ",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			searchfields: []database.Field{
				{
					Field:            "f1",
					Italic:           true,
					Prefix:           "",
					Suffix:           " - ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f2",
					Italic:           true,
					Prefix:           "",
					Suffix:           " - ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "f3",
					Italic:           true,
					Prefix:           "",
					Suffix:           " - ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: `\argI in: \textit{\argIV}, `,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if got, want := GenerateCommand(c.fields, c.searchfields), c.expected; got != want {
				t.Errorf("got command %s, want: %s", got, want)
			}
		})
	}
}

func TestGenerateAssignment(t *testing.T) {
	categories := []database.BibCategory{
		{
			Name: "test1",
		},
		{
			Name: "TestCamelCase",
		},
	}
	expectedBib := "\t\t{test1}{\\printtest1}\n\t\t{TestCamelCase}{\\printTestCamelCase}\n"

	gotBib := GenerateAssignment(categories, "\\print")

	if gotBib != expectedBib {
		t.Errorf("got %s, want %s", gotBib, expectedBib)
	}

	expectedCite := "\t\t{test1}{\\citetest1}\n\t\t{TestCamelCase}{\\citeTestCamelCase}\n"

	gotCite := GenerateAssignment(categories, "\\cite")

	if gotCite != expectedCite {
		t.Errorf("got %s, want %s", gotCite, expectedCite)
	}
}

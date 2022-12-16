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
		title    string
		fields   []database.Field
		expected string
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
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if got, want := GenerateCommand(c.fields), c.expected; got != want {
				t.Errorf("got command %s, want: %s", got, want)
			}
		})
	}
}

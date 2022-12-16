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

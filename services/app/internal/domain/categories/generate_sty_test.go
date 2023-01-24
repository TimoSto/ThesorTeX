package categories

import (
	"testing"
)

func TestGeneratePrintCommands(t *testing.T) {
	cases := []struct {
		title         string
		categories    []Category
		expectedBib   string
		expectedCites string
	}{
		{
			title: "one empty category",
			categories: []Category{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilter:   nil,
					BibFields:      nil,
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
}
`,
		},
		{
			title: "one category with one bib and citefield",
			categories: []Category{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilter:   nil,
					BibFields: []field{
						{
							Name: "f1",
							Format: format{
								Italic:       true,
								Prefix:       "in: ",
								Suffix:       ". ",
								Preformatted: false,
							},
							CitaviMapping: nil,
						},
					},
					CiteFields: []field{
						{
							Name: "f2",
							Format: format{
								Italic:       true,
								Prefix:       "",
								Suffix:       ", ",
								Preformatted: false,
							},
							CitaviMapping: nil,
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
}
`,
		},
		{
			title: "one category with the same bib and citefield",
			categories: []Category{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilter:   nil,
					BibFields: []field{
						{
							Name: "f1",
							Format: format{
								Italic:       true,
								Prefix:       "in: ",
								Suffix:       ". ",
								Preformatted: false,
							},
							CitaviMapping: nil,
						},
					},
					CiteFields: []field{
						{
							Name: "f1",
							Format: format{
								Italic:       true,
								Prefix:       "",
								Suffix:       ", ",
								Preformatted: false,
							},
							CitaviMapping: nil,
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
}
`,
		},
		{
			title: "one category with overlapping bib and citefields but different styles",
			categories: []Category{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilter:   nil,
					BibFields: []field{
						{
							Name: "f1",
							Format: format{
								Italic:       true,
								Prefix:       "in: ",
								Suffix:       ". ",
								Preformatted: false,
							},
							CitaviMapping: nil,
						},
						{
							Name: "f2",
							Format: format{
								Italic:       true,
								Prefix:       "(",
								Suffix:       ")",
								Preformatted: false,
							},
							CitaviMapping: nil,
						},
					},
					CiteFields: []field{
						{
							Name: "f3",
							Format: format{
								Italic:       true,
								Prefix:       "",
								Suffix:       ", ",
								Preformatted: false,
							},
							CitaviMapping: nil,
						},
						{
							Name: "f2",
							Format: format{
								Italic:       true,
								Prefix:       "Jahr: ",
								Suffix:       "",
								Preformatted: false,
							},
							CitaviMapping: nil,
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
}
`,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			bib, cites := GeneratePrintCommands(c.categories)
			if bib != c.expectedBib {
				t.Errorf("got bib prints %s, want: %v", bib, c.expectedBib)
			}
			if cites != c.expectedCites {
				t.Errorf("got cite prints \n%s\n, want: \n%v\n", cites, c.expectedCites)
			}
		})
	}
}

func TestGenerateCommand(t *testing.T) {
	cases := []struct {
		title        string
		fields       []field
		searchfields []field
		expected     string
	}{
		{
			title: "one plain, one with pre- and suffix, one italic with pre- and suffix",
			fields: []field{
				{
					Name: "f1",
					Format: format{
						Italic:       false,
						Prefix:       "",
						Suffix:       "",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f2",
					Format: format{
						Italic:       false,
						Prefix:       "(",
						Suffix:       "). ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f3",
					Format: format{
						Italic:       true,
						Prefix:       "in: ",
						Suffix:       ", ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
			},
			expected: `\argI(\argII). in: \textit{\argIII}, `,
		},
		{
			title: "one italic with suffix, one with pre- and suffix, one italic",
			fields: []field{
				{
					Name: "f1",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       " - ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f2",
					Format: format{
						Italic:       false,
						Prefix:       "(",
						Suffix:       "). ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f3",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       "",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
			},
			expected: `\textit{\argI} - (\argII). \textit{\argIII}`,
		},
		{
			title: "one plain, one with pre- and suffix, one italic with pre- and suffix with searchfields (less than fields)",
			fields: []field{
				{
					Name: "f1",
					Format: format{
						Italic:       false,
						Prefix:       "",
						Suffix:       "",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f2",
					Format: format{
						Italic:       false,
						Prefix:       "(",
						Suffix:       "). ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f3",
					Format: format{
						Italic:       true,
						Prefix:       "in: ",
						Suffix:       ", ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
			},
			searchfields: []field{
				{
					Name: "f2",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       " - ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
			},
			expected: `\argII(\argI). in: \textit{\argIII}, `,
		},
		{
			title: "one plain, one with pre- and suffix, one italic with pre- and suffix with searchfields (equal num than fields)",
			fields: []field{
				{
					Name: "f1",
					Format: format{
						Italic:       false,
						Prefix:       "",
						Suffix:       "",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f2",
					Format: format{
						Italic:       false,
						Prefix:       "(",
						Suffix:       "). ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f4",
					Format: format{
						Italic:       true,
						Prefix:       "in: ",
						Suffix:       ", ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
			},
			searchfields: []field{
				{
					Name: "f1",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       " - ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f2",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       " - ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f3",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       " - ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
			},
			expected: `\argI(\argII). in: \textit{\argIV}, `,
		},
		{
			title: "one plain, one with pre- and suffix, one italic with pre- and suffix with searchfields (less than fields)",
			fields: []field{
				{
					Name: "f1",
					Format: format{
						Italic:       false,
						Prefix:       "",
						Suffix:       " ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f4",
					Format: format{
						Italic:       true,
						Prefix:       "in: ",
						Suffix:       ", ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
			},
			searchfields: []field{
				{
					Name: "f1",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       " - ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f2",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       " - ",
						Preformatted: false,
					},
					CitaviMapping: nil,
				},
				{
					Name: "f3",
					Format: format{
						Italic:       true,
						Prefix:       "",
						Suffix:       " - ",
						Preformatted: false,
					},
					CitaviMapping: nil,
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
	categories := []Category{
		{
			Name: "test1",
		},
		{
			Name: "TestCamelCase",
		},
	}

	expectedBib := "\t\t{test1}{\\printtest1}%\n\t\t{TestCamelCase}{\\printTestCamelCase}%\n"
	expectedCite := "\t\t{test1}{\\citetest1}%\n\t\t{TestCamelCase}{\\citeTestCamelCase}%\n"

	gotBib, gotCite := GenerateAssignment(categories)

	if gotBib != expectedBib {
		t.Errorf("got %s, want %s", gotBib, expectedBib)
	}

	if gotCite != expectedCite {
		t.Errorf("got %s, want %s", gotCite, expectedCite)
	}
}

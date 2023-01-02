package bib_categories

import (
	"regexp"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database/fake_store"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

func TestSaveCategoriesToSty(t *testing.T) {
	fake := fake_store.Store{}

	categories := []BibCategory{
		{
			Name: "test1",
			Fields: []Field{
				{
					Field:  "t11",
					Suffix: " ",
				},
				{
					Field:  "t12",
					Prefix: "(",
					Suffix: "), ",
					Italic: true,
				},
				{
					Field:  "t13",
					Suffix: "",
					Prefix: "",
				},
			},
			CiteFields: []Field{
				{
					Field:  "t14",
					Suffix: ", ",
				},
				{
					Field:  "t12",
					Suffix: ")",
					Prefix: "(",
					Italic: true,
				},
			},
		},
		{
			Name: "test2",
			Fields: []Field{
				{
					Field:  "t21",
					Suffix: ", ",
				},
				{
					Field:  "t22",
					Suffix: "",
					Prefix: "",
				},
			},
			CiteFields: []Field{
				{
					Field:  "t23",
					Suffix: ", ",
				},
				{
					Field:  "t22",
					Suffix: ")",
					Prefix: "(",
					Italic: true,
				},
			},
		},
	}

	tmpl, _ := project_template.ProjectTemplate.ReadFile("template" + styFilePath)

	fake.WriteFileInProject("", styFilePath, tmpl)

	err := SaveCategoriesToSty(&fake, "", categories)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	result, _ := fake.ReadFileInProject("", styFilePath)

	expectedBibCommands := `% begin bib commands
\newcommand{\printtest1}[0]{%
	\hangindent=\bibparindent%
	\parindent 0pt%
	\hangafter=1%
	\argI (\textit{\argII}), \argIII%
	\\%
	\vspace{-12pt}%

}
\newcommand{\printtest2}[0]{%
	\hangindent=\bibparindent%
	\parindent 0pt%
	\hangafter=1%
	\argI, \argII%
	\\%
	\vspace{-12pt}%

}
% end bib commands`
	m1 := regexp.MustCompile(`(?s)% begin bib commands(.*?)% end bib commands`)
	bibCommandsResult := m1.FindStringSubmatch(string(result))

	if len(bibCommandsResult) == 0 {
		t.Errorf("did not find bib commands")
		return
	}

	if bibCommandsResult[0] != expectedBibCommands {
		t.Errorf("got %v want %v", bibCommandsResult[0], expectedBibCommands)
	}

	expectedCiteCommands := `% begin cite commands
\newcommand{\citetest1}[0]{%
	\argIV, (\textit{\argII})%
}
\newcommand{\citetest2}[0]{%
	\argIII, (\textit{\argII})%
}
% end cite commands`

	m2 := regexp.MustCompile(`(?s)% begin cite commands(.*?)% end cite commands`)
	citeCommandsResult := m2.FindStringSubmatch(string(result))

	expectedBibAssignments := `% begin bib assignment
		{test1}{\printtest1}%
		{test2}{\printtest2}%
	% end bib assignment`

	if len(citeCommandsResult) == 0 {
		t.Errorf("did not find cite commands")
	}

	if citeCommandsResult[0] != expectedCiteCommands {
		t.Errorf("got %v want %v", citeCommandsResult[0], expectedCiteCommands)
	}

	m3 := regexp.MustCompile(`(?s)% begin bib assignment(.*?)% end bib assignment`)
	bibAssignResult := m3.FindStringSubmatch(string(result))

	if len(bibAssignResult) == 0 {
		t.Errorf("did not find bib assignment")
	}

	if bibAssignResult[0] != expectedBibAssignments {
		t.Errorf("got \n%v\n want \n%v\n", bibAssignResult[0], expectedBibAssignments)
	}

	expectedCiteAssignments := `% begin cite assignment
		{test1}{\citetest1}%
		{test2}{\citetest2}%
	% end cite assignment`

	m4 := regexp.MustCompile(`(?s)% begin cite assignment(.*?)% end cite assignment`)
	citeAssignResult := m4.FindStringSubmatch(string(result))

	if len(citeAssignResult) == 0 {
		t.Errorf("did not find cite assignment")
	}

	if citeAssignResult[0] != expectedCiteAssignments {
		t.Errorf("got %v want %v", citeAssignResult[0], expectedCiteAssignments)
	}

	//fileRes, _ := fake.GetBibliographySty("")

	//m1 := regexp.MustCompile(`(?s)% begin bib commands(.*?)% end bib commands`)
	//bibCommands := m1.FindStringSubmatch(fileRes)

	//m2 := regexp.MustCompile(`(?s)% begin cite commands(.*?)% end cite commands`)
	//newFile = m2.ReplaceAllString(newFile, citeCommands)

	//newFile = m3.ReplaceAllString(newFile, bibAssignments)
	//m4 := regexp.MustCompile(`(?s)% begin cite assignments(.*?)% end cite assignments`)
	//newFile = m4.ReplaceAllString(newFile, citeAssignments)
}

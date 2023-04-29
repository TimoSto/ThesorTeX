package categories

import (
	"regexp"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/pkg/backend/project_template"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
)

func TestSaveCategoriesToSty(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "",
	}

	categories := []Category{
		{
			Name: "test1",
			BibFields: []field{
				{
					Name: "t11",
					Format: format{
						Suffix: " ",
					},
				},
				{
					Name: "t12",
					Format: format{
						Prefix: "(",
						Suffix: "), ",
						Italic: true,
					},
				},
				{
					Name: "t13",
					Format: format{
						Suffix: "",
						Prefix: "",
					},
				},
			},
			CiteFields: []field{
				{
					Name: "t14",
					Format: format{
						Suffix: ", ",
					},
				},
				{
					Name: "t12",
					Format: format{
						Suffix: ")",
						Prefix: "(",
						Italic: true,
					},
				},
			},
		},
		{
			Name: "test2",
			BibFields: []field{
				{
					Name: "t21",
					Format: format{
						Suffix: ", ",
					},
				},
				{
					Name: "t22",
					Format: format{
						Suffix: "",
						Prefix: "",
					},
				},
			},
			CiteFields: []field{
				{
					Name: "t23",
					Format: format{
						Suffix: ", ",
					},
				},
				{
					Name: "t22",
					Format: format{
						Suffix: ")",
						Prefix: "(",
						Italic: true,
					},
				},
			},
		},
	}

	tmpl, _ := project_template.ProjectTemplate.ReadFile("template/" + styFilePath)

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "", styFilePath), tmpl)

	err := SaveCategoriesToSty(&fs, cfg, "", categories)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	result, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "", styFilePath))

	expectedBibCommands := `% begin bib commands
\newcommand{\printtest1}[0]{%
	\hangindent=\bibparindent%
	\parindent 0pt%
	\hangafter=1%
	\argI{ }(\textit{\argII}),{ }\argIII%
	\\%
	\vspace{-12pt}%

}
\newcommand{\printtest2}[0]{%
	\hangindent=\bibparindent%
	\parindent 0pt%
	\hangafter=1%
	\argI,{ }\argII%
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
	\argIV,{ }(\textit{\argII})%
}
\newcommand{\citetest2}[0]{%
	\argIII,{ }(\textit{\argII})%
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

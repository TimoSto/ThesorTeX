package categories

import (
	"regexp"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
)

const (
	styFilePath = "styPackages/bibliography.sty"
)

func SaveCategoriesToSty(fs filesystem.FileSystem, cfg config.Config, project string, categories []Category) error {
	file, err := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, project, styFilePath))
	if err != nil {
		return err
	}

	bibCommands, citeCommands := GeneratePrintCommands(categories)
	bibCommands = "% begin bib commands\n" + bibCommands + "% end bib commands"
	citeCommands = "% begin cite commands\n" + citeCommands + "% end cite commands"

	bibAssignments, citeAssignments := GenerateAssignment(categories)
	bibAssignments = "% begin bib assignment\n" + bibAssignments + "\t% end bib assignment"
	citeAssignments = "% begin cite assignment\n" + citeAssignments + "\t% end cite assignment"

	m1 := regexp.MustCompile(`(?s)% begin bib commands(.*?)% end bib commands`)
	newFile := m1.ReplaceAllString(string(file), bibCommands)
	m2 := regexp.MustCompile(`(?s)% begin cite commands(.*?)% end cite commands`)
	newFile = m2.ReplaceAllString(newFile, citeCommands)
	m3 := regexp.MustCompile(`(?s)% begin bib assignment(.*?)% end bib assignment`)
	newFile = m3.ReplaceAllString(newFile, bibAssignments)
	m4 := regexp.MustCompile(`(?s)% begin cite assignment(.*?)% end cite assignment`)
	newFile = m4.ReplaceAllString(newFile, citeAssignments)

	return fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, project, styFilePath), []byte(newFile))
}

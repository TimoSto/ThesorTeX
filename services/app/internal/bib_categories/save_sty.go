package bib_categories

import (
	"regexp"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func SaveCategoriesToSty(store database.ThesorTeXStore, project string, categories []database.BibCategory) error {
	file, err := store.GetBibliographySty(project)
	if err != nil {
		return err
	}

	bibCommands, citeCommands := GeneratePrintCommands(categories)
	bibCommands = "% begin bib commands\n" + bibCommands + "% end bib commands\n"
	citeCommands = "% begin cite commands\n" + citeCommands + "% end cite commands\n"

	bibAssignments, citeAssignments := GenerateAssignment(categories)
	bibAssignments = "% begin bib assignment\n" + bibAssignments + "\t% end bib assignment\n"
	citeAssignments = "% begin cite assignment\n" + citeAssignments + "\t% end cite assignment\n"

	m1 := regexp.MustCompile(`(?s)% begin bib commands(.*?)% end bib commands`)
	newFile := m1.ReplaceAllString(file, bibCommands)
	m2 := regexp.MustCompile(`(?s)% begin cite commands(.*?)% end cite commands`)
	newFile = m2.ReplaceAllString(newFile, citeCommands)
	m3 := regexp.MustCompile(`(?s)% begin bib assignment(.*?)% end bib assignment`)
	newFile = m3.ReplaceAllString(newFile, bibAssignments)
	m4 := regexp.MustCompile(`(?s)% begin cite assignment(.*?)% end cite assignment`)
	newFile = m4.ReplaceAllString(newFile, citeAssignments)

	return store.WriteBibliographySty(project, []byte(newFile))
}

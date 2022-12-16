package bib_categories

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/pkg/romannumerals"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

const (
	commandTemplate = `\newcommand{\print%s}[0]{%%
	\hangindent=\bibparindent%%
	\parindent 0pt%%
	\hangafter=1%%
	%s%%
	\\%%
	\vspace{-12pt}%%

}
`
	citeTemplate = `\newcommand{\cite%s}[0]{%%
	%s%%
}`
)

// GeneratePrintCommands This func gives the print commands for the bibliography and the cites
func GeneratePrintCommands(categories []database.BibCategory) (string, string) {
	bibCommands := ""
	citeCommands := ""

	for _, c := range categories {
		bibCommands += fmt.Sprintf(commandTemplate, c.Name, GenerateCommand(c.Fields, nil))
		citeCommands += fmt.Sprintf(citeTemplate, c.Name, GenerateCommand(c.CiteFields, c.Fields))
	}

	return bibCommands, citeCommands
}

func GenerateCommand(fields []database.Field, searchfields []database.Field) string {
	command := ""

	// count the fields, found in search to determine correct index of field not found in searchfields
	foundCount := 0

	for i, f := range fields {
		command += f.Prefix
		index := i
		if searchfields != nil {
			found := false
			for n, sf := range searchfields {
				if sf.Field == f.Field {
					found = true
					foundCount++
					index = n
				}
			}
			if !found {
				index += len(searchfields) - foundCount
			}
		}
		arg := fmt.Sprintf(`\arg%s`, romannumerals.IntegerToRoman(index+1))
		if f.Italic {
			arg = fmt.Sprintf(`\textit{%s}`, arg)
		}
		command += arg
		command += f.Suffix
	}

	return command
}

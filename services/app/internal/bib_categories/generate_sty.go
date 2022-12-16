package bib_categories

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/pkg/romannumerals"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

const (
	commandTemplate = `\newcommand{\print%s}[%d]{%%
	\hangindent=\bibparindent%%
	\parindent 0pt%%
	\hangafter=1%%
	%s%%
	\\%%
	\vspace{-12pt}%%

}
`
	citeTemplate = `\newcommand{\cite%s}[%d]{%%
	%s%%
}`
)

// GeneratePrintCommands This func gives the print commands for the bibliography and the cites
func GeneratePrintCommands(categories []database.BibCategory) (string, string) {
	bibCommands := ""
	citeCommands := ""

	for _, c := range categories {
		bibCommands += fmt.Sprintf(commandTemplate, c.Name, len(c.Fields), "")
		citeCommands += fmt.Sprintf(citeTemplate, c.Name, len(c.Fields), "")
	}

	return bibCommands, citeCommands
}

func GenerateCommand(fields []database.Field) string {
	command := ""

	for i, f := range fields {
		command += f.Prefix
		arg := fmt.Sprintf(`\arg%s`, romannumerals.IntegerToRoman(i+1))
		if f.Italic {
			arg = fmt.Sprintf(`\textit{%s}`, arg)
		}
		command += arg
		command += f.Suffix
	}

	return command
}

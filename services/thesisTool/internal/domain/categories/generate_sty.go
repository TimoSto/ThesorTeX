package categories

import (
	"fmt"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/romannumerals"
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
}
`
)

// GeneratePrintCommands This func gives the print commands for the bibliography and the cites
func GeneratePrintCommands(categories []Category) (string, string) {
	bibCommands := ""
	citeCommands := ""

	for _, c := range categories {
		bibCommands += fmt.Sprintf(commandTemplate, c.Name, GenerateCommand(c.BibFields, nil))
		citeCommands += fmt.Sprintf(citeTemplate, c.Name, GenerateCommand(c.CiteFields, c.BibFields))
	}

	return bibCommands, citeCommands
}

// GenerateCommand this function generates the command, that prints the fields with style, prefix and suffix
func GenerateCommand(fields []field, searchfields []field) string {
	command := ""

	// count the fields, found in search to determine correct index of field not found in searchfields
	foundCount := 0

	for i, f := range fields {
		command += strings.Replace(f.Format.Prefix, " ", "{ }", -1)
		index := i
		if searchfields != nil {
			found := false
			for n, sf := range searchfields {
				if sf.Name == f.Name {
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
		if f.Format.Italic {
			arg = fmt.Sprintf(`\textit{%s}`, arg)
		}
		command += arg
		command += strings.Replace(f.Format.Suffix, " ", "{ }", -1)
	}

	return command
}

func GenerateAssignment(categories []Category) (string, string) {
	bibAssignments := ""
	citeAssignments := ""

	for _, c := range categories {
		bibAssignments += fmt.Sprintf("\t\t{%s}{%s%s}%%\n", c.Name, `\print`, c.Name)
		citeAssignments += fmt.Sprintf("\t\t{%s}{%s%s}%%\n", c.Name, `\cite`, c.Name)
	}

	return bibAssignments, citeAssignments
}

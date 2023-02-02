package tex_escaping

import (
	"fmt"
	"strings"
)

type Pair struct {
	String string
	TeX    string
}

var pairs = []Pair{
	{
		String: "_",
		TeX:    "{{\\_}}",
	},
	{
		String: ";",
		TeX:    "{{;}}",
	},
	{ //TODO: ggf vorher amp ersetzen
		String: "&",
		TeX:    "{{\\&}}",
	},
	{
		String: "$",
		TeX:    "{{\\$}}",
	},
	{
		String: "#",
		TeX:    "{{\\#}}",
	},
	{
		String: "%",
		TeX:    "{{\\%}}",
	},
	{
		String: "á",
		TeX:    "{{\\'{a}}}",
	},
	{
		String: "â",
		TeX:    "{{\\u{a}}}",
	},
	{
		String: "å",
		TeX:    "{{\\r{a}}}",
	},
	{
		String: "é",
		TeX:    "{{\\'{e}}}",
	},
	{
		String: "è",
		TeX:    "{{\\`{e}}}",
	},
	{
		String: "<",
		TeX:    "{{\\textless}}",
	},
	{
		String: ">",
		TeX:    "{{\\textgreater}}",
	},
	{
		String: "°",
		TeX:    "{{\\degree}}",
	},
	{
		String: "š",
		TeX:    "{{\\v{s}}}",
	},
	{
		String: "č",
		TeX:    "{{\\v{c}}}",
	},
}

func EscapeField(field string) string {
	for _, p := range pairs {
		if p.TeX == fmt.Sprintf("{{\\%s}}", p.String) || p.TeX == "{{;}}" {
			// if the tex-value is just the value surrounded by {{\\}}
			offset := 0
			for {
				i := strings.Index(field[offset:], p.String)
				if i == -1 {
					break
				}
				iOffset := i + offset
				if iOffset > 2 && iOffset < len(field)-2 {
					surrounded := false
					// dont escape if it is already escaped
					if p.TeX == "{{;}}" && field[iOffset-2:iOffset+2] == "{{;}}" || field[iOffset-3:iOffset+3] == fmt.Sprintf("{{\\%s}}", p.String) {
						surrounded = true
					}

					if !surrounded {
						field = field[:iOffset] + p.TeX + field[iOffset+1:]
					}
				}
				offset += i + 6
			}
		} else {
			field = strings.ReplaceAll(field, p.String, p.TeX)
		}
	}

	return field
}

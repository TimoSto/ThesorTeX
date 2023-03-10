package parser

import (
	"fmt"
	"strings"
)

func GetRelevantNotes(file []byte, version string) string {
	var notes string

	strFile := string(file)

	begin := strings.Index(strFile, fmt.Sprintf("# v%s", version))

	end := strings.Index(strFile[begin+3:], "# v")

	if end > 0 {
		notes = strFile[begin : begin+end+1]
	} else {
		notes = strFile[begin:]
	}

	return notes
}

package parser

import (
	"fmt"
	"strings"
)

type Versions struct {
	ThesisTool     string
	ThesisTemplate string
	CVTemplate     string
}

func GetVersions(file []byte, prefix string) string {
	var v string

	lines := strings.Split(string(file), "\n")

	for _, l := range lines {
		if len(l) > len(prefix)+1 && l[:len(prefix)+1] == fmt.Sprintf("%s=", prefix) {
			v = l[len(prefix)+1:]
		}
	}

	return v
}

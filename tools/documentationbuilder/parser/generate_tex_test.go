package parser

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateContentForTeX(t *testing.T) {
	bodies := simpleExpectedBody

	content, err := GenerateContentForTeX("test", bodies)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedContent := `\part{test}
\section{Doc 1}
Some content
\section{Doc 2 styled}
Simple content \textit{with} \textbf{some} \textit{\textbf{styling}}
\section{Extra spacing}
Simple\\
content \textit{with} \textbf{some} \textit{\textbf{styling}}
`
	expSlice := strings.Split(expectedContent, "\n")
	gotSlice := strings.Split(string(content), "\n")
	if diff := cmp.Diff(expSlice, gotSlice); diff != "" {
		t.Errorf("%s", diff)
	}

}

func TestGenerateContentForTeXWithCode(t *testing.T) {
	bodies := withCodeExpectedBody

	content, err := GenerateContentForTeX("test 2", bodies)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedContent := `\part{test 2}
\section{Doc 1}
Some content
\begin{verbatim}
\testcommdand{}
hallo
\end{verbatim}
test
`
	expSlice := strings.Split(expectedContent, "\n")
	gotSlice := strings.Split(string(content), "\n")
	if diff := cmp.Diff(expSlice, gotSlice); diff != "" {
		t.Errorf("%s", diff)
	}

}

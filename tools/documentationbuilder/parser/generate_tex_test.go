package parser

import (
	"fmt"
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

func TestGenerateContentForTeXWithImage(t *testing.T) {
	bodies := withImageExpectedBody

	content, err := GenerateContentForTeX("test 3", bodies)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedContent := `\part{test 3}
\section{Doc 1}
Some content
\begin{figure}[H]
\includegraphics[width=\textwidth]{./images/img1.png}
\caption{test image}
\end{figure}
\noindent test
`
	expSlice := strings.Split(expectedContent, "\n")
	gotSlice := strings.Split(string(content), "\n")
	if diff := cmp.Diff(expSlice, gotSlice); diff != "" {
		t.Errorf("%s", diff)
	}

}

func TestGenerateContentForTeXWithLinks(t *testing.T) {
	bodies := withLinksExpectedBody

	content, err := GenerateContentForTeX("test 3", bodies)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedContent := `\part{test 3}
\section{Doc 1}
Some content \href{https://localhost.test1}{\nolinkurl{link 1}} and \href{https://localhost.test2}{\nolinkurl{https://localhost.test2}}.
`
	expSlice := strings.Split(expectedContent, "\n")
	gotSlice := strings.Split(string(content), "\n")
	if diff := cmp.Diff(expSlice, gotSlice); diff != "" {
		t.Errorf("%s", diff)
	}

}

func TestGenerateContentForTeX_WithFootnote(t *testing.T) {
	bodies := withFootnoteExpectedBody

	content, err := GenerateContentForTeX("test 3", bodies)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	fmt.Println(string(content))

	expectedContent := `\part{test 3}
\section{Doc 1}
Some content
\section{Doc 2 styled}
Simple content \textit{with} \textbf{some} \textit{\textbf{styling}} and footnote.\footnote{footnote \textit{content}}
\section{Extra spacing}
Simple\\
content \textit{with} \textbf{some} \textit{\textbf{styling}}\footnote{\href{https://testsite/test}{\nolinkurl{https://testsite/test}}}
`

	expSlice := strings.Split(expectedContent, "\n")
	gotSlice := strings.Split(string(content), "\n")
	if diff := cmp.Diff(expSlice, gotSlice); diff != "" {
		t.Errorf("%s", diff)
	}
}

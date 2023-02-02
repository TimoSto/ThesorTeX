package tex_escaping

import "testing"

func TestEscapeField(t *testing.T) {
	scenarios := []struct {
		title    string
		input    string
		expected string
	}{
		{
			title:    "nothing necessary",
			input:    "hallo du",
			expected: "hallo du",
		},
		{
			title:    "one escape (surrounded)",
			input:    "hallo_ du",
			expected: "hallo{{\\_}} du",
		},
		{
			title:    "two escapes nest to each other (surrounded)",
			input:    "hallo__ du",
			expected: "hallo{{\\_}}{{\\_}} du",
		},
		{
			title:    "escape ; (needed for csv)",
			input:    "hallo ; du",
			expected: "hallo {{;}} du",
		},
		{
			title:    "escape more complicated",
			input:    "hallo è du",
			expected: "hallo {{\\`{e}}} du",
		},
		{
			title:    "already escaped",
			input:    "hallo {{\\_}} du",
			expected: "hallo {{\\_}} du",
		},
		{
			title:    "already escaped next to unescaped",
			input:    "hallo {{\\_}}# du",
			expected: "hallo {{\\_}}{{\\#}} du",
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			got := EscapeField(s.input)
			if got != s.expected {
				t.Errorf("got: %s, want %s", got, s.expected)
			}
		})
	}
}

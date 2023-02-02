package entries

import (
	"strings"
	"testing"
)

type csvScenario struct {
	title         string
	entries       []Entry
	expectedLines []string
}

func TestGenerateCsvForEntries(t *testing.T) {
	scenarios := []csvScenario{
		{
			title: "single entry",
			entries: []Entry{
				{
					Key:      "test1",
					Category: "aufsatz",
					Fields: []string{
						"test",
						"of",
						"function",
					},
				},
			},
			expectedLines: []string{
				"key;type;a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z;",
				"test1;aufsatz;test;of;function;;;;;;;;;;;;;;;;;;;;;;;;",
				"empty;empty;a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z;",
			},
		},
		{
			title: "two entries",
			entries: []Entry{
				{
					Key:      "test1",
					Category: "aufsatz",
					Fields: []string{
						"test",
						"of",
						"function",
					},
				},
				{
					Key:      "test2",
					Category: "aufsatz",
					Fields: []string{
						"test",
						"of",
						"function",
					},
				},
			},
			expectedLines: []string{
				"key;type;a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z;",
				"test1;aufsatz;test;of;function;;;;;;;;;;;;;;;;;;;;;;;;",
				"test2;aufsatz;test;of;function;;;;;;;;;;;;;;;;;;;;;;;;",
				"empty;empty;a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z;",
			},
		},
		{
			title:   "zero entries",
			entries: []Entry{},
			expectedLines: []string{
				"key;type;a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z;",
				"empty;empty;a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z;",
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			file := GenerateCsvForEntries(s.entries, nil)
			expected := strings.Join(s.expectedLines, "\n") + "\n"
			if file != expected {
				t.Errorf("expected %v but got %v", expected, string(file))
			}
		})
	}
}

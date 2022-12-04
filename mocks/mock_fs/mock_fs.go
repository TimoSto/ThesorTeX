package mock_fs

import (
	"encoding/json"
	"io/fs"
)

func ReadDir(dir string) ([]fs.FileInfo, error) {
	switch dir {
	case "/test2/":
		return []fs.FileInfo{
			&MockFI{
				name:  "testproject1",
				isDir: true,
			},
			&MockFI{
				name:  "testproject2",
				isDir: true,
			},
		}, nil
	default:
		return []fs.FileInfo{}, nil
	}
}

func Mkdir(name string, p fs.FileMode) error {
	return nil
}

//To prevent circle-dependencies, here types are repeated
type project struct {
	Name            string
	Created         string
	LastModified    string
	NumberOfEntries int
}

type bibEntry struct {
	Key        string
	Category   string
	Fields     []string
	Comment    string
	CiteNumber int
}

type bibCategory struct {
	Name                  string
	CitaviType            string
	CitaviNecessaryFields []string //z.B. nur dieser Typ wenn doi existiert
	Fields                []field
	CiteFields            []field
}

type field struct {
	Field            string
	Style            string
	Prefix           string
	Suffix           string
	TexValue         bool
	CitaviAttributes []string
}

func ReadFile(filename string) ([]byte, error) {
	switch filename {
	case "/test2/testproject1/config.json":
		data := project{
			Name:         "",
			Created:      "2022-11-13",
			LastModified: "2022-11-20",
		}
		return json.Marshal(data)
	case "/test2/testproject2/config.json":
		data := project{
			Name:         "",
			Created:      "2022-11-14",
			LastModified: "2022-11-20",
		}
		return json.Marshal(data)
	case "/test/test_threeEntries/bib_entries.json":
		data := []bibEntry{
			{
				Key:        "test1",
				Category:   "testc",
				Fields:     []string{"a", "bb", "ccc"},
				Comment:    "hallo",
				CiteNumber: 0,
			},
			{
				Key:        "test2",
				Category:   "testc",
				Fields:     []string{"a", "bb", "ccc", "ddd"},
				Comment:    "hallo",
				CiteNumber: 0,
			},
			{
				Key:        "test3",
				Category:   "testc2",
				Fields:     []string{"a", "bb", "ccc"},
				Comment:    "",
				CiteNumber: 0,
			},
		}
		return json.Marshal(data)
	case "/test/test_threeCategories/bib_categories.json":
		data := []bibCategory{
			{
				Name:                  "test1",
				CitaviType:            "",
				CitaviNecessaryFields: nil,
				Fields: []field{
					{
						Field:            "t11",
						Style:            "normal",
						Prefix:           "",
						Suffix:           ", ",
						TexValue:         false,
						CitaviAttributes: nil,
					},
					{
						Field:            "t12",
						Style:            "italic",
						Prefix:           " ",
						Suffix:           ".",
						TexValue:         false,
						CitaviAttributes: nil,
					},
				},
				CiteFields: []field{
					{
						Field:            "t11",
						Style:            "normal",
						Prefix:           "",
						Suffix:           ", ",
						TexValue:         false,
						CitaviAttributes: nil,
					},
				},
			},
			{
				Name:                  "test2",
				CitaviType:            "citaviTestDoi",
				CitaviNecessaryFields: []string{"doi"},
				Fields: []field{
					{
						Field:            "t21",
						Style:            "normal",
						Prefix:           "",
						Suffix:           ", ",
						TexValue:         false,
						CitaviAttributes: []string{"author"},
					},
					{
						Field:            "t22",
						Style:            "italic",
						Prefix:           " ",
						Suffix:           ".",
						TexValue:         false,
						CitaviAttributes: []string{"doi"},
					},
				},
				CiteFields: []field{
					{
						Field:            "t23",
						Style:            "normal",
						Prefix:           "",
						Suffix:           ", ",
						TexValue:         false,
						CitaviAttributes: nil,
					},
				},
			},
			{
				Name:                  "test3",
				CitaviType:            "citaviTest",
				CitaviNecessaryFields: nil,
				Fields: []field{
					{
						Field:            "t31",
						Style:            "normal",
						Prefix:           "",
						Suffix:           ", ",
						TexValue:         false,
						CitaviAttributes: []string{"author"},
					},
					{
						Field:            "t32",
						Style:            "italic",
						Prefix:           " ",
						Suffix:           ".",
						TexValue:         false,
						CitaviAttributes: nil,
					},
				},
				CiteFields: []field{
					{
						Field:            "t33",
						Style:            "normal",
						Prefix:           "",
						Suffix:           ", ",
						TexValue:         false,
						CitaviAttributes: nil,
					},
				},
			},
		}
		return json.Marshal(data)
	default:
		return []byte{}, nil
	}
}

var WrittenFiles = []string{}

func WriteFile(path string, data []byte, p fs.FileMode) error {
	WrittenFiles = append(WrittenFiles, path)
	return nil
}

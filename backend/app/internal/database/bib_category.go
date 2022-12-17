package database

type BibCategory struct {
	Name           string
	CitaviCategory string
	CitaviFilters  []string //z.B. nur dieser Typ wenn doi existiert
	Fields         []Field
	CiteFields     []Field
}

type Field struct {
	Field            string
	Italic           bool
	Prefix           string
	Suffix           string
	TexValue         bool
	CitaviAttributes []string
}
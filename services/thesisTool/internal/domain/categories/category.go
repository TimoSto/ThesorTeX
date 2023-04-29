package categories

type Category struct {
	Name           string
	BibFields      []field
	CiteFields     []field
	CitaviCategory string
	CitaviFilter   []string
}

type field struct {
	Name          string
	Format        format
	CitaviMapping []string
}

type format struct {
	Prefix       string
	Suffix       string
	Italic       bool
	Preformatted bool
}

const (
	categoriesFile = "/data/bib_categories.json"
)

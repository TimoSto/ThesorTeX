package database

type BibEntry struct {
	Key        string
	Category   string
	Fields     []string
	Comment    string
	CiteNumber int
}

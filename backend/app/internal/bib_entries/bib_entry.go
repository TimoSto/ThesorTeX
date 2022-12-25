package bib_entries

const (
	entriesJsonFile = "/data/bib_entries.json"
	entriesCsvFile  = "/bib_entries.csv"
)

type BibEntry struct {
	Key        string
	Category   string
	Fields     []string
	Comment    string
	CiteNumber int
}

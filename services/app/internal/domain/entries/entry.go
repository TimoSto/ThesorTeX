package entries

type Entry struct {
	Key      string
	Category string
	Fields   []string
}

const (
	entriesFile = "/data/bib_entries.json"

	csvFile = "bib_entries.csv"
)

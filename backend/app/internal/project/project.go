package project

const (
	metaDataFile = "/data/config.json"
)

type ProjectMetaData struct {
	Created         string
	LastModified    string
	NumberOfEntries int
}

type ProjectMetaDataWithName struct {
	ProjectMetaData
	Name string
}

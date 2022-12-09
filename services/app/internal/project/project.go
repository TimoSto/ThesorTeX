package project

import "github.com/TimoSto/ThesorTeX/services/app/internal/database"

type ProjectData struct {
	Entries    []database.BibEntry
	Categories []database.BibCategory
}

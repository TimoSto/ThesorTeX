package project

import "github.com/TimoSto/ThesorTeX/backend/app/internal/database"

type ProjectData struct {
	Entries    []database.BibEntry
	Categories []database.BibCategory
}

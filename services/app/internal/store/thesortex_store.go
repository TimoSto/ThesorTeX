package store

import "github.com/TimoSto/ThesorTeX/services/app/internal/store/filesystem"

type Store struct {
	fileSystem filesystem.FileSystem
}

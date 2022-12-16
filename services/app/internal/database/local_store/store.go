package local_store

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type Store struct {
	Config conf.Config
}

const (
	projectEntriesFile = "data/bib_entries.json"

	projectCategoriesFile = "data/bib_categories.json"

	entriesCsvFile = "bib_entries.csv"

	bibliographyStyFile = "styPackages/bibliography.sty"
)

//todo: unit tests expand

func (s *Store) GetAllProjects() ([]database.ProjectMetaData, error) {
	projectsPath := s.Config.ProjectsDir
	dirContent, err := os.ReadDir(projectsPath)
	if err != nil {
		err = os.Mkdir(projectsPath, 0777)
		if err != nil {
			return nil, err
		}
		dirContent, err = os.ReadDir(projectsPath)
		if err != nil {
			return nil, err
		}
	}

	var p []database.ProjectMetaData

	for _, f := range dirContent {
		if f.IsDir() {
			newProj := database.ProjectMetaData{
				Name: f.Name(),
			}
			confFile, err := ioutil.ReadFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, f.Name(), "data/config.json"))
			if err != nil {
				log.Error("got error while trying to read project config: %v", err)
				log.Info("Using empty config for project %s", f.Name())
			} else {
				var pConf database.ProjectMetaData
				err = json.Unmarshal(confFile, &pConf)
				if err != nil {
					log.Error("got error while trying to read project config: %v", err)
					log.Info("Using empty config for project %s", f.Name())
				} else {
					newProj.Created = pConf.Created
					newProj.LastModified = pConf.LastModified
				}
			}
			p = append(
				p,
				newProj,
			)
		}
	}

	return p, nil
}

func (s *Store) CreateProject(metaData database.ProjectMetaData, template fs.FS) error {

	//todo:special error if project already exists
	err := os.MkdirAll(pathbuilder.GetProjectPath(s.Config.ProjectsDir, metaData.Name), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir(pathbuilder.GetPathInProject(s.Config.ProjectsDir, metaData.Name, "styPackages"), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir(pathbuilder.GetPathInProject(s.Config.ProjectsDir, metaData.Name, "data"), os.ModePerm)
	if err != nil {
		return err
	}

	err = fs.WalkDir(template, ".", func(path string, d fs.DirEntry, err error) error {
		// cannot happen
		if err != nil {
			panic(err)
		}
		if d.IsDir() {
			return nil
		}
		b, err := fs.ReadFile(template, path)
		if err != nil {
			return err // or panic or ignore
		}
		path = strings.TrimPrefix(path, "template/")
		path = strings.Replace(path, "example.tex", fmt.Sprintf("%s.tex", metaData.Name), 1)

		return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, metaData.Name, path), b, 0644)
	})

	if err != nil {
		return err
	}

	data, err := json.Marshal(metaData)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, metaData.Name, "data/config.json"), data, 0644)
}

func (s *Store) DeleteProject(name string) error {
	path := pathbuilder.GetProjectPath(s.Config.ProjectsDir, name)

	return os.RemoveAll(path)
}

func (s *Store) GetProjectEntries(project string) ([]database.BibEntry, error) {
	file, err := ioutil.ReadFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, projectEntriesFile))
	if err != nil {
		return nil, err
	}

	var data []database.BibEntry
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Store) SaveProjectEntries(project string, data []database.BibEntry) error {
	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, projectEntriesFile), file, 0777)
}

func (s *Store) WriteCSV(project string, file []byte) error {
	return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, entriesCsvFile), file, 0777)
}

func (s *Store) GetProjectCategories(project string) ([]database.BibCategory, error) {
	file, err := ioutil.ReadFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, projectCategoriesFile))
	if err != nil {
		return nil, err
	}

	var data []database.BibCategory
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Store) SaveProjectCategories(project string, data []database.BibCategory) error {
	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, projectCategoriesFile), file, 0777)
}

func (s *Store) GetBibliographySty(project string) (string, error) {
	file, err := ioutil.ReadFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, bibliographyStyFile))
	if err != nil {
		return "", err
	}
	return string(file), nil
}
func (s *Store) WriteBibliographySty(project string, file []byte) error {
	return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, bibliographyStyFile), file, 0777)
}

//TODO: do i really need 0777 everywhere

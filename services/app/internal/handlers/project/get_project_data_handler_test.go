package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestGetProjectDataHandler_POST(t *testing.T) {
	req, err := http.NewRequest("POST", "/getProjectData", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProjectDataHandler(&fs, cfg))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %v, but got %v", http.StatusBadRequest, status)
	}
}

func TestGetProjectDataHandler_GET(t *testing.T) {
	reqD := getProjectDataData{Name: "test"}
	reqData, _ := json.Marshal(reqD)
	req, err := http.NewRequest("GET", "/getProjectData", bytes.NewReader(reqData))
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	entriesFile := []byte("[\n  {\n    \"Key\": \"testEntry\",\n    \"Category\": \"aufsatz\",\n    \"Fields\": [\n      \"Autor xyz\",\n      \"1999\",\n      \"ThesorTeX - Ein tolles Tool\",\n      \"Random Zeitschrift\",\n      \"99/03\",\n      \"20-22\",\n      \"https://wikipedia.org/\"\n    ]\n  }\n]")
	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", "/data/bib_entries.json"), entriesFile)

	categoriesFile := []byte("[\n  {\n    \"Name\": \"aufsatz\",\n    \"CitaviCategory\": \"\",\n    \"CitaviFilters\": null,\n    \"BibFields\": [\n      {\n        \"Name\": \"Autor\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Jahr\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Titel\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Zeitschrift\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Ausgabe\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Seiten\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"url\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      }\n    ],\n    \"CiteFields\": [\n      {\n        \"Name\": \"Autor_kurz\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Jahr\",\n        \"Format\": {\n          \"Style\": \"italic\",\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      }\n    ]\n  }\n]")
	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", "/data/bib_categories.json"), categoriesFile)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProjectDataHandler(&fs, cfg))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, but got %v", http.StatusOK, status)
	}

	respData, _ := ioutil.ReadAll(rr.Body)
	var respObj ProjectData
	json.Unmarshal(respData, &respObj)

	if len(respObj.Entries) != 1 {
		t.Errorf("expected 1 entry, got %v", respObj.Entries)
	}

	if len(respObj.Categories) != 1 {
		t.Errorf("expected 1 entry, got %v", respObj.Categories)
	}
}

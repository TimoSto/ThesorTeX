package categories

import (
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestDeleteCategory(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "/projects",
	}

	//TODO: generalize fake setup
	projects.CreateProject("test", &fs, cfg)

	allBefore, _ := GetAllCategories("test", &fs, cfg)

	err := DeleteCategory("test", "CitaviArticle", &fs, cfg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	allAfter, err := GetAllCategories("test", &fs, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(allAfter) != len(allBefore)-1 {
		t.Errorf("expected 1 category less but got %v and %v", len(allAfter), len(allBefore))
	}

	for _, c := range allAfter {
		if c.Name == "CitaviArticle" {
			t.Error("did not expect to find CitaviArticle, but did")
		}
	}
}

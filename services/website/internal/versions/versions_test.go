package versions

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/website/internal/buckethandler"
)

func TestGetToolVersions(t *testing.T) {
	db := buckethandler.FakeBucketHandler{}

	db.SetItem("thesisTool/v1.1.0", "test-2")
	db.SetItem("thesisTool/v1.0.0", "test-1")
	db.SetItem("thesisTemplate/v1.0.0", "test-1")

	versions, err := GetToolVersions(false, &db)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := []VersionInfo{
		{
			Name: "v1.1.0",
			Date: "01-01-2022",
		},
		{
			Name: "v1.0.0",
			Date: "01-01-2022",
		},
	}

	if !reflect.DeepEqual(versions, expected) {
		t.Errorf("expected %v, got %v", expected, versions)
	}
}

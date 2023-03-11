package versions

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/buckethandler"
)

type Versions struct {
	Tool           []VersionInfo
	ThesisTemplate []VersionInfo
	CvTemplate     []VersionInfo
}

func GetVersions(bucket *buckethandler.BucketHandler) (Versions, error) {
	obj := Versions{
		Tool:           nil,
		ThesisTemplate: nil,
		CvTemplate:     nil,
	}
	toolsV, err := GetToolVersions(*bucket)
	if err != nil {
		return obj, err
	}
	obj.Tool = toolsV

	thesisV, err := GetThesisVersions(*bucket)
	if err != nil {
		return obj, err
	}
	obj.ThesisTemplate = thesisV

	cvV, err := GetCvVersions(*bucket)
	if err != nil {
		return obj, err
	}
	obj.CvTemplate = cvV

	return obj, nil
}

func GetToolVersions(bucket buckethandler.BucketHandler) ([]VersionInfo, error) {
	var versions []VersionInfo

	output, err := bucket.ListElementsInBucket(context.TODO(), "thesortex-artifacts", "thesisTool")

	if err != nil {
		log.Error("could not read bucket items: %v", err)
		return versions, err
	}

	foundVersions := ""

	for _, object := range output {
		pathParts := strings.Split(*object.Key, "/")
		if pathParts[1] != "latest" && strings.Index(foundVersions, pathParts[1]) == -1 {
			foundVersions += fmt.Sprintf(";%s", pathParts[1])
			versions = append(versions, VersionInfo{
				Name: pathParts[1],
				Date: object.LastModified.Format("02-01-2006"),
			})
		}

	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Name > versions[j].Name
	})

	return versions, nil
}

func GetThesisVersions(bucket buckethandler.BucketHandler) ([]VersionInfo, error) {
	var versions []VersionInfo

	output, err := bucket.ListElementsInBucket(context.TODO(), "thesortex-artifacts", "thesisTemplate")

	if err != nil {
		log.Error("could not read bucket items: %v", err)
		return versions, err
	}

	foundVersions := ""

	for _, object := range output {
		pathParts := strings.Split(*object.Key, "/")
		if pathParts[1] != "latest" && strings.Index(foundVersions, pathParts[1]) == -1 {
			foundVersions += fmt.Sprintf(";%s", pathParts[1])
			versions = append(versions, VersionInfo{
				Name: pathParts[1],
				Date: object.LastModified.Format("02-01-2006"),
			})
		}

	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Name > versions[j].Name
	})

	return versions, nil
}

func GetCvVersions(bucket buckethandler.BucketHandler) ([]VersionInfo, error) {
	var versions []VersionInfo

	output, err := bucket.ListElementsInBucket(context.TODO(), "thesortex-artifacts", "cvTemplate")

	if err != nil {
		log.Error("could not read bucket items: %v", err)
		return versions, err
	}

	foundVersions := ""

	for _, object := range output {
		pathParts := strings.Split(*object.Key, "/")
		if pathParts[1] != "latest" && strings.Index(foundVersions, pathParts[1]) == -1 {
			foundVersions += fmt.Sprintf(";%s", pathParts[1])
			versions = append(versions, VersionInfo{
				Name: pathParts[1],
				Date: object.LastModified.Format("02-01-2006"),
			})
		}

	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Name > versions[j].Name
	})

	return versions, nil
}

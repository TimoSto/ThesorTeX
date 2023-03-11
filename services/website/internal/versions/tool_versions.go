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

func GetVersions(dev bool, bucket *buckethandler.BucketHandler) (Versions, error) {
	obj := Versions{
		Tool:           nil,
		ThesisTemplate: nil,
		CvTemplate:     nil,
	}
	toolsV, err := GetToolVersions(dev, *bucket)
	if err != nil {
		return obj, err
	}
	obj.Tool = toolsV

	thesisV, err := GetThesisVersions(dev, *bucket)
	if err != nil {
		return obj, err
	}
	obj.ThesisTemplate = thesisV

	cvV, err := GetCvVersions(dev, *bucket)
	if err != nil {
		return obj, err
	}
	obj.CvTemplate = cvV

	return obj, nil
}

func GetToolVersions(dev bool, bucket buckethandler.BucketHandler) ([]VersionInfo, error) {
	var versions []VersionInfo

	if dev {
		versions = []VersionInfo{
			{
				Name: "v1.0.0",
				Date: "10-01-2023",
			},
			{
				Name: "v0.0.2",
				Date: "02-01-2023",
			},
			{
				Name: "v0.0.1",
				Date: "20-12-2022",
			},
		}
	} else {
		output, err := bucket.ListElementsInBucket(context.TODO(), "thesortex-artifacts", "tool")

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

		versions[0].Name = fmt.Sprintf("%s (latest)", versions[0].Name)
	}

	return versions, nil
}

func GetThesisVersions(dev bool, bucket buckethandler.BucketHandler) ([]VersionInfo, error) {
	var versions []VersionInfo

	if dev {
		versions = []VersionInfo{
			{
				Name: "v1.0.0",
				Date: "10-01-2023",
			},
			{
				Name: "v0.0.2",
				Date: "02-01-2023",
			},
			{
				Name: "v0.0.1",
				Date: "20-12-2022",
			},
		}
	} else {
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

		versions[0].Name = fmt.Sprintf("%s (latest)", versions[0].Name)
	}

	return versions, nil
}

func GetCvVersions(dev bool, bucket buckethandler.BucketHandler) ([]VersionInfo, error) {
	var versions []VersionInfo

	if dev {
		versions = []VersionInfo{
			{
				Name: "v1.0.0",
				Date: "10-01-2023",
			},
			{
				Name: "v0.0.2",
				Date: "02-01-2023",
			},
			{
				Name: "v0.0.1",
				Date: "20-12-2022",
			},
		}
	} else {
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

		versions[0].Name = fmt.Sprintf("%s (latest)", versions[0].Name)
	}

	return versions, nil
}

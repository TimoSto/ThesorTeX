package versions

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Versions struct {
	Tool           []VersionInfo
	ThesisTemplate []VersionInfo
	CvTemplate     []VersionInfo
}

func GetVersions(dev bool, s3Client *s3.Client) (Versions, error) {
	obj := Versions{
		Tool:           nil,
		ThesisTemplate: nil,
		CvTemplate:     nil,
	}
	toolsV, err := GetToolVersions(dev, s3Client)
	if err != nil {
		return obj, err
	}
	obj.Tool = toolsV

	thesisV, err := GetThesisVersions(dev, s3Client)
	if err != nil {
		return obj, err
	}
	obj.ThesisTemplate = thesisV

	cvV, err := GetCvVersions(dev, s3Client)
	if err != nil {
		return obj, err
	}
	obj.CvTemplate = cvV

	return obj, nil
}

func GetToolVersions(dev bool, s3Client *s3.Client) ([]VersionInfo, error) {
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
		output, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String("thesortex-artifacts"),
			Prefix: aws.String("tool"),
		})

		if err != nil {
			log.Error("could not read bucket items: %v", err)
			return versions, err
		}

		foundVersions := ""

		for _, object := range output.Contents {
			pathParts := strings.Split(*object.Key, "/")
			if pathParts[0] != "latest" && strings.Index(foundVersions, pathParts[0]) == -1 {
				foundVersions += fmt.Sprintf(";%s", pathParts[0])
				versions = append(versions, VersionInfo{
					Name: pathParts[0],
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

func GetThesisVersions(dev bool, s3Client *s3.Client) ([]VersionInfo, error) {
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
		output, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String("thesortex-artifacts"),
			Prefix: aws.String("thesisTemplate"),
		})

		if err != nil {
			log.Error("could not read bucket items: %v", err)
			return versions, err
		}

		foundVersions := ""

		for _, object := range output.Contents {
			pathParts := strings.Split(*object.Key, "/")
			if pathParts[0] != "latest" && strings.Index(foundVersions, pathParts[0]) == -1 {
				foundVersions += fmt.Sprintf(";%s", pathParts[0])
				versions = append(versions, VersionInfo{
					Name: pathParts[0],
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

func GetCvVersions(dev bool, s3Client *s3.Client) ([]VersionInfo, error) {
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
		output, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String("thesortex-artifacts"),
			Prefix: aws.String("cvTemplate"),
		})

		if err != nil {
			log.Error("could not read bucket items: %v", err)
			return versions, err
		}

		foundVersions := ""

		for _, object := range output.Contents {
			pathParts := strings.Split(*object.Key, "/")
			if pathParts[0] != "latest" && strings.Index(foundVersions, pathParts[0]) == -1 {
				foundVersions += fmt.Sprintf(";%s", pathParts[0])
				versions = append(versions, VersionInfo{
					Name: pathParts[0],
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

package versions

import (
	"context"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetToolVersions(dev bool, s3Client *s3.Client) ([]VersionInfo, error) {
	var versions []VersionInfo

	if dev {
		versions = []VersionInfo{
			{
				Name: "1.0.0",
				Date: "10-01-2023",
			},
			{
				Name: "0.0.2",
				Date: "02-01-2023",
			},
			{
				Name: "0.0.1",
				Date: "20-12-2022",
			},
		}
	} else {
		output, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String("thesortex-artifacts"),
		})

		if err != nil {
			log.Error("could not read bucket items: %v", err)
			return versions, err
		}

		for _, object := range output.Contents {
			versions = append(versions, VersionInfo{
				Name: *object.Key,
				Date: object.LastModified.String(),
			})
		}
	}

	return versions, nil
}

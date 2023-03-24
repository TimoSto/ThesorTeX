package s3

import (
	"context"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateS3Client() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Error("could not create s3 client: %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return client, nil
}

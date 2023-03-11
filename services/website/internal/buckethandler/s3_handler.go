package buckethandler

import (
	"context"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3BucketHandler struct {
	client *s3.Client
}

func New(client *s3.Client) S3BucketHandler {
	return S3BucketHandler{
		client: client,
	}
}

func (s *S3BucketHandler) ListElementsInBucket(ctx context.Context, bucket string, subDir string) ([]BucketItem, error) {
	output, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(subDir),
	})

	if err != nil {
		log.Error("could not read bucket items: %v", err)
		return nil, err
	}

	var result []BucketItem

	for _, object := range output.Contents {
		result = append(result, BucketItem{
			Key:          object.Key,
			LastModified: object.LastModified,
		})
	}

	return result, nil
}

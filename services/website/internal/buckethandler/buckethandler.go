package buckethandler

import (
	"context"
	"time"
)

type BucketHandler interface {
	ListElementsInBucket(ctx context.Context, bucket string, subDir string) ([]BucketItem, error)
}

type BucketItem struct {
	Key          *string
	LastModified *time.Time
}

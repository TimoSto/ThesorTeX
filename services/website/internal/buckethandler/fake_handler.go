package buckethandler

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type FakeBucketHandler struct {
	items map[string]string
}

func (s *FakeBucketHandler) SetItem(key, value string) {
	if s.items == nil {
		s.items = make(map[string]string)
	}

	s.items[key] = value
}

func (s *FakeBucketHandler) ListElementsInBucket(ctx context.Context, bucket string, subDir string) ([]BucketItem, error) {
	var result []BucketItem

	t := time.Date(2022, time.January, 1, 1, 1, 1, 1, time.UTC)

	for k, _ := range s.items {
		fmt.Println(k, strings.Index(k, subDir) >= 0)
		if strings.Index(k, subDir) >= 0 {
			key := k
			result = append(result, BucketItem{
				Key:          &key,
				LastModified: &t,
			})
		}
	}

	fmt.Println(result)

	return result, nil
}

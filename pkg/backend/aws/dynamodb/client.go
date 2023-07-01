package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamoClient(opts func(*config.LoadOptions) error) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return client, nil
}

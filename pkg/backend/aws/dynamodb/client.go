package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamoClient(opts func(*config.LoadOptions) error) (*dynamodb.Client, error) {
	var cfg aws.Config
	var err error
	if opts != nil {
		cfg, err = config.LoadDefaultConfig(context.Background(), opts)
	} else {
		cfg, err = config.LoadDefaultConfig(context.Background())
	}

	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return client, nil
}

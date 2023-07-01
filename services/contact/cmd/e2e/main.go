package main

import (
	"context"
	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/dynamodb"
	dynamocontainer "github.com/TimoSto/ThesorTeX/pkg/backend/container/dynamodb"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	dynamobasic "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"time"
)

func main() {
	log.Info("Starting the fake contact application...")

	c, err := dynamocontainer.StartDynamoDBLocally()
	if err != nil {
		log.Fatal("could not start local dynamo db: %v", err)
	}

	err = createFakeTable(c.Endpoint())
	if err != nil {
		log.Fatal("could not create local dynamo db table: %v", err)
	}

	time.Sleep(5 * time.Second)

	defer c.Purge()
}

func createFakeTable(endpoint string) error {
	opts := config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{URL: endpoint}, nil
		}),
	)

	client, err := dynamodb.GetDynamoClient(opts)
	if err != nil {
		return err
	}

	_, err = client.CreateTable(context.Background(), &dynamobasic.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName: aws.String("feedbacks"),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

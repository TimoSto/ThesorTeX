package feedback_test

import (
	"context"
	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/dynamodb"
	dynamocontainer "github.com/TimoSto/ThesorTeX/pkg/backend/container/dynamodb"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/feedback"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	dynamobasic "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"testing"
)

func TestStore_SaveFeedback(t *testing.T) {
	c, err := dynamocontainer.RunDynamoDBLocally()
	if err != nil {
		t.Fatalf("could not start local dynamo db: %v", err)
	}

	defer c.Purge()

	opts := config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{URL: c.Endpoint()}, nil
		}),
	)

	client, err := dynamodb.GetDynamoClient(opts)
	if err != nil {
		t.Fatalf("could not create dynamo client: %v", err)
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

	s := feedback.New(client)

	err = s.SaveFeedback("foo", "bar")
	if err != nil {
		t.Errorf("unexpected error saving feedback: %v", err)
	}

	var feedbacks []feedback.Feedback

	data, err := client.Scan(context.Background(), &dynamobasic.ScanInput{
		TableName: aws.String("feedbacks"),
	})
	if err != nil {
		log.Error("could not read db: %v", err)
		return
	}

	err = attributevalue.UnmarshalListOfMaps(data.Items, &feedbacks)
	if err != nil {
		log.Error("UnmarshalListOfMaps: %v\n", err)
		return
	}

	if len(feedbacks) != 1 {
		t.Errorf("expected feedbacks to have a length of 1 but got %v", feedbacks)
	}

	if feedbacks[0].Subject != "foo" || feedbacks[0].Message != "bar" {
		t.Errorf("expected foo bar but got %s %s", feedbacks[0].Subject, feedbacks[0].Message)
	}
}

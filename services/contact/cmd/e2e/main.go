package main

import (
	"context"
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/dynamodb"
	dynamocontainer "github.com/TimoSto/ThesorTeX/pkg/backend/container/dynamodb"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/backend"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/feedback"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	dynamobasic "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"net/http"
	"os"
)

func main() {
	log.Info("Starting the fake contact application...")

	c, err := dynamocontainer.StartDynamoDBLocally()
	if err != nil {
		log.Fatal("could not start local dynamo db: %v", err)
	}

	defer c.Purge()

	testClient, err := createFakeTable(c.Endpoint())
	if err != nil {
		log.Fatal("could not create local dynamo db table: %v", err)
	}

	tables, err := testClient.ListTables(
		context.TODO(), &dynamobasic.ListTablesInput{})

	fmt.Println("got tables: ", tables.TableNames)

	sigs := make(chan os.Signal, 1)

	mux := http.NewServeMux()

	setupFakeGetter(mux, testClient)

	opts := config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{URL: c.Endpoint()}, nil
		}),
	)

	cfg := backend.Config{
		Mux:        mux,
		DynamoOpts: opts,
	}

	err = backend.StartApp(cfg)
	if err != nil {
		log.Fatal("could not start backend: %v", err)
	}

	<-sigs
}

func createFakeTable(endpoint string) (*dynamobasic.Client, error) {
	opts := config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{URL: endpoint}, nil
		}),
	)

	client, err := dynamodb.GetDynamoClient(opts)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return client, nil
}

func setupFakeGetter(mux *http.ServeMux, dbClient *dynamobasic.Client) {
	mux.HandleFunc("/getE2EFeedbacks", func(writer http.ResponseWriter, request *http.Request) {
		var feedbacks []feedback.Feedback

		data, err := dbClient.Scan(context.Background(), &dynamobasic.ScanInput{
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

		log.Info("got %v feedback messages", len(feedbacks))
		for _, f := range feedbacks {
			log.Info("[%s] - %s : '%s'", f.ID, f.Date, f.Message)
		}
	})
}

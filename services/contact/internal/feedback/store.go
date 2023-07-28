package feedback

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	dynamobasic "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"time"
)

type Store struct {
	DynamoClient *dynamobasic.Client
}

func New(db *dynamobasic.Client) Store {
	return Store{
		DynamoClient: db,
	}
}

func (s *Store) SaveFeedback(message string) error {
	feedbackObj := Feedback{
		ID:      uuid.New().String(),
		Date:    time.Now().Format(time.RFC822),
		Message: message,
	}

	data, err := attributevalue.MarshalMap(feedbackObj)
	if err != nil {
		return err
	}

	_, err = s.DynamoClient.PutItem(context.TODO(), &dynamobasic.PutItemInput{
		TableName: aws.String("feedbacks"),
		Item:      data,
	})

	if err != nil {
		return err
	}

	return nil
}
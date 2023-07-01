package main

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/container/dynamodb"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"time"
)

func main() {
	log.Info("Starting the fake contact application...")

	c, err := dynamodb.StartDynamoDBLocally()
	if err != nil {
		log.Fatal("could not start local dynamo db: %v", err)
	}

	time.Sleep(5 * time.Second)

	defer c.Purge()
}

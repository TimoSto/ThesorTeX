package dynamodb

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/container"
	"github.com/ory/dockertest/v3"
	"os"
)

func RunDynamoDBLocally() (*container.Container, error) {
	err := os.Setenv("AWS_ACCESS_KEY_ID", "foo")
	if err != nil {
		return nil, err
	}
	err = os.Setenv("AWS_SECRET_ACCESS_KEY", "bar")
	if err != nil {
		return nil, err
	}
	err = os.Setenv("AWS_SESSION_TOKEN", "foobar")
	if err != nil {
		return nil, err
	}

	c, err := container.RunImage(&dockertest.RunOptions{
		Repository:   "docker.io/amazon/dynamodb-local",
		Tag:          "latest",
		ExposedPorts: []string{"8000/tcp"},
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

package dynamodb

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/container"
	"github.com/ory/dockertest/v3"
	"os"
)

func SetTestAWSEnvironment() {
	env := map[string]string{
		"AWS_ACCESS_KEY_ID":     "foo",
		"AWS_SECRET_ACCESS_KEY": "bar",
		"AWS_SESSION_TOKEN":     "baz",
	}

	for key, val := range env {
		os.Setenv(key, val)
	}
}

func StartDynamoDBLocally() (*container.Container, error) {
	SetTestAWSEnvironment()

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

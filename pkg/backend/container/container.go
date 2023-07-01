package container

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"os"
	"time"
)

type Container struct {
	endpoint string
	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func (c *Container) Purge() error {
	return c.pool.Purge(c.resource)
}

func connectToPool() (*dockertest.Pool, error) {
	dockerHost := os.Getenv("DOCKER_HOST")

	pool, err := dockertest.NewPool(dockerHost)
	if err != nil {
		return nil, fmt.Errorf("could not connect to docker: %s", err)
	}

	return pool, nil
}

func RunImage(options *dockertest.RunOptions) (*Container, error) {
	pool, err := connectToPool()
	if err != nil {
		return nil, err
	}

	res, err := pool.RunWithOptions(options)
	if err != nil {
		return nil, err
	}

	if err := pool.Retry(func() error {
		time.Sleep(2 * time.Second)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("could not connect to container: %s", err)
	}

	port := res.GetPort(options.ExposedPorts[0])

	return &Container{
		endpoint: "http://localhost:" + port,
		pool:     pool,
		resource: res,
	}, nil
}

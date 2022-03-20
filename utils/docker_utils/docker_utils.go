package docker_utils

import (
	"context"

	"github.com/docker/docker/client"
)

func Ping() (bool, error) {
	var err error

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return false, err
	}

	_, err = cli.ServerVersion(ctx)
	if err == nil {
		return true, nil
	}
	return false, err
}

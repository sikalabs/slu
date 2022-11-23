package docker_utils

import (
	"context"

	"github.com/docker/docker/api/types"
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

func ListContainerIDs() ([]string, error) {
	var err error

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}

	var containerIDs []string
	for _, container := range containers {
		containerIDs = append(containerIDs, container.ID)
	}

	return containerIDs, nil
}

func RemoveContainerByID(id string, force bool) error {
	var err error

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	err = cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{
		Force: force,
	})
	if err != nil {
		return err
	}

	return nil
}

func ForceRemoveAllContainers() error {
	var err error

	containers, err := ListContainerIDs()
	if err != nil {
		return err
	}

	for _, container := range containers {
		err = RemoveContainerByID(container, true)
		if err != nil {
			return err
		}
	}

	return nil
}

package docker

import (
	"context"
	"fmt"

	"github.com/Doer-org/ketos/internal/docker"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CreateContainer() (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	// コンテナを作成
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: docker.ImageName,
	}, nil, nil, nil, docker.ContainerName)
	if err != nil {
		return "", err
	}
	fmt.Printf("Container has been created\n")
	return resp.ID, nil
}

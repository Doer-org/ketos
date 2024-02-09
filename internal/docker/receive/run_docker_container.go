package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunConrainer(respID string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// コンテナを起動
	if err := cli.ContainerStart(ctx, respID, container.StartOptions{}); err != nil {
		return err
	}
	fmt.Printf("Container has been started\n")
	return nil
}

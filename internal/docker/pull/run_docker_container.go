package docker

import (
	"context"
	"fmt"

	"github.com/Doer-org/ketos/internal/docker"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunContainer(respID string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// Start the container
	if err := cli.ContainerStart(ctx, respID, container.StartOptions{}); err != nil {
		return err
	}
	fmt.Printf("Container %s has been started\n",docker.ContainerName)
	return nil
}

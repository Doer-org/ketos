package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/Doer-org/ketos/internal/docker"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func CreateContainer(port string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	// ポートのマッピングを解析する
	portMappings := strings.Split(port, ":")
	if len(portMappings) != 2 {
		return "", fmt.Errorf("invalid port mapping format")
	}
	hostPort := portMappings[0]
	containerPort := portMappings[1] + "/tcp"

	// PortBindingsの設定
	portBindings := nat.PortMap{
		nat.Port(containerPort): []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort,
			},
		},
	}

	// コンテナを作成
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        docker.ImageName,
		ExposedPorts: nat.PortSet{nat.Port(containerPort): struct{}{}},
	}, &container.HostConfig{
		PortBindings: portBindings,
	}, nil, nil, docker.ContainerName)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	fmt.Printf("Container has been created with ID: %s\n", resp.ID)
	return resp.ID, nil
}

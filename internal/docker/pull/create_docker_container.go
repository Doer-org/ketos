package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/Doer-org/ketos/internal/docker"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func StopAndRemoveContainer(containerName string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// search for the container by name
	filter := filters.NewArgs()
	filter.Add("name", containerName)
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true, Filters: filter})
	if err != nil {
		return err
	}
	for _, containerFiltered := range containers {
		if containerFiltered.Names[0] == "/"+containerName {
			// stop the container
			if err := cli.ContainerStop(ctx, containerFiltered.ID, container.StopOptions{}); err != nil {
				return err
			}
			// remove the container
			if err := cli.ContainerRemove(ctx, containerFiltered.ID, container.RemoveOptions{}); err != nil {
				return err
			}
			fmt.Printf("Container %s has been stopped and removed.\n", containerName)
			break
		}
	}	
	return nil
}


func CreateContainer(port string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	// Stop and remove the container if it already exists
	err = StopAndRemoveContainer(docker.ContainerName)
	if err != nil {
		return "", err
	}

	// Parse the port mappings
	portMappings := strings.Split(port, ":")
	if len(portMappings) != 2 {
		return "", fmt.Errorf("invalid port mapping format")
	}
	hostPort := portMappings[0]
	containerPort := portMappings[1] + "/tcp"

	// Set up PortBindings
	portBindings := nat.PortMap{
		nat.Port(containerPort): []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort,
			},
		},
	}

	// Create the container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        docker.ImageName,
		ExposedPorts: nat.PortSet{nat.Port(containerPort): struct{}{}},
	}, &container.HostConfig{
		PortBindings: portBindings,
	}, nil, nil, docker.ContainerName)
	if err != nil {
		return "", err
	}
	fmt.Printf("Container %s has been created with ID: %s\n", docker.ContainerName,resp.ID)
	return resp.ID, nil
}

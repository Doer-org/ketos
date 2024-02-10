package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/Doer-org/ketos/internal/docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func CreateImage(dockerFile bool, languageType string, directory string, dockerfilename string) error {

	if dockerFile {
		err := createImageWithDockerFile(directory, dockerfilename)
		if err != nil {
			return err
		}
	} else {
		err := createImageWithBuildPacks(directory, dockerfilename, languageType)
		if err != nil {
			return err
		}
	}

	return nil
}

func createImageWithDockerFile(directory string, dockerfilename string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	tar, err := archive.TarWithOptions(directory, &archive.TarOptions{})
	if err != nil {
		return err
	}
	fmt.Println("Creating image with Dockerfile")
	fmt.Println("Path: ", directory)
	fmt.Println("Dockerfile: ", dockerfilename)
	defer tar.Close()
	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		tar,
		types.ImageBuildOptions{
			Tags:       []string{docker.ImageName},
			Dockerfile: dockerfilename,
			Remove:     true,
			NoCache:    true,
		},
	)
	if err != nil {
		return err
	}

	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}
	return nil
}

func createImageWithBuildPacks(directory string, dockerfilename string, language string) error {
	builder := responseBuilder(language)
	cmd := exec.Command("pack", "build", docker.ImageName, "--builder", builder, "--path", directory)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func responseBuilder(language string) string {
	switch language {
	case "node":
		return "gcr.io/buildpacks/builder:v1"
	case "go":
		return "gcr.io/buildpacks/builder:v1"
	case "python":
		return "gcr.io/buildpacks/builder:v1"
	case "java":
		return "gcr.io/buildpacks/builder:v1"
	case "ruby":
		return "paketobuildpacks/builder:base"
	default:
		return "gcr.io/buildpacks/builder:v1"
	}
}

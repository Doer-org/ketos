package docker

import (
	"compress/gzip"
	"context"
	"io"
	"os"

	"github.com/Doer-org/ketos/internal/docker"
	"github.com/docker/docker/client"
)

func CompressImageToTarGz() error {
	// imageをtar.gzに圧縮
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	if _, err := os.Stat(docker.TarTmpDir); os.IsNotExist(err) {
		os.Mkdir(docker.TarTmpDir, 0777)
	}
	tarGzFileName := docker.TarTmpDir + "/" + docker.ImageName + ".tar.gz"
	imageSaveResponse, err := cli.ImageSave(ctx, []string{docker.ImageName})
	if err != nil {
		return err
	}
	defer imageSaveResponse.Close()
	file, err := os.Create(tarGzFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	gw := gzip.NewWriter(file)
	defer gw.Close()

	if _, err = io.Copy(gw, imageSaveResponse); err != nil {
		return err
	}

	return nil
}

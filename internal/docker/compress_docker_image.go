package docker

import (
	"context"
	"os"

	"github.com/docker/docker/client"
)

func CompressImageToTar() {
	// imageをtarに圧縮
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(TarTmpDir); os.IsNotExist(err) {
		os.Mkdir(TarTmpDir, 0777)
	}
	tarFileName := TarTmpDir + "/" + ImageName + ".tar"
	imageSaveResponse, err := cli.ImageSave(ctx, []string{ImageName})
	if err != nil {
		panic(err)
	}
	defer imageSaveResponse.Close()
	file, err := os.Create(tarFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.ReadFrom(imageSaveResponse)
	if err != nil {
		panic(err)
	}
}

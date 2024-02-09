package docker

import (
	"context"
	"os"

	"github.com/Doer-org/ketos/internal/docker"
	"github.com/docker/docker/client"
)

func DecompressTarToImage() error {
	// tarをimageに展開
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	tarFileName := docker.TarTmpDir + "/" + docker.ImageName + ".tar"
	imageTar, err := os.Open(tarFileName)
	if err != nil {
		return err
	}
	defer imageTar.Close()
	loadResponse, err := cli.ImageLoad(ctx, imageTar, true)
	if err != nil {
		return err
	}
	defer loadResponse.Body.Close()
	// tarファイルを削除
	// err = os.Remove(tarFileName)
	// if err != nil {
	// 	panic(err)
	// }
	return nil
}

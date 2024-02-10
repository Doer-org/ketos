package docker

import (
	"compress/gzip"
	"context"
	"os"

	"github.com/Doer-org/ketos/internal/docker"
	"github.com/docker/docker/client"
)

func DecompressTarGzToImage() error {
	// tar.gzをimageに展開
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	tarGzFileName := "../../../tmp-tar" + "/" + docker.ImageName + ".tar.gz"

	gzFile, err := os.Open(tarGzFileName)
	if err != nil {
		return err
	}
	defer gzFile.Close()

	gzReader, err := gzip.NewReader(gzFile)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	loadResponse, err := cli.ImageLoad(ctx, gzReader, true)
	if err != nil {
		return err
	}
	defer loadResponse.Body.Close()

	// tar.gzファイルを削除するコードはコメントアウトされていますが、必要に応じて有効化してください
	// err = os.Remove(tarGzFileName)
	// if err != nil {
	// 	panic(err)
	// }

	return nil
}

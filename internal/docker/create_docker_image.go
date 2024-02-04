package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func CreateImageWithPack() {
	// Packを外部コマンドで実行してimageを作成
	cmd := exec.Command("pack", "build", ImageName, "--builder", "gcr.io/buildpacks/builder:v1")
	if dirPath != "" {
		cmd.Dir = dirPath
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func createImageWithDockerFile(path string, dockerfilename string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	tar, err := archive.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}
	defer tar.Close()
	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		tar,
		types.ImageBuildOptions{
			Tags:       []string{"ketos-tmp-image:latest"},
			Dockerfile: dockerfilename, // またはDockerfileの相対パス
			Remove:     true,           // ビルド後に中間コンテナを削除
			NoCache:    true,           // キャッシュを使わない
		},
	)
	if err != nil {
		return err
	}
	defer imageBuildResponse.Body.Close()
	// ビルドログを標準出力に表示
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}
	fmt.Println("Image build successful!")
	return nil
}

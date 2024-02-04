package docker

import (
	"os"
	"os/exec"
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

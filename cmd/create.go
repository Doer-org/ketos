/*
Copyright © 2024 Do'er
*/
package cmd

import (
	"context"
	"os"
	"os/exec"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

const (
	TarTmpDir   = "tmp-tar"
	ImageName   = "ketos-tmp-image"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Docker image based on your local environment",
	Long: `This command creates a docker image based on the local environment, 
	compresses it, and sends it to the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		createImageWithPack()
		compressImageToTar()
	},
}

func createImageWithPack() {
	// Packを外部コマンドで実行してimageを作成
	cmd := exec.Command("pack", "build", "ketos-tmp-image", "--builder", "gcr.io/buildpacks/builder:v1")
	cmd.Dir = "examples/go"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func compressImageToTar() {
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

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

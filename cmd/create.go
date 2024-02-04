/*
Copyright © 2024 Do'er
*/
package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Docker image based on your local environment",
	Long: `This command creates a docker image based on the local environment, 
	compresses it, and sends it to the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		createImageWithPack()
		compressImage()
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

func compressImage() {
	// imageをtarに圧縮
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

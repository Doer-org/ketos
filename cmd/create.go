/*
Copyright © 2024 Do'er
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Doer-org/ketos/internal/api"
	"github.com/Doer-org/ketos/internal/docker"
)

var dirPath string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Docker image based on your local environment",
	Long: `This command creates a docker image based on the local environment, 
	compresses it, and sends it to the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		docker.CreateImageWithPack()
		docker.CompressImageToTar()
		api.SendTarToServer()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&dirPath, "path", "p", "", "directory path to create docker image")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

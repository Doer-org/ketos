/*
Copyright © 2024 Do'er
*/
package cmd

import (
	"github.com/Doer-org/ketos/internal/api"
	"github.com/Doer-org/ketos/internal/docker"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull Docker image from the server and run it",
	Long: `This command pulls a docker image from the server and runs it.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.ReceiveTarFromServer()
		docker.DecompressTarToImage()
		// docker.RunImage()
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

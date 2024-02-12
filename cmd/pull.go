/*
Copyright © 2024 Do'er
*/
package cmd

import (
	"github.com/Doer-org/ketos/internal"
	"github.com/Doer-org/ketos/internal/api"
	"github.com/Doer-org/ketos/internal/docker"
	dockerpull "github.com/Doer-org/ketos/internal/docker/pull"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull Docker image from the server and run it",
	Long:  `This command pulls a docker image from the server and runs it.`,
	// Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		port, err := api.ReceiveTarGzFromServer(id)
		if err != nil {
			return err
		}
		tarGzFileName := "./tmp-tar" + "/" + docker.ImageName + ".tar.gz"
		err = dockerpull.DecompressTarGzToImage(tarGzFileName)
		if err != nil {
			return err
		}
		respID, err := dockerpull.CreateContainer(port)
		if err != nil {
			return err
		}
		err = dockerpull.RunContainer(respID)
		if err != nil {
			return err
		}
		internal.PrintKetos()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().StringP("id", "i", "", "ketos docker image id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

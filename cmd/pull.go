/*
Copyright © 2024 Do'er
*/
package cmd

import (
	"fmt"

	"github.com/Doer-org/ketos/internal/api"
	docker "github.com/Doer-org/ketos/internal/docker/receive"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull Docker image from the server and run it",
	Long:  `This command pulls a docker image from the server and runs it.`,
	Args:  cobra.ExactArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
	     __ __ ________________  _____
        / //_// ____/_  __/ __ \/ ___/
       / ,<  / __/   / / / / / /\__ \ 
      / /| |/ /___  / / / /_/ /___/ / 
     /_/ |_/_____/ /_/  \____//____/  
                                       						  				   
	`)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := api.ReceiveTarFromServer()
		if err != nil {
			return err
		}
		err = docker.DecompressTarToImage()
		if err != nil {
			return err
		}
		respID, err := docker.CreateContainer()
		if err != nil {
			return err
		}
		err = docker.RunConrainer(respID)
		if err != nil {
			return err
		}
		return nil
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

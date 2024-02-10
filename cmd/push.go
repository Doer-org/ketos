/*
Copyright © 2024 Do'er
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Doer-org/ketos/internal/api"
	"github.com/Doer-org/ketos/internal/docker"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "push",
	Short: "Create Docker image based on your local environment",
	Long: `This command creates a docker image based on the local environment, 
	compresses it, and sends it to the server.`,
	Args: cobra.ExactArgs(0),
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
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}
		language, err := cmd.Flags().GetString("language")
		if err != nil {
			return err
		}
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			return err
		}
		dockerfile, err := cmd.Flags().GetBool("dockerfile")
		if err != nil {
			return err
		}
		fmt.Println("path: ", path)

		err = docker.CreateImage(dockerfile, language, path, filename)
		if err != nil {
			return err
		}
		docker.CompressImageToTar()
		err = api.SendTarToServer()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("path", "p", "", "directory path to create docker image")
	createCmd.Flags().StringP("language", "l", "", "language type to create docker image")
	createCmd.Flags().StringP("filename", "f", "", "dockerfile name to create docker image")
	createCmd.Flags().BoolP("dockerfile", "d", false, "dockerfile or buildpacks")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

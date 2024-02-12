/*
Copyright © 2024 Do'er
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Doer-org/ketos/internal/api"
	docker "github.com/Doer-org/ketos/internal/docker/push"
)

// pushCmd represents the create command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Create Docker image based on your local environment",
	Long: `This command creates a docker image based on the local environment, 
	compresses it, and sends it to the server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: 余分な引数がある場合はエラーを返したい
		// if len(args) > 0 {
		// 	return fmt.Errorf("unexpected argument(s): %v\nUsage: %s", args, cmd.UseLine())
		// }
		directory, err := cmd.Flags().GetString("directory")
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
		publishList, err := cmd.Flags().GetStringSlice("publish")
		if err != nil {
			return err
		}
		envList, err := cmd.Flags().GetStringSlice("env")
		if err != nil {
			return err
		}
		fmt.Println("directory: ", directory)

		err = docker.CreateImage(dockerfile, language, directory, filename)
		if err != nil {
			return err
		}
		err = docker.CompressImageToTarGz()
		if err != nil {
			return err
		}
		err = api.SendTarToServer(publishList, envList)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.Flags().StringP("directory", "d", "", "Directory path to create docker image")
	pushCmd.Flags().StringP("language", "l", "", "Language type to create docker image")
	pushCmd.Flags().StringP("filename", "f", "", "Dockerfile name to create docker image")
	pushCmd.Flags().BoolP("dockerfile", "D", false, "Dockerfile or buildpacks")
	pushCmd.Flags().StringSliceP("publish", "p", []string{}, "Publish a container's port(s) to the host")
	pushCmd.Flags().StringSliceP("env", "e", []string{}, "Set environment variable(s)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

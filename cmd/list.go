/*
Copyright © 2024 jj_ <jj_@team-jj.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list workspaces under the base directory",
	Long:  `list workspaces under the base directory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usecase, err := NewUsecase(baseFile)
		if err != nil {
			return err
		}

		wss, err := usecase.ListWorkspace()
		if err != nil {
			return err
		}

		if len(wss) == 0 {
			fmt.Fprintln(os.Stdout, "No workspace found")
		}

		ShowWorkspaces(wss...)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

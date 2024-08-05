/*
Copyright © 2024 jj_ <jj_@team-jj.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [WORKSPACE FILE PATH]",
	Short: "add workspace to the list",
	Long:  `add workspace to the list`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := func() error {
			usecase, err := NewUsecase(baseFile)
			if err != nil {
				return err
			}

			path := args[0]

			var name string
			name, err = cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			if err := usecase.AddWorkspace(name, path); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("name", "n", "", "name of the workspace. If empty, the filename without the extension is used.")
}

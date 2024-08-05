/*
Copyright © 2024 jj_ <jj_@team-jj.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open [WORKSPACE NAME]",
	Short: "open workspace",
	Long:  `open workspace`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := func() error {
			usecase, err := NewUsecase(baseFile)
			if err != nil {
				return err
			}

			name := args[0]

			if err := usecase.OpenWorkspace(name); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

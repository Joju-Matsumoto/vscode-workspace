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
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 最初の引数以外は補完しない
		if len(args) >= 1 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		wss, err := usecase.ListWorkspace()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, cobra.ShellCompDirectiveError
		}

		vargs := ListWorkspaceNamesWithDescription(wss)

		return vargs, cobra.ShellCompDirectiveDefault
	},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := func() error {
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

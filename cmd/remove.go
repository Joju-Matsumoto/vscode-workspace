/*
Copyright © 2024 jj_ <jj_@team-jj.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Joju-Matsumoto/vscode-workspace/vscodeworkspace"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove [NAME]...",
	Aliases: []string{"rm"},
	Short:   "remove workspace from the list",
	Long:    `remove workspace from the list`,
	Args:    cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		wss, err := usecase.ListWorkspace(vscodeworkspace.ListWorkspaceUsecaseOption{})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, cobra.ShellCompDirectiveError
		}

		vargs := ListWorkspaceNamesWithDescription(wss)

		return vargs, cobra.ShellCompDirectiveDefault
	},
	Run: func(cmd *cobra.Command, args []string) {
		names := args

		for _, name := range names {
			if err := usecase.DeleteWorkspace(name); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

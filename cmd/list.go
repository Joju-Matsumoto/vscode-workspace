/*
Copyright © 2024 jj_ <jj_@team-jj.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Joju-Matsumoto/vscode-workspace/vscodeworkspace"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list workspaces",
	Long:    `list workspaces`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		sortBy, err := cmd.Flags().GetString("sort")
		if err != nil {
			return err
		}

		wss, err := usecase.ListWorkspace(vscodeworkspace.ListWorkspaceUsecaseOption{
			SortBy: sortBy,
		})
		if err != nil {
			return err
		}

		if len(wss) == 0 {
			fmt.Fprintln(os.Stdout, "No workspace found")
			return nil
		}

		ShowWorkspaces(wss...)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	sortBy := NewEnum([]string{"name", "opened_at", "count"}, "opened_at")
	listCmd.Flags().VarP(sortBy, "sort", "s", fmt.Sprintf("sort by [%s]", strings.Join(sortBy.Allowed, ",")))
}

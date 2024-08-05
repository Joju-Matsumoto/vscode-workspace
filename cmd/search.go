/*
Copyright © 2024 jj_ <jj_@team-jj.com>
*/
package cmd

import (
	"github.com/Joju-Matsumoto/vscode-workspace/vscodeworkspace"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [BASE DIRECTORY]...",
	Short: "search workspaces",
	Long: `search workspaces

  With no [BASE DIRECTOR], current directory is used.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dirs := args
		if len(dirs) == 0 {
			dirs = append(dirs, ".")
		}

		usecase, err := NewUsecase(baseFile)
		if err != nil {
			return err
		}

		workspaces := make([]vscodeworkspace.Workspace, 0)

		for _, dir := range dirs {
			wss, err := usecase.SearchWorkspaceFromDirectory(dir, "")
			if err != nil {
				return err
			}
			workspaces = append(workspaces, wss...)
		}

		ShowWorkspaces(workspaces...)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

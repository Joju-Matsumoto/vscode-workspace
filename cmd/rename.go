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

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename <OLD NAME> <NEW NAME>",
	Short: "rename workspace",
	Long:  `rename workspace`,
	Args:  cobra.ExactArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 最初の引数以外は補完しない
		if len(args) >= 1 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		wss, err := usecase.ListWorkspace(vscodeworkspace.ListWorkspaceUsecaseOption{})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, cobra.ShellCompDirectiveError
		}

		vargs := ListWorkspaceNamesWithDescription(wss)

		return vargs, cobra.ShellCompDirectiveDefault
	},
	Run: func(cmd *cobra.Command, args []string) {
		oldName := args[0]
		newName := args[1]

		if err := usecase.RenameWorkspace(oldName, newName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

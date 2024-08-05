/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [BASE DIRECTORY]",
	Short: "search workspaces",
	Long:  `search workspaces`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no base directory specified")
		}
		baseDir := args[0]

		usecase, err := NewUsecase(baseFile)
		if err != nil {
			return err
		}

		wss, err := usecase.SearchWorkspaceFromDirectory(baseDir, "")
		if err != nil {
			return err
		}

		ShowWorkspaces(wss...)

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

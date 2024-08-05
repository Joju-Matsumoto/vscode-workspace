/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [NAME]...",
	Short: "remove workspace from the list",
	Long:  `remove workspace from the list`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		names := args

		usecase, err := NewUsecase(baseFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		for _, name := range names {
			if err := usecase.DeleteWorkspace(name); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

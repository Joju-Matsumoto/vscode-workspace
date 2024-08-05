/*
Copyright © 2024 jj_ <jj_@team-jj.com>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var baseFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vscode-workspace",
	Short: "A VSCode Workspace CLI tool",
	Long:  `A VSCode Workspace CLI tool`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	defaultBaseFile := filepath.Join(userHomeDir, ".vscodeworkspaces.json")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vscode-workspace.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&baseFile, "file", "f", defaultBaseFile, "File path to save the list of workspaces")
}

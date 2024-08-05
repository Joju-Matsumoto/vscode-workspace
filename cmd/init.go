/*
Copyright © 2024 jj_ <jj_@team-jj.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [ROOT DIRECTORY]",
	Short: "init workspace",
	Long: `init workspace

With no [ROOT DIRECTORY], current directory is used`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := func() error {
			dir := "."
			if len(args) > 0 {
				dir = args[0]
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			usecase, err := NewUsecase(baseFile)
			if err != nil {
				return err
			}

			ws, err := usecase.InitWorkspace(name, dir)
			if err != nil {
				return err
			}

			ShowWorkspaces(ws)

			return nil
		}(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCmd.Flags().StringP("name", "n", "", "name of the workspace. If empty, the directory name is used.")
}

package cmd

import (
	"os"

	"github.com/mheob/update-mac/cmdutil"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "update-mac",
	Version: "1.0.0",
	Short:   "Update applications on the macOS system",
	Long: `Update applications on the macOS system via command line tools.
You can use this tool to update applications via brew, composer, npm and so on.`,
	Run: func(cmd *cobra.Command, args []string) {
		hasFlag := false
		if cmd.Flags().Lookup("brew").Value.String() == "true" {
			hasFlag = updateBrew()
		}
		if cmd.Flags().Lookup("npm").Value.String() == "true" {
			hasFlag = updateNpm()
		}
		if cmd.Flags().Lookup("omz").Value.String() == "true" {
			hasFlag = updateOmz()
		}
		if !hasFlag {
			updateBrew()
			updateOmz()
			updateNpm()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("brew", "b", false, "update brew and its applications")
	rootCmd.Flags().BoolP("npm", "n", false, "check for updates of global npm packages")
	rootCmd.Flags().BoolP("omz", "o", false, "update Oh My Zsh")
}

func updateBrew() bool {
	cmdutil.PrintUpdateStart("brew")
	cmdutil.CallCmd("brew", "update")
	cmdutil.CallCmd("brew", "upgrade")
	cmdutil.CallCmd("brew", "cu", "--all", "--yes", "--cleanup")
	cmdutil.CallCmd("brew", "update")
	cmdutil.PrintUpdateEnd("brew")
	return true
}

func updateNpm() bool {
	cmdutil.PrintNpmOutdatedStart("npm")
	cmdutil.CallCmd("npm", "-g", "outdated")
	cmdutil.PrintNpmOutdatedEnd("npm")
	return true
}

func updateOmz() bool {
	cmdutil.PrintUpdateStart("omz")
	cmdutil.CallCmd("sh", "-c", "~/.oh-my-zsh/tools/upgrade.sh")
	cmdutil.PrintUpdateEnd("omz")
	return true
}

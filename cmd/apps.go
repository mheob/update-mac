package cmd

import (
	"fmt"

	"github.com/mheob/update-mac/cmdutil"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Update applications on the macOS system",
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

func init() {
	rootCmd.AddCommand(appsCmd)
	appsCmd.Flags().BoolP("brew", "b", false, "update brew and its applications")
	appsCmd.Flags().BoolP("npm", "n", false, "check for updates of global npm packages")
	appsCmd.Flags().BoolP("omz", "o", false, "update Oh My Zsh")
}

func updateBrew() bool {
	cmdutil.PrintCommandStart("brew")

	err := cmdutil.CallCmd("brew", "update")
	if err != nil {
		panic(err)
	}

	err = cmdutil.CallCmd("brew", "upgrade")
	if err != nil {
		panic(err)
	}

	err = cmdutil.CallCmd("brew", "cu", "--all", "--yes", "--cleanup")
	if err != nil {
		panic(err)
	}

	err = cmdutil.CallCmd("brew", "update")
	if err != nil {
		panic(err)
	}

	cmdutil.PrintCommandEnd("brew")

	return true
}

func updateNpm() bool {
	fmt.Printf(
		"%sCheck global %s for updates starting ...%s\n",
		cmdutil.Color["purple"],
		"npm",
		cmdutil.Color["default"],
	)

	err := cmdutil.CallCmd("npm", "-g", "outdated")
	if err != nil && err.Error() != "exit status 1" {
		panic(err)
	}

	fmt.Printf(
		"%sCheck global %s for updates finished%s\n\n",
		cmdutil.Color["purple"],
		"npm",
		cmdutil.Color["default"],
	)

	return true
}

func updateOmz() bool {
	cmdutil.PrintCommandStart("omz")

	err := cmdutil.CallCmd("sh", "-c", "~/.oh-my-zsh/tools/upgrade.sh")
	if err != nil {
		panic(err)
	}

	cmdutil.PrintCommandEnd("omz")

	return true
}

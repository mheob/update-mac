package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "update-mac",
	Version: "1.0.0",
	Short:   "Update applications or Node/NPM on the macOS system",
	Long:    "Update applications or Node/NPM on the macOS system via command line tools.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.AddCommand(apps.Cmd)
}

package cmdutil

import (
	"os"
	"os/exec"
)

// `CallCmd` calls the command with the given name and arguments
func CallCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() >= getFirstErrorCode(args) {
				panic(err)
			}
		}
	}
}

// `getFirstErrorCode` returns the number of first available error code.
//
// If the command to run is `npm outdated` it returns `2` otherwise `1`.
func getFirstErrorCode(args []string) int {
	if args[len(args)-1] == "outdated" {
		return 2
	}
	return 1
}

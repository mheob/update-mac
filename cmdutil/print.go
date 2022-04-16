package cmdutil

import "fmt"

var Color = map[string]string{
	"default": "\033[0m",
	"purple":  "\033[35m",
}

// `PrintUpdateStart` prints the start of the update command
func PrintUpdateStart(command string) {
	fmt.Printf("%sUpdate %s starting ...%s\n", Color["purple"], command, Color["default"])
}

// `PrintUpdateEnd` prints the end of the update command
func PrintUpdateEnd(command string) {
	fmt.Printf("%sUpdate %s finished%s\n\n", Color["purple"], command, Color["default"])
}

// `PrintNpmOutdatedStart` prints the start of the npm outdated command
func PrintNpmOutdatedStart(command string) {
	fmt.Printf(
		"%sCheck global %s for updates starting%s\n\n", Color["purple"], command, Color["default"],
	)
}

// `PrintNpmOutdatedEnd` prints the end of the npm outdated command
func PrintNpmOutdatedEnd(command string) {
	fmt.Printf(
		"%sCheck global %s for updates finished%s\n\n", Color["purple"], command, Color["default"],
	)
}

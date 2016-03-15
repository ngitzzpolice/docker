package libcontainerd

import (
	"strings"
	"syscall"

	"github.com/Sirupsen/logrus"
)

// setupEnvironmentVariables convert a string array of environment variables
// into a map as required by the HCS. Source array is in format [v1=k1] [v2=k2] etc.
func setupEnvironmentVariables(a []string) map[string]string {
	r := make(map[string]string)
	for _, s := range a {
		arr := strings.Split(s, "=")
		if len(arr) == 2 {
			r[arr[0]] = arr[1]
		}
	}
	return r
}

func createCommandLine(args []string) string {
	escapedArgs := make([]string, len(args))
	// Convert the args array into the escaped command line.
	for i, arg := range args {
		escapedArgs[i] = syscall.EscapeArg(arg)
	}
	commandLine := strings.Join(escapedArgs, " ")
	logrus.Debugf("commandLine: %s", commandLine)

	return commandLine
}

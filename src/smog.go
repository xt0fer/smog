package smog

import (
	"fmt"
	"os"
)

// variables to hold the separator characters
var pathSeparator string
var fileSeparator string
var arguments []string
var universe Universe

func Smog(args []string) int {
	// Setup the system path and file separators

	pathSeparator = string(os.PathListSeparator)
	fileSeparator = string(os.PathSeparator)

	// Check for command line switches
	arguments := handleArguments(arguments)

	// Initialize the known universe
	universe = *newUniverse().Initialize(arguments)

	// Exit with error code 0
	fmt.Println("Exiting Smog")

	return 0
}

func handleArguments(arguments []string) []string {
	// Check for command line switches
	// This function will return a list of arguments
	// that are not switches
	return arguments
}

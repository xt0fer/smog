package main

import (
	"fmt"
	"os"

	"github.com/xtofer/src/smog"
)

func main() {
	args := os.Args
	fmt.Println("Hello, Smog!")
	// Use args here
	os.Exit(smog.Smog(args))
}

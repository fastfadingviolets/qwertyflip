package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command file> <input file>\n", os.Args[0])
		os.Exit(1)
	}
	// commandFile, inputFile := os.Args[1], os.Args[2]
}

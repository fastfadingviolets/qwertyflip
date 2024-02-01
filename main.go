package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fastfadingviolets/qwertyflip/flip"
)

func getTransform(flipper *flip.Flipper, fromFile string) (flip.Transform, error) {
	cmdFile, err := os.Open(fromFile)
	if err != nil {
		return nil, fmt.Errorf("Unable to open command file %s: %v", fromFile, err)
	}
	defer cmdFile.Close()
	bytes, err := io.ReadAll(cmdFile)
	if err != nil {
		return nil, fmt.Errorf("Unable to read command file %s: %v", fromFile, err)
	}
	cmdString := strings.TrimSpace(string(bytes))
	return flipper.ParseCommand(cmdString)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command file> <input file>\n", os.Args[0])
		os.Exit(1)
	}
	commandFile, inputFile := os.Args[1], os.Args[2]
	flipper := flip.NewFlipper()
	transform, err := getTransform(flipper, commandFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open input file %s: %v", inputFile, err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := transform.Apply(scanner.Text())
		fmt.Println(line)
	}
}

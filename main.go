package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fastfadingviolets/qwertyflip/flip"
)

func getCommand(fromFile string) (string, error) {
	cmdFile, err := os.Open(fromFile)
	if err != nil {
		return "", fmt.Errorf("Unable to open command file %s: %v", fromFile, err)
	}
	defer cmdFile.Close()
	bytes, err := io.ReadAll(cmdFile)
	if err != nil {
		return "", fmt.Errorf("Unable to read command file %s: %v", fromFile, err)
	}
	return strings.TrimSpace(string(bytes)), nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command file> <input file>\n", os.Args[0])
		os.Exit(1)
	}
	commandFile, inputFile := os.Args[1], os.Args[2]
	command, err := getCommand(commandFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	flipper := flip.NewFlipper()
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open input file %s: %v", inputFile, err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := flipper.RunCommand(command, scanner.Text())
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(line)
	}
}

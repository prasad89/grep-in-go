package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func grep(filename, pattern string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", filename, err)
	}
	defer file.Close()

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("invalid pattern %q: %w", pattern, err)
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if regex.MatchString(line) {
			fmt.Printf("%d: %s\n", lineNumber, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: grep <file> <pattern>")
		os.Exit(1)
	}
	filename, pattern := os.Args[1], os.Args[2]
	if err := grep(filename, pattern); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check that the user gave us something to print
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . \"Hello\"")
		os.Exit(1)
	}

	// Grab the text the user typed
	input := os.Args[1]

	// If empty string, print nothing
	if input == "" {
		return
	}

	// Choose banner file (default is standard)
	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	// Load the banner file
	banner, err := loadBanner(bannerName + ".txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Replace literal \n with a real newline
	input = strings.ReplaceAll(input, `\n`, "\n")

	// Split into lines and print each one
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			fmt.Println()
		} else {
			printLine(line, banner)
		}
	}
}

func loadBanner(filename string) (map[rune][]string, error) {
	// Read the whole file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %q: %w", filename, err)
	}

	content := string(data)

	// Fix Windows line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")

	// Split into blocks — one block per character
	blocks := strings.Split(content, "\n\n")

	// Build the map: character → its 8 lines
	banner := make(map[rune][]string)
	startChar := rune(32) // space character

	for i, block := range blocks {
		ch := startChar + rune(i)
		lines := strings.Split(block, "\n")
		if len(lines) >= 8 {
			banner[ch] = lines[:8]
		}
	}

	return banner, nil
}

func printLine(text string, banner map[rune][]string) {
	// Print row 0 of all characters, then row 1, etc.
	for row := 0; row < 8; row++ {
		for _, ch := range text {
			lines, found := banner[ch]
			if !found {
				fmt.Print("  ")
				continue
			}
			fmt.Print(lines[row])
		}
		fmt.Println()
	}
}
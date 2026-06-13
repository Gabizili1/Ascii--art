package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// TestLoadBanner verifies that banner files are parsed correctly into 8-line character slices.
func TestLoadBanner(t *testing.T) {
	// Create a temporary mock banner file
	mockContent := "line1\nline2\nline3\nline4\nline5\nline6\nline7\nline8\n\nlineA\nlineB\nlineC\nlineD\nlineE\nlineF\nlineG\nlineH"
	tmpFile, err := os.CreateTemp("", "mock_banner_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test

	if _, err := tmpFile.WriteString(mockContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Run the loader function
	banner, err := loadBanner(tmpFile.Name())
	if err != nil {
		t.Fatalf("loadBanner returned an unexpected error: %v", err)
	}

	// Verify the space character (index 0, rune 32)
	spaceLines, found := banner[rune(32)]
	if !found {
		t.Fatalf("Expected to find character rune 32 (' ') in map")
	}
	if len(spaceLines) != 8 || spaceLines[0] != "line1" {
		t.Errorf("Unexpected lines for space character: %v", spaceLines)
	}

	// Verify the exclamation mark character (index 1, rune 33)
	exclLines, found := banner[rune(33)]
	if !found {
		t.Fatalf("Expected to find character rune 33 ('!') in map")
	}
	if len(exclLines) != 8 || exclLines[7] != "lineH" {
		t.Errorf("Unexpected lines for '!' character: %v", exclLines)
	}
}

// TestPrintLine captures stdout to verify that rows are printed side-by-side correctly.
func TestPrintLine(t *testing.T) {
	// Build a mock banner map for characters 'A' and 'B'
	mockBanner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
		'B': {"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8"},
	}

	// Redirect standard output to a buffer to capture the print results
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the print function
	printLine("AB", mockBanner)

	// Close writer and restore original stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Expected output should combine row 0, then row 1, etc., followed by newlines
	expectedLines := []string{
		"A1B1", "A2B2", "A3B3", "A4B4", "A5B5", "A6B6", "A7B7", "A8B8", "",
	}
	expectedOutput := strings.Join(expectedLines, "\n")

	if output != expectedOutput {
		t.Errorf("Output mismatch.\nExpected:\n%q\nGot:\n%q", expectedOutput, output)
	}
}

// TestMissingFile verifies that loadBanner returns an error if the file doesn't exist.
func TestMissingFile(t *testing.T) {
	_, err := loadBanner("non_existent_file_xyz.txt")
	if err == nil {
		t.Error("Expected an error when loading a non-existent file, but got nil")
	}
}

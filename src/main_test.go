package main

import (
	"testing"
)

// TODO: Add TestCountLinesOfCode function

func TestCountLinesOfFile(t *testing.T) {

	tests := []struct {
		filePath    string
		expectedLines int
	}{
		{"../tests/data/hello.go", 4},
		{"../tests/data/stack.ml", 28},
	}

	for _, tt := range tests {
		lines, err := countLinesOfFile(tt.filePath)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if lines != tt.expectedLines {
			t.Errorf("Expected %d lines, but got %d", tt.expectedLines, lines)
		}
	}
}
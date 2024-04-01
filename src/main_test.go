package main

import (
	"reflect"
	"testing"
)

func TestCountLinesOfCode(t *testing.T) {

	tests := []struct {
		filePath        string
		expectedClocMap clocMap
	}{
		{"../tests/data/hello.go", clocMap{"Go": 4}},
		{"../tests/data/stack.ml", clocMap{"OCaml": 28}},
	}

	for _, tt := range tests {
		clocMap, err := countLinesOfCode(tt.filePath)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if !reflect.DeepEqual(clocMap, tt.expectedClocMap) {
			t.Errorf("Expected %v, but got %v", tt.expectedClocMap, clocMap)
		}
	}
}

func TestCountLinesOfFile(t *testing.T) {

	tests := []struct {
		filePath      string
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

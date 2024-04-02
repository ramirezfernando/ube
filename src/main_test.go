package main

import (
	"cloc-tool/src/terminal"
	"reflect"
	"testing"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func TestGetMessage(t *testing.T) {
	lines := clocMap{
		"Go":     4,
		"OCaml": 28,
		"Plain Text": 0,
		"JavaScript": 24,
	}

	_, err := countLinesOfCode("../tests/data/nonexistent.txt")

	tests := []struct {
		filePath         string
		expectedMessage  terminal.ClocCompleted
	}{
		{"../tests/data", terminal.ClocCompleted{Table: generateTable(lines), Help: help.New()}},
		{"../tests/data/nonexistent.txt", terminal.ClocCompleted{Err: err}}, // Non-existent file

	}
	
	for _, tt := range tests {
		message := getMessage(tt.filePath)
		if !reflect.DeepEqual(message, tt.expectedMessage) {
			t.Errorf("Expected %v, but got %v", tt.expectedMessage, message)
		}
	}
}

func TestCountLinesOfCode(t *testing.T) {

	tests := []struct {
		filePath        string
		expectedClocMap clocMap
	}{
		{"../tests/data/hello.go", clocMap{"Go": 4}},
		{"../tests/data/stack.ml", clocMap{"OCaml": 28}},
		{"../tests/data/empty.txt", clocMap{"Plain Text": 0}},
		{"../tests/data/person.js", clocMap{"JavaScript": 24}},
		{"../tests/data", clocMap{"Go": 4, "OCaml": 28, "Plain Text": 0, "JavaScript": 24}}, // Directory
		{"../tests/data/nonexistent.txt", clocMap{}}, // Non-existent file
	}

	for _, tt := range tests {
		clocMap, _ := countLinesOfCode(tt.filePath)

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
		{"../tests/data/empty.txt", 0},
		{"../tests/data/person.js", 24},
		{"../tests/data/nonexistent.txt", -1}, // Non-existent file
	}

	for _, tt := range tests {
		lines, _ := countLinesOfFile(tt.filePath)

		if lines != tt.expectedLines {
			t.Errorf("Expected %d lines, but got %d", tt.expectedLines, lines)
		}
	}
}

func TestGenerateTable(t *testing.T) {
	lines := clocMap{
		"Go":     100,
		"Python": 200,
		"Java":   150,
	}

	expectedColumns := []table.Column{
		{Title: "Language", Width: 15},
		{Title: "Lines of Code", Width: 15},
	}

	expectedRows := []table.Row{
		{"Python", "200"},
		{"Java", "150"},
		{"Go", "100"},
		{"", ""},
		{"Total", "450"},
	}

	expectedTable := table.New(
		table.WithColumns(expectedColumns),
		table.WithRows(expectedRows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	// lipgloss style
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("#5433ff")).
		Bold(false)
	expectedTable.SetStyles(s)

	actualTable := generateTable(lines)

	if !reflect.DeepEqual(expectedTable, actualTable) {
		t.Errorf("Expected %v, but got %v", expectedTable, actualTable)
	}
}
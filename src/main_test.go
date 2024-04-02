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
	}

	tests := []struct {
		filePath         string
		expectedMessage  terminal.ClocCompleted
	}{
		{"../tests/data", terminal.ClocCompleted{Table: generateTable(lines), Help: help.New()}},
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
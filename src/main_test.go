package main

import (
	"ube/src/terminal"
	"reflect"
	"testing"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func TestGetMessage(t *testing.T) {
	llc := map[string]languageDetails{
		"Go":         {4, 1},
		"OCaml":      {28, 1},
		"Plain Text": {0, 1},
		"JavaScript": {24, 1},
	}

	_, err := countLinesOfCode("../tests/data/nonexistent.txt")

	tests := []struct {
		path            string
		expectedMessage terminal.ClocCompleted
	}{
		{"../tests/data", terminal.ClocCompleted{Table: generateTable(llc), Help: help.New()}},
		{"../tests/data/nonexistent.txt", terminal.ClocCompleted{Err: err}}, // Non-existent file

	}

	for _, tt := range tests {
		message := getMessage(tt.path)
		if !reflect.DeepEqual(message, tt.expectedMessage) {
			t.Errorf("Expected %v, but got %v", tt.expectedMessage, message)
		}
	}
}

func TestCountLinesOfCode(t *testing.T) {

	tests := []struct {
		path            string
		expectedClocMap map[string]languageDetails
	}{
		{"../tests/data/hello.go", map[string]languageDetails{"Go": {4, 1}}},
		{"../tests/data/stack.ml", map[string]languageDetails{"OCaml": {28, 1}}},
		{"../tests/data/empty.txt", map[string]languageDetails{"Plain Text": {0, 1}}},
		{"../tests/data/person.js", map[string]languageDetails{"JavaScript": {24, 1}}},
		{"../tests/data", map[string]languageDetails{"Go": {4, 1}, "OCaml": {28, 1}, "Plain Text": {0, 1}, "JavaScript": {24, 1}}}, // Directory
		{"../tests/data/nonexistent.txt", map[string]languageDetails{}},                                                            // Non-existent file
		{"../tests/data/unsupported.xyz", map[string]languageDetails{}},                                                            // Unsupported file extension

	}

	for _, tt := range tests {
		clocMap, _ := countLinesOfCode(tt.path)

		if !reflect.DeepEqual(clocMap, tt.expectedClocMap) {
			t.Errorf("Expected %v, but got %v", tt.expectedClocMap, clocMap)
		}
	}
}

func TestCountLinesOfFile(t *testing.T) {
	tests := []struct {
		path          string
		expectedLines int
	}{
		{"../tests/data/hello.go", 4},
		{"../tests/data/stack.ml", 28},
		{"../tests/data/empty.txt", 0},
		{"../tests/data/person.js", 24},
		{"../tests/data/nonexistent.txt", -1}, // Non-existent file
	}

	for _, tt := range tests {
		lines, _ := countLinesOfFile(tt.path)

		if lines != tt.expectedLines {
			t.Errorf("Expected %d lines, but got %d", tt.expectedLines, lines)
		}
	}
}

func TestGenerateTable(t *testing.T) {
	llc := map[string]languageDetails{
		"Go":     {100, 1},
		"Python": {200, 1},
		"Java":   {150, 1},
	}

	expectedColumns := []table.Column{
		{Title: "Language", Width: 16},
		{Title: "Lines", Width: 16},
		{Title: "Files", Width: 10},
	}

	expectedRows := []table.Row{
		{"Python", "200", "1"},
		{"Java", "150", "1"},
		{"Go", "100", "1"},
		{"", "", ""},
		{"Total", "450", "3"},
	}

	expectedTable := table.New(
		table.WithColumns(expectedColumns),
		table.WithRows(expectedRows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

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

	actualTable := generateTable(llc)

	if !reflect.DeepEqual(expectedTable, actualTable) {
		t.Errorf("Expected %v, but got %v", expectedTable, actualTable)
	}
}

func TestFormatStringInteger(t *testing.T) {
	tests := []struct {
		number         string
		expectedNumber string
	}{
		{"1234", "1,234"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"0", "0"},
	}
	
	for _, tt := range tests {
		num := FormatStringInteger(tt.number)

		if num != tt.expectedNumber {
			t.Errorf("Expected %s lines, but got %s", tt.expectedNumber, num)
		}
	} 
}

package main

import (
	"cloc-tool/src/language"
	"cloc-tool/src/terminal"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type clocMap = map[string]int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cloc <folder>")
		os.Exit(1)
	}

	folderPath := os.Args[1]

	lines, err := countLinesOfCode(folderPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

    t := generateTable(lines)

	m := terminal.Model{ExecutionTime: stopwatch.NewWithInterval(time.Millisecond), IsRunning: true, Table: t, Help: help.New()}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func countLinesOfCode(folderPath string) (clocMap, error) {
	cloc := make(clocMap)
	var fc = 0
	err := filepath.WalkDir(folderPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Type().IsRegular() {
			if val, exists := language.Exts[filepath.Ext(path)]; exists {
				lines, err := countLinesOfFile(path)
				if err != nil {
					return err
				}
				cloc[val] += lines
				fc += 1
			}
		}

		return nil
	})
	if err != nil {
		return cloc, err
	}
	return cloc, nil
}

func countLinesOfFile(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	return len(lines), nil
}

func generateTable(lines clocMap) table.Model {
    columns := []table.Column{
		{Title: "Language", Width: 15},
		{Title: "Lines of Code", Width: 15},
	}

	rows := []table.Row{}
	for lang, li := range lines {
		rows = append(rows, table.Row{lang, strconv.Itoa(li)})
	}
	sort.Slice(rows, func(i, j int) bool {
		li1, _ := strconv.Atoi(rows[i][1])
		li2, _ := strconv.Atoi(rows[j][1])
		return li1 > li2
	})

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
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
	t.SetStyles(s)

    return t
}
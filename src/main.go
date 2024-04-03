package main

import (
	"bytes"
	"cloc-tool/src/language"
	"cloc-tool/src/terminal"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// LanguageLineCount is a map of programming languages to line count
// Example: { "Python": 100, "Go": 200 }
type LanguageLineCount = map[string]int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cloc <folder>")
		os.Exit(1)
	}

	path := os.Args[1]

	m := terminal.Model{ExecutionTime: stopwatch.NewWithInterval(time.Millisecond), IsRunning: true}
	p := tea.NewProgram(m)

	go func() {
		msg := getMessage(path)
		p.Send(msg)
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func getMessage(path string) tea.Msg {
	llc, err := countLinesOfCode(path)
	if err != nil {
		fmt.Println("Error running program:", err)
		return terminal.ClocCompleted{Err: err}
	}

	t := generateTable(llc)
	h := help.New()
	return terminal.ClocCompleted{Table: t, Help: h}
}

func countLinesOfCode(path string) (LanguageLineCount, error) {
	llc := make(LanguageLineCount)
	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Type().IsRegular() {
			if language, exists := language.Exts[filepath.Ext(path)]; exists {
				lines, err := countLinesOfFile(path)
				if err != nil {
					return err
				}
				llc[language] += lines
			}
		}

		return nil
	})
	if err != nil {
		return llc, err
	}
	return llc, nil
}

// Counts the number of '\n' characters in a file
func countLinesOfFile(filePath string) (int, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	var count int
	var read int
	var target []byte = []byte("\n")

	buffer := make([]byte, 32*1024)

	for {
		read, err = file.Read(buffer)
		if err != nil {
			break
		}

		count += bytes.Count(buffer[:read], target)
	}

	if err == io.EOF {
		return count, nil
	}

	return count, err
}

func generateTable(llc LanguageLineCount) table.Model {
	columns := []table.Column{
		{Title: "Language", Width: 16},
		{Title: "Lines of Code", Width: 16},
	}

	rows := []table.Row{}
	total := 0
	for language, lineCount := range llc {
		rows = append(rows, table.Row{language, strconv.Itoa(lineCount)})
		total += lineCount
	}
	sort.Slice(rows, func(i, j int) bool {
		li1, _ := strconv.Atoi(rows[i][1])
		li2, _ := strconv.Atoi(rows[j][1])
		return li1 > li2
	})

	rows = append(rows, table.Row{"", ""})
	rows = append(rows, table.Row{"Total", strconv.Itoa(total)})

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

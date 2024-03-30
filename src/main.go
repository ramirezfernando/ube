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

type clocMap = map[string]int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cloc <folder>")
		os.Exit(1)
	}

	folderPath := os.Args[1]

	m := terminal.Model{ExecutionTime: stopwatch.NewWithInterval(time.Millisecond), IsRunning: true}
	p := tea.NewProgram(m)

	go func() {
		msg := getMessage(folderPath)
		p.Send(msg)
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func getMessage(folderPath string) tea.Msg {
	lines, err := countLinesOfCode(folderPath)
	if err != nil {
		return terminal.ClocCompleted{Err: err}
	}

	t := generateTable(lines)
	h := help.New()
	return terminal.ClocCompleted{Table: t, Help: h}
}

func countLinesOfCode(folderPath string) (clocMap, error) {
	cloc := make(clocMap)
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
			}
		}

		return nil
	})
	if err != nil {
		return cloc, err
	}
	return cloc, nil
}

func countLines(r io.Reader) int {
	count, err := CountLines(r)
	if err != nil {
		fmt.Println("Error running program:", err)
	}
    return count
}

func CountLines(r io.Reader) (int, error) {
	var count int
	var read int
	var err error
	var target []byte = []byte("\n")

	buffer := make([]byte, 32*1024)

	for {
		read, err = r.Read(buffer)
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

func countLinesOfFile(filename string) (int, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0444)
	if err != nil {
		return 0, err
	}
    defer file.Close()
    return countLines(file), err
}

func generateTable(lines clocMap) table.Model {
	columns := []table.Column{
		{Title: "Language", Width: 15},
		{Title: "Lines of Code", Width: 15},
	}

	rows := []table.Row{}
    total := 0
	for lang, li := range lines {
		rows = append(rows, table.Row{lang, strconv.Itoa(li)})
        total += li
	}
	sort.Slice(rows, func(i, j int) bool {
		li1, _ := strconv.Atoi(rows[i][1])
		li2, _ := strconv.Atoi(rows[j][1])
		return li1 > li2
	})

    totalRow := table.Row{"Total", strconv.Itoa(total)}
    rows = append(rows, totalRow)

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

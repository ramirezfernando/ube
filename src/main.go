package main

import (
	"bytes"
	"cloc-tool/src/language"
	"cloc-tool/src/terminal"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(`Usage: cloc <folder>`)
		os.Exit(1)
	} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		log.Fatal("cloc: no such file or directory")
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
		log.Fatal(err)
	}

	os.Exit(0)
}

func getMessage(path string) tea.Msg {
	llc, err := countLinesOfCode(path)
	if err != nil {
		log.Error(err)
		return terminal.ClocCompleted{Err: err}
	}

	t := generateTable(llc)
	h := help.New()
	return terminal.ClocCompleted{Table: t, Help: h}
}

func countLinesOfCode(path string) (map[string]int, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	lineCount := make(map[string]int)

	err := filepath.WalkDir(path, func(currPath string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !info.Type().IsRegular(){
			return nil
		}
		language, exists := language.Exts[filepath.Ext(currPath)];
		if !exists {
			return nil
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			lines, err := countLinesOfFile(currPath)
			if err != nil {
				log.Error("Error counting lines in %s: %v\n", currPath, err)
				return
			}

			mu.Lock()
			lineCount[language] += lines
			mu.Unlock()
		}()

		return nil
	})

	if err != nil {
		return lineCount, err
	}

	wg.Wait()
	return lineCount, nil
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

func generateTable(lineCount map[string]int) table.Model {
	columns := []table.Column{
		{Title: "Language", Width: 16},
		{Title: "Lines of Code", Width: 16},
	}

	rows := []table.Row{}
	total := 0
	for language, count := range lineCount {
		rows = append(rows, table.Row{language, strconv.Itoa(count)})
		total += count
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

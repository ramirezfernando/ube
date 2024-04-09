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

type languageDetails struct {
	// TODO: Change to uint
	lines int
	files int
}

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

func countLinesOfCode(path string) (map[string]languageDetails, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	lineCount := make(map[string]languageDetails)

	err := filepath.WalkDir(path, func(currPath string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !info.Type().IsRegular() {
			return nil
		}
		language, exists := language.Exts[filepath.Ext(currPath)]
		if !exists {
			return nil
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			lines, err := countLinesOfFile(currPath)
			if err != nil {
				log.Error(err)
				return
			}

			mu.Lock()
			ld, exists := lineCount[language]
			if !exists {
				ld = languageDetails{}
			}
			ld.lines += lines
			ld.files++
			lineCount[language] = ld
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

func generateTable(lineCount map[string]languageDetails) table.Model {
	columns := []table.Column{
		{Title: "Language", Width: 16},
		{Title: "Lines", Width: 16},
		{Title: "Files", Width: 10},
	}

	rows := []table.Row{}
	lineTotal := 0
	fileTotal := 0
	for language, details := range lineCount {
		rows = append(rows, table.Row{language, strconv.Itoa(details.lines), strconv.Itoa(details.files)})
		lineTotal += details.lines
		fileTotal += details.files
	}
	sort.Slice(rows, func(i, j int) bool {
		li1, _ := strconv.Atoi(rows[i][1])
		li2, _ := strconv.Atoi(rows[j][1])
		return li1 > li2
	})

	rows = append(rows, table.Row{"", "", ""})
	rows = append(rows, table.Row{"Total", strconv.Itoa(lineTotal), strconv.Itoa(fileTotal)})

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

package main

import (
	"bytes"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"
	"ube/src/language"
	"ube/src/terminal"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type languageDetails struct {
	lines int
	files int
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(`Usage: ube <folder>`)
		os.Exit(1)
	} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		log.Fatal("ube: no such file or directory")
		os.Exit(1)
	}

	path := os.Args[1]

	m := terminal.Model{ElapsedTime: stopwatch.NewWithInterval(time.Millisecond), IsRunning: true}
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
	stats, err := countLinesOfCode(path)
	if err != nil {
		log.Error(err)
		return terminal.CompletionResult{Err: err}
	}

	t := generateTable(stats)
	h := help.New()
	return terminal.CompletionResult{Table: t, Help: h}
}

func isValidFile(fileName string, info fs.DirEntry) bool {
	if info.IsDir() || !info.Type().IsRegular() {
		return false
	}
	_, exists := getLanguageName(fileName)
	return exists
}

func getLanguageName(fileName string) (string, bool) {
	language, exists := language.Exts[getLanguageExtension(fileName)]
	if !exists {
		return "", false
	}
	return language, true
}

func getLanguageExtension(fileName string) string {
	return filepath.Ext(fileName)
}

func countLinesOfCode(path string) (map[string]languageDetails, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	lineCount := make(map[string]languageDetails)

	err := filepath.WalkDir(path, func(currPath string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !isValidFile(currPath, info) {
			return nil
		}
		language, _ := getLanguageName(currPath)

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

	// Sort by lines of code
	sort.Slice(rows, func(i, j int) bool {
		li1, _ := strconv.Atoi(rows[i][1])
		li2, _ := strconv.Atoi(rows[j][1])
		return li1 > li2
	})

	for i, row := range rows {
		rows[i][1] = FormatStringInteger(row[1])
		rows[i][2] = FormatStringInteger(row[2])
	}

	rows = append(rows, table.Row{"", "", ""})
	rows = append(rows, table.Row{"Total", FormatStringInteger(strconv.Itoa(lineTotal)), FormatStringInteger(strconv.Itoa(fileTotal))})

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

func FormatStringInteger(n string) string {
	if len(n) < 4 {
		return n
	}

	var formatted string
	for i, r := range n {
		if i != 0 && (len(n)-i)%3 == 0 {
			formatted += ","
		}
		formatted += string(r)
	}

	return formatted
}

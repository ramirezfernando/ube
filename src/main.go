package main

import (
	"cloc-tool/src/language"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
    "time"
)

type clocMap = map[string]int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cloc <folder>")
		os.Exit(1)
	}
	folderPath := os.Args[1]

    startTime := time.Now()
	lines, fileCount, err := countLinesOfCode(folderPath)
    endTime := time.Now()
    executionTime := endTime.Sub(startTime)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for lang, li := range lines {
		fmt.Printf("Total lines of code for %q: %d\n", lang, li)
	}
    fmt.Printf("Total files read: %d\n", fileCount)
    fmt.Printf("Executed in %v\n", executionTime)

}

func countLinesOfCode(folderPath string) (clocMap, int, error) {
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
		return cloc, fc, err
	}
	return cloc, fc, nil
}

func countLinesOfFile(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	return len(lines), nil
}

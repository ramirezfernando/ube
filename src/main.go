package main

import (
	"cloc-tool/src/language"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

    for lang, li := range lines {
        fmt.Printf("Total lines of code for %q: %d\n", lang, li)
    }
}

func countLinesOfCode(folderPath string) (clocMap, error) {
    cloc := make(clocMap)
    err := filepath.WalkDir(folderPath, func(path string, info fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && info.Type().IsRegular() {
            val, exists := language.Exts[filepath.Ext(path)]; if exists {
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

func countLinesOfFile(filename string) (int, error) {
    content, err := os.ReadFile(filename)
    if err != nil {
        return 0, err
    }
    lines := strings.Split(string(content), "\n")
    return len(lines), nil
}
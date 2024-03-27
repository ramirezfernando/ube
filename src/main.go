package main

/*
Idea for this tool:
1. run "cloc ."
2. count lines of code within each file of
3. display cloc info using bubble tea and bubbles
*/

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "cloc-tool/src/language"
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
// I got Error: read ../../code-coogs/bot/node_modules/discord.js: is a directory, why?
// I need to add a check for directories in countLinesOfFile
func countLinesOfCode(folderPath string) (clocMap, error) {
    cloc := make(clocMap)
    err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            val, exists := language.Exts[filepath.Ext(path)]
            if exists {
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
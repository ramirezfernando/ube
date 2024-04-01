# cloc-tool

A command line tool to count lines of code (CLOC), written in [Go](https://go.dev/) & styled with [bubbletea](https://github.com/charmbracelet/bubbletea).

## Example
<img width="508" alt="Screen Shot 2024-03-31 at 1 26 14 AM" src="https://github.com/ramirezfernando/cloc-tool/assets/91701930/4b188369-39d9-48b0-8fed-0d414b067e75">

## Features
- cloc-tool has a huge range of languages, supporting over **220** language extensions.
- cloc-tool is **accurate**, and **consistent** as it counts the number of newline characters `/n` present in a specified path. This ensures consistency across different platforms and text editors. Different text editors may interpret line endings differently (e.g., `\n` in Unix-like systems, `\r\n` in Windows), which could lead to discrepancies in line counts if you try to match the exact number of lines displayed in a specific editor.

## What's next?
- Add unit tests to double check the accuracy
- Making your command available to Homebrew users

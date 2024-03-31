# cloc-tool

A command line tool to count lines of code (CLOC), written in Go.

## Example
<img width="500" alt="Screen Shot 2024-03-28 at 1 16 24 AM" src="https://github.com/ramirezfernando/cloc-tool/assets/91701930/fce391c5-c470-4d4b-8d61-4a484d24f8dc">

## Features
- cloc-tool has huge range of languages, supporting over **220** languages extensions.
- cloc-tool is **accurate**, and **consistent** as it counts the number of newline characters (`/n`) that are present in a specified path. This ensures consitency across different platforms and text editors. Different text editors may interpret line endings differently (e.g., \n in Unix-like systems, \r\n in Windows), which could lead to discrepancies in line counts if you try to match the exact number of lines displayed in a specific editor.

## What's next?
- Add unit tests to double check the accuracy
- Making your command available to Homebrew users

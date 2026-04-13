package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {
	if len(os.Args) > 2 {
		fmt.Println("Please provide a filename")
		return
	}

	filename  := os.Args[1]
	output := "stat.txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s\n", err)
		return
	}
	content := string(data)
	charCount := len(content)

	words:= strings.Fields(content)
	wordCount := len(words)

	lines := strings.Split(content, "\n")
	lenLines := len(lines)

	summary := fmt.Sprintf("File: %s\nCharacters: %d\nWords: %d\nLines: %d\n",
							filename, charCount, wordCount, lenLines,
	)

	err = os.WriteFile(output, []byte(summary), 0644)
	if err != nil {
		fmt.Printf("error reading file %s\n", err)
		return
	}
	fmt.Println("Summary successfully written to stats.txt")
}
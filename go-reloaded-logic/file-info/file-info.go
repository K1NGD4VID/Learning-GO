package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	if len(os.Args) > 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}
	file := os.Args[1]
	output := "stats.txt"
	
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %s\n", err)
		return
	}

	content := string(data)

	charCount := len(content)

	words := strings.Fields(content)
	wordCount := len(words)

	lines := strings.Split(content, "\n")
	lineCount := len(lines)

	summary := fmt.Sprintf(
		"File : %s\nCharcaters: %d\nWords, %d\nlines, %d\n",
		file, charCount, wordCount, lineCount,
	)

	err = os.WriteFile(output, []byte(summary), 0644)
	fmt.Println("Stats saved to stats.txt")
}
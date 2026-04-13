package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("usage : go run . <filename>.txt")
		return
	}
	filename := os.Args[1]
	output := "upper_output.txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file%s", err)
		return
	}

	upperContent := strings.ToUpper(string(data))

	err = os.WriteFile(output, []byte(upperContent), 0644)
	if err != nil {
		fmt.Printf("Error reading file %s\n", err)
		return
	}
	
}
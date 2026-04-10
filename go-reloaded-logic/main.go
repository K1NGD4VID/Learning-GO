package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {

	if len(os.Args) > 2 {
		fmt.Printf("Usage: go run . <filename>.txt")
		return
	}

	input := os.Args[1]
	output := "upper_content.txt"

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("Error reading file %s\n", err)
		return
	}
	upper_content := strings.ToUpper(string(data))

	err = os.WriteFile(output, []byte(upper_content), 0644)
	if err != nil {
		fmt.Printf("Error reading file%s\n", err)
		return
	}
	fmt.Printf("Sucess! Upper case content saved to %s\n", output)
}
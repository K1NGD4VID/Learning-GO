package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage := go run . <filename>.txt")
		return
	}
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s\n", err)
		return
	}
	s := string(data)
	revTag := regexp.MustCompile(`\((reverse|double|shout)\)`)
	revTagmatch := revTag.FindStringSubmatchIndex(s)
	fmt.Println(revTagmatch)

}

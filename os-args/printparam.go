package main

import (
	"fmt"
	"os"
	"strings"
)

func printparam(s []string) string {
	words := os.Args[1:]
	return strings.Join(words, "\n")
}

func main() {
	fmt.Println(printparam(os.Args))
}

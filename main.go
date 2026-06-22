package main

import (
	"fmt"
	"strings"
)

func joinWords(words []string) string {

	return strings.Join(words, " ")
}
func main() {
	words := []string{"samuel", "is", "a", "bitch"}
	fmt.Println(joinWords(words))
}

package main

import (
	"strings"
)

func main() {
	// Replace the word just before "(delete)" with an empty string
	input := "keep this word (delete) and keep this"
	// Expected output: "keep this  and keep this"
	// Hint: use strings.Index to find the tag, then strings.Split to get words

	// tag := regexp.MustCompile(`\(.+?\)`)
	// tagMatch := tag.FindAllString(input, -1)

	pos := strings.Index(input, "(delete)")

}

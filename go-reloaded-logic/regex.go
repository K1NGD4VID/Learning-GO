package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "hello (up) world (low, 3) test (hex) done (cap, 2)"
	// Should find: ["(up)", "(low, 3)", "(hex)", "(cap, 2)"]

	re := regexp.MustCompile(`\(.+?\)`)
	matches := re.FindAllString(input, -1)

	fmt.Println(matches)
}

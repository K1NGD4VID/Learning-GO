package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// input := "keep this word (delete) and (delete) keep this hgbjibgsjnbgfnsnreuifnuivehfngvuihernfvuibneafvibdvfhbjc (delete) and this"

	// tag := regexp.MustCompile(`\w+\s*\(.+?\)`)

	// matches := tag.ReplaceAllString(input, "")

	// final := strings.Join(strings.Fields(matches), " ")

	// 	// fmt.Println(final)
	// 	words := []string{"one", "two", "three", "four", "five"}
	// 	// Print the last 3 words using len() and slicing
	// 	// Print the first 2 words using slicing
	// 	// Print words from index 1 to 3
	// 	final := words[1:3]
	// final := words[len(words) - 3:]

	// 	fmt.Println(final)

	Input := "my name is david (up, 2),"

	tag := regexp.MustCompile(`\(.+?\)`)
	tagMatches := tag.FindAllString(Input, -1)

	if len(tagMatches) > 0 {
		re := regexp.MustCompile(`\d+`)
		numStr := re.FindAllString(tagMatches[0], -1)

		pos := strings.Index(Input, tagMatches[0])
		beforeTag := Input[:pos]
		words := strings.Fields(beforeTag)
		n, _ := strconv.Atoi(numStr[0])

		start := len(words) - n
		if start < 0 {
			start = 0
		}

		for i := start; i < len(words); i++ {
			words[i] = strings.ToUpper(words[i])
		}
		afterTag := Input[pos+len(tagMatches[0]):]
		result := strings.Join(words, " ") + afterTag

		fmt.Println(result)

	}

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		input := "I am a boy"
		subStr := input[0:4]
		fmt.Println(subStr)
	*/

	/*
		str := "Hello, World!"
		for i, r := range str {
			fmt.Printf("index %d, Rune: %c\n", i, r)
		}
	*/

	/*

		input := "My name is Fuhad 1234"

		fmt.Println(strings.Contains(input, "is"))
		fmt.Println(strings.Count(input, "a"))
		fmt.Println(strings.Index(input, "a"))
		fmt.Println(strings.LastIndex(input, "a"))
		fmt.Println(strings.ToUpper(input))
		fmt.Println(strings.ToLower(input))
		fmt.Println(strings.Replace(input, "a", "x", 2))
		fmt.Println(strings.Split(input, ""))
	*/

	first := "abc"
	sec := "xyz"

	fmt.Println(strings.Compare(first, sec))

}

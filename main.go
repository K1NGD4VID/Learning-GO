package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "SHOUTING ALL THE TIME (down, 3)"
	words := strings.Split(s, " ")
	for _, word := range words {
		fmt.Println(word)
		if strings.HasPrefix(word, "("){
			fmt.Println("found a tag")
		}
	}

}
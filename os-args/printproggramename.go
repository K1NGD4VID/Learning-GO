package main

import (
	"fmt"
	"os"
)

func printprogrammename(s []string) string {
	if len(s) > 0 {
		return s[0] // Safely get the program name
	}
	return "no name found"
}

// 	name := os.Args[0]

//		return name
//	}
func main() {
	fmt.Println(printprogrammename(os.Args))

}

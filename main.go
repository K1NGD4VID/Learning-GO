package main

import "fmt"

func main() {
	items := 3
	pricePeritem := 50

	if total := items * pricePeritem; total >= 100 {
		fmt.Println("Eligible for shipping", total)
	} else {
		fmt.Println("Not Eligible for shipping", total)
	}
}

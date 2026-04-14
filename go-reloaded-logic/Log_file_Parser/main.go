package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	input := "app.log"
	output := "report.txt"

	content, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("Error reading file %s\n")
		return
	}
	s := string(content)

	Error := regexp.MustCompile(`\[ERROR\]`)
	Dates := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	Ports := regexp.MustCompile(`port \d+`)

	Errmatches := Error.FindAllString(s, -1)
	DatesMatches := Dates.FindAllString(s, -1)
	PortMatches := Ports.FindAllString(s, -1)

	summary := fmt.Sprintf("Error: %s\nDates: %s\nPorts: %s\n",
		Errmatches, DatesMatches, PortMatches,
	)

	err = os.WriteFile(output, []byte(summary), 0644)
	if err != nil {
		fmt.Printf("Error writing file %s\n", err)
		return
	}

	fmt.Println("Succecefuly saves to report.txt")

}

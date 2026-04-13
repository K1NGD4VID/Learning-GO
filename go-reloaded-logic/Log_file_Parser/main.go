package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	output := "report.txt"
	content, err := os.ReadFile("app.log")
	if err != nil {
		fmt.Printf("Error reading file %s\n", err)
		return
	}
	s := string(content)
	Error := regexp.MustCompile(`\[ERROR\]`)
	Dates := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	ports := regexp.MustCompile(`port \d+`)

	Errmatches := Error.FindAllString(s, -1)
	DatesMatches := Dates.FindAllString(s, -1)
	PortsMatches := ports.FindAllString(s, -1)

	summary := fmt.Sprintf("Error: %s\nDates: %s\nPorts: %s\n",
		Errmatches, DatesMatches, PortsMatches,
	)
	err = os.WriteFile(output, []byte(summary), 0644)
	if err != nil {
		fmt.Printf("Error reading file %s\n", err)
		return
	}
	fmt.Println("Summary successfully written to reports.txt")
}

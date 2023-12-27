package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

func main() {
	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	lines := splitLines(string(content))
	for _, line := range lines {
		firstDigit := extractDigit(line, true)
		lastDigit := extractDigit(line, false)

		calibrationValue := firstDigit*10 + lastDigit

		sum += calibrationValue
	}

	fmt.Printf("The sum of calibration values is: %d\n", sum)
}

func splitLines(text string) []string {
	var lines []string

	for _, line := range text {
		if line == '\n' {
			lines = append(lines, "")
		} else {
			if len(lines) == 0 {
				lines = append(lines, "")
			}
			lines[len(lines)-1] += string(line)
		}
	}

	return lines
}

func extractDigit(s string, isFirst bool) int {
	for _, c := range s {
		if unicode.IsDigit(c) {
			if isFirst {
				return int(c - '0')
			}
			lastDigit := int(c - '0')
			for _, c := range s {
				if unicode.IsDigit(c) {
					lastDigit = int(c - '0')
				}
			}
			return lastDigit
		}
	}

	return 0
}

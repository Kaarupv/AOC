package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	CubeCount()
}

func CubeCount() int {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	var bufferCount int

	for i := 0; scanner.Scan(); i++ {
		lines := scanner.Text()
		for j := 0; j < len(lines); j++ {
			if lines[j] >= '0' && lines[j] <= '9' {
				bufferCount *= 10
				bufferCount += int(lines[j] - '0')
				continue
			}
			if lines[j] == ' ' {
				continue
			}
			switch lines[j] {
			case 'r':
				if bufferCount > 12 {
					goto exit
				}
			case 'g':
				if bufferCount > 13 {
					goto exit
				}
			case 'b':
				if bufferCount > 14 {
					goto exit
				}
			}

			bufferCount = 0
		}
		total += i + 1
	exit:
	}

	fmt.Println(total)
	return total
}

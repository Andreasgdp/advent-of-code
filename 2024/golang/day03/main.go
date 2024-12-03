package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	// file, err := os.Open("sample_input1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Regular expression to match valid mul(X,Y) instructions
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) == 3 {
				x, err1 := strconv.Atoi(match[1])
				y, err2 := strconv.Atoi(match[2])
				if err1 == nil && err2 == nil {
					total += x * y
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total:", total)
}

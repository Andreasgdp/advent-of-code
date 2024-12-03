package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func findDoAndDontAndMulInstructions(text string) []string {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	matches := re.FindAllString(text, -1)
	return matches
}

func main() {
	// Open the input file
	// file, err := os.Open("sample_input2.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var instructions []string

	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, findDoAndDontAndMulInstructions(line)...)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Process the collected instructions
	total := 0
	mulEnabled := true // mul instructions are enabled by default

	for _, instruction := range instructions {
		if instruction == "do()" {
			mulEnabled = true
		} else if instruction == "don't()" {
			mulEnabled = false
		} else if mulEnabled {
			mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
			matches := mulRe.FindStringSubmatch(instruction)
			if len(matches) == 3 {
				x, err1 := strconv.Atoi(matches[1])
				y, err2 := strconv.Atoi(matches[2])
				if err1 == nil && err2 == nil {
					total += x * y
				}
			}
		}
	}

	fmt.Println("Total:", total)
}

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

func readLines(file *os.File, lines chan<- string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
}

func processInstructions(lines <-chan string, instructions chan<- string) {
	for line := range lines {
		for _, instruction := range findDoAndDontAndMulInstructions(line) {
			instructions <- instruction
		}
	}
	close(instructions)
}

func calculateTotal(instructions <-chan string, results chan<- int) {
	total := 0
	mulEnabled := true

	for instruction := range instructions {
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
	results <- total
	close(results)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lines := make(chan string)
	instructions := make(chan string)
	results := make(chan int)

	go readLines(file, lines)
	go processInstructions(lines, instructions)
	go calculateTotal(instructions, results)

	total := <-results
	fmt.Println("Total:", total)
}

package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func findDoAndDontAndMulInstructions(text string) []string {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	matches := re.FindAllString(text, -1)
	return matches
}

func readLines(input io.Reader, lines chan<- string) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
}

func processInstructions(lines <-chan string, instructions chan<- string, includeDoDont bool) {
	for line := range lines {
		for _, instruction := range findDoAndDontAndMulInstructions(line) {
			if includeDoDont || regexp.MustCompile(`mul\(\d+,\d+\)`).MatchString(instruction) {
				instructions <- instruction
			}
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

func fs1(input io.Reader) int {
	lines := make(chan string)
	instructions := make(chan string)
	results := make(chan int)

	go readLines(input, lines)
	go processInstructions(lines, instructions, false)
	go calculateTotal(instructions, results)

	total := <-results
	return total
}

func fs2(input io.Reader) int {
	lines := make(chan string)
	instructions := make(chan string)
	results := make(chan int)

	go readLines(input, lines)
	go processInstructions(lines, instructions, true)
	go calculateTotal(instructions, results)

	total := <-results
	return total
}

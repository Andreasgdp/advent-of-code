package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Function to calculate the total distance between two lists
func totalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	totalDistance := 0
	for i := range left {
		totalDistance += abs(left[i] - right[i])
	}
	return totalDistance
}

// Function to calculate the similarity score between two lists
func similarityScore(left, right []int) int {
	rightCount := make(map[int]int)
	for _, num := range right {
		rightCount[num]++
	}
	similarityScore := 0
	for _, num := range left {
		similarityScore += num * rightCount[num]
	}
	return similarityScore
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	// Part 1
	fmt.Println("Total Distance:", totalDistance(left, right))

	// Part 2
	fmt.Println("Similarity Score:", similarityScore(left, right))

	// Read input line by line from input.txt
	// File input.txt contains a lot of lines lines with space separated integers
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	leftInput := make([]int, 0)
	rightInput := make([]int, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Fields(line)
		leftInput = append(leftInput, parseInt(parts[0]))
		rightInput = append(rightInput, parseInt(parts[1]))
	}

	// Part 1 for input.txt
	fmt.Println("Total Distance (input.txt):", totalDistance(leftInput, rightInput))

	// Part 2 for input.txt
	fmt.Println("Similarity Score (input.txt):", similarityScore(leftInput, rightInput))
}

func parseInt(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("Error converting %s to int: %v", line, err)
	}
	return num
}

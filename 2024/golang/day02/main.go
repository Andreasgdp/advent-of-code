package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Function to check if a report is safe without removing any level
func isSafe(report []int) bool {
	increasing := false
	decreasing := false

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if diff > 0 {
			increasing = true
		}
		if diff < 0 {
			decreasing = true
		}
	}

	if increasing && decreasing {
		return false
	}

	return increasing || decreasing
}

// if we remove any single level (element) from the report and it is still safe,
// it is considered safe with dampener
func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		// remove the ith element from the report
		var newReport []int
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if isSafe(newReport) {
			return true
		}
	}

	return false
}

// Function to read reports from a file
func readReports(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func main() {
	filenames := []string{"sample_input1.txt", "sample_input2.txt", "input.txt"}
	// filenames := []string{"sample_input1.txt", "sample_input2.txt"}

	for _, filename := range filenames {
		reports, err := readReports(filename)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filename, err)
			continue
		}

		safeCount := 0
		safeWithDampenerCount := 0

		for _, report := range reports {
			if isSafe(report) {
				safeCount++
				safeWithDampenerCount++
			} else if isSafeWithDampener(report) {
				safeWithDampenerCount++
			}
		}

		fmt.Printf("File: %s\n", filename)
		fmt.Printf("Safe reports: %d\n", safeCount)
		fmt.Printf("Safe reports with dampener: %d\n", safeWithDampenerCount)
	}
}

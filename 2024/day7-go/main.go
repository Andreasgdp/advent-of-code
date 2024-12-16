package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Equation struct {
	testVal int
	numbers []int
}

// example input
// 190: 10 19
// 3267: 81 40 27
// 83: 17 5
// 156: 15 6
// 7290: 6 8 6 15
// 161011: 16 10 13
// 192: 17 8 14
// 21037: 9 7 18 13
// 292: 11 6 16 20

// takes reader, returns slice of tuple of ( int, slice of int )
func readInput(input io.Reader) []Equation {
	scanner := bufio.NewScanner(input)
	eqList := []Equation{}

	for scanner.Scan() {
		line := scanner.Text()
		eq := Equation{}
		testValStr := line[:strings.Index(line, ":")]
		eq.testVal, _ = strconv.Atoi(testValStr)

		numberStrs := strings.Fields(line[strings.Index(line, ":")+1:])
		eq.numbers = make([]int, len(numberStrs))
		for i, numStr := range numberStrs {
			eq.numbers[i], _ = strconv.Atoi(numStr)
		}
		eqList = append(eqList, eq)
	}

	return eqList
}

func getCombinationsResults(numbers []int, operators []string) []int {
	if len(numbers) == 1 {
		return numbers
	}

	var results []int
	var generateCombinations func(nums []int, currentResult int, index int)

	generateCombinations = func(nums []int, currentResult int, index int) {
		if index == len(nums) {
			results = append(results, currentResult)
			return
		}

		for _, op := range operators {
			if op == "+" {
				generateCombinations(nums, currentResult+nums[index], index+1)
			} else if op == "*" {
				generateCombinations(nums, currentResult*nums[index], index+1)
			} else if op == "||" {
				concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentResult, nums[index]))
				generateCombinations(nums, concatenated, index+1)
			}
		}
	}

	for i := 1; i < len(numbers); i++ {
		for _, op := range operators {
			if op == "+" {
				generateCombinations(numbers, numbers[0]+numbers[i], i+1)
			} else if op == "*" {
				generateCombinations(numbers, numbers[0]*numbers[i], i+1)
			} else if op == "||" {
				concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", numbers[0], numbers[i]))
				generateCombinations(numbers, concatenated, i+1)
			}
		}
	}

	return results
}

// takes testVal and numbers, returns bool if testVal can be made using multiply + and times *
// numbers in list must preserve order
// ex: 190: 10 19 -> true as 10*19 = 190, while 10+19 != 190
// ex: 3267: 81 40 27 -> true as 81+40*27 = 3267, but all other combinations are false
func canEqTestValUsingPlusAndTimes(testVal int, numbers []int, operators []string) bool {
	if len(numbers) == 1 {
		return numbers[0] == testVal
	}

	// in between each number, we can either add or multiply, if one of the combinations is true, return true
	combinations := getCombinationsResults(numbers, operators)
	fmt.Println("testVal", testVal, "combinations", combinations)

	for _, result := range combinations {
		if result == testVal {
			fmt.Println("result", result, "testVal", testVal)
			return true
		}
	}

	return false
}

func fs1(input io.Reader) int {
	eqList := readInput(input)
	total := 0
	operators := []string{"+", "*"}
	for _, eq := range eqList {
		if canEqTestValUsingPlusAndTimes(eq.testVal, eq.numbers, operators) {
			total += eq.testVal
		}
	}
	return total
}

func fs2(input io.Reader) int {
	eqList := readInput(input)
	total := 0
	operators := []string{"+", "*", "||"}
	for _, eq := range eqList {
		if canEqTestValUsingPlusAndTimes(eq.testVal, eq.numbers, operators) {
			total += eq.testVal
		}
	}
	return total
}

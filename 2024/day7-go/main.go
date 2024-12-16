package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Equation struct {
	numbers []int
	testVal int
}

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
	var applyOperator func(a, b int, op string) int
	var generateCombinations func(nums []int, currentResult int, index int)

	applyOperator = func(before, after int, operator string) int {
		switch operator {
		case "+":
			return before + after
		case "*":
			return before * after
		case "||":
			concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", before, after))
			return concatenated
		default:
			return 0
		}
	}

	generateCombinations = func(nums []int, currentResult int, index int) {
		if index == len(nums) {
			results = append(results, currentResult)
			return
		}

		for _, op := range operators {
			newResult := applyOperator(currentResult, nums[index], op)
			generateCombinations(nums, newResult, index+1)
		}
	}

	for i := 1; i < len(numbers); i++ {
		for _, op := range operators {
			newResult := applyOperator(numbers[0], numbers[i], op)
			generateCombinations(numbers, newResult, i+1)
		}
	}

	return results
}

// takes testVal and numbers, returns bool if testVal can be made using multiply + and times *
// numbers in list must preserve order
// ex: 190: 10 19 -> true as 10*19 = 190, while 10+19 != 190
// ex: 3267: 81 40 27 -> true as 81+40*27 = 3267, but all other combinations are false
func canEqTestVal(testVal int, numbers []int, operators []string) bool {
	if len(numbers) == 1 {
		return numbers[0] == testVal
	}

	// in between each number, we can either add or multiply, if one of the combinations is true, return true
	combinations := getCombinationsResults(numbers, operators)

	for _, result := range combinations {
		if result == testVal {
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
		if canEqTestVal(eq.testVal, eq.numbers, operators) {
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
		if canEqTestVal(eq.testVal, eq.numbers, operators) {
			total += eq.testVal
		}
	}
	return total
}

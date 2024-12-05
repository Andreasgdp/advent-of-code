package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func parseInput(input io.Reader) ([]string, [][]int) {
	scanner := bufio.NewScanner(input)
	var rules []string
	var updates [][]int
	isUpdateSection := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isUpdateSection = true
			continue
		}
		if isUpdateSection {
			updateNumbers := strings.Split(line, ",")
			var update []int
			for _, numStr := range updateNumbers {
				num, _ := strconv.Atoi(numStr)
				update = append(update, num)
			}
			updates = append(updates, update)
		} else {
			rules = append(rules, line)
		}
	}
	return rules, updates
}

// Update is valid if all rules are satisfied
func isUpdateValid(rules []string, update []int) bool {
	for _, rule := range rules {
		ruleParts := strings.Split(rule, "|")
		before, _ := strconv.Atoi(ruleParts[0])
		after, _ := strconv.Atoi(ruleParts[1])

		if !updateSatisfiesRule(update, before, after) {
			return false
		}
	}

	return true
}

func sumMiddleValueOfValidUpdates(validUpdateIndexes []int, updates [][]int) int {
	// Sum the middle value of all valid updates
	// The middle value is the value at the index of the middle of the update
	// If the update has an even number of values, the middle value is the value at the index of the middle - 1
	sum := 0
	for _, updateIndex := range validUpdateIndexes {
		update := updates[updateIndex]
		middleIndex := len(update) / 2
		if len(update)%2 == 0 {
			middleIndex--
		}
		sum += update[middleIndex]
	}

	return sum
}

// A rule defines an order of what pages (number) are valid e.g. rule 77|52 means that 66 should be before 33
// An update is e.g. 83,84,11,26,77,34,14,85,71,52,18
// In this case the update is valid because 77 is before 52, it is not valid if there were another rule of e.g. 84|83
func updateSatisfiesRule(update []int, before int, after int) bool {
	index1 := -1
	index2 := -1
	for i, num := range update {
		if num == before {
			index1 = i
		}
		if num == after {
			index2 = i
		}
	}

	if index1 == -1 || index2 == -1 {
		return true
	}

	return index1 < index2
}

func fs1(input io.Reader) int {
	rules, updates := parseInput(input)
	validUpdateIndexes := make([]int, 0)
	for i, update := range updates {
		if isUpdateValid(rules, update) {
			validUpdateIndexes = append(validUpdateIndexes, i)
		}
	}

	return sumMiddleValueOfValidUpdates(validUpdateIndexes, updates)
}

func reorderUpdate(rules []string, update []int) []int {
	for {
		allRulesSatisfied := true
		for _, rule := range rules {
			ruleParts := strings.Split(rule, "|")
			before, _ := strconv.Atoi(ruleParts[0])
			after, _ := strconv.Atoi(ruleParts[1])
			if !updateSatisfiesRule(update, before, after) {
				allRulesSatisfied = false
				index1 := -1
				index2 := -1
				for i, num := range update {
					if num == before {
						index1 = i
					}
					if num == after {
						index2 = i
					}
				}
				update[index1], update[index2] = update[index2], update[index1]
			}
		}
		if allRulesSatisfied {
			break
		}
	}
	return update
}

func fs2(input io.Reader) int {
	rules, updates := parseInput(input)
	validReorderedUpdateIndexes := make([]int, 0)
	for i, update := range updates {
		if !isUpdateValid(rules, update) {
			reorderedUpdate := reorderUpdate(rules, update)
			if isUpdateValid(rules, reorderedUpdate) {
				validReorderedUpdateIndexes = append(validReorderedUpdateIndexes, i)
			}

		}
	}

	return sumMiddleValueOfValidUpdates(validReorderedUpdateIndexes, updates)
}

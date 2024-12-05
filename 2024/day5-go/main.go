package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func parseInput(input io.Reader) ([]string, [][]string) {
	scanner := bufio.NewScanner(input)
	var rules []string
	var updates [][]string
	isUpdateSection := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isUpdateSection = true
			continue
		}
		if isUpdateSection {
			updates = append(updates, parseUpdate(line))
		} else {
			rules = append(rules, line)
		}
	}
	return rules, updates
}

func parseUpdate(line string) []string {
	return strings.Split(line, ",")
}

func isUpdateValid(rules []string, update []string) bool {
	for _, rule := range rules {
		if !updateSatisfiesRule(update, rule) {
			return false
		}
	}
	return true
}

func sumMiddleValueOfValidUpdates(validUpdateIndexes []int, updates [][]string) int {
	sum := 0
	for _, idx := range validUpdateIndexes {
		update := updates[idx]
		middleIndex := len(update) / 2
		if len(update)%2 == 0 {
			middleIndex--
		}
		num, _ := strconv.Atoi(update[middleIndex])
		sum += num
	}
	return sum
}

func updateSatisfiesRule(update []string, rule string) bool {
	parts := strings.Split(rule, "|")
	before := parts[0]
	after := parts[1]

	index1, index2 := -1, -1
	for i, num := range update {
		if num == before {
			index1 = i
		}
		if num == after {
			index2 = i
		}
	}

	return index1 == -1 || index2 == -1 || index1 < index2
}

func fs1(input io.Reader) int {
	rules, updates := parseInput(input)
	validUpdateIndexes := []int{}
	for i, update := range updates {
		if isUpdateValid(rules, update) {
			validUpdateIndexes = append(validUpdateIndexes, i)
		}
	}
	return sumMiddleValueOfValidUpdates(validUpdateIndexes, updates)
}

func reorderUpdate(rules []string, update []string) []string {
	for {
		allRulesSatisfied := true
		for _, rule := range rules {
			if !updateSatisfiesRule(update, rule) {
				allRulesSatisfied = false
				swapUpdateElements(update, rule)
			}
		}
		if allRulesSatisfied {
			break
		}
	}
	return update
}

func swapUpdateElements(update []string, rule string) {
	parts := strings.Split(rule, "|")
	before := parts[0]
	after := parts[1]

	index1, index2 := -1, -1
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

func fs2(input io.Reader) int {
	rules, updates := parseInput(input)
	validReorderedUpdateIndexes := []int{}
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

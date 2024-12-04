package main

import (
	"io"

	aoc "github.com/teivah/advent-of-code"
)

func fs1(input io.Reader) int {
	grid := aoc.ReaderToStrings(input)
	word := "XMAS"
	return countWordOccurrences(grid, word)
}

func countWordOccurrences(grid []string, word string) int {
	directions := []struct {
		dx, dy int
	}{
		{1, 0},   // horizontal right
		{-1, 0},  // horizontal left
		{0, 1},   // vertical down
		{0, -1},  // vertical up
		{1, 1},   // diagonal down-right
		{-1, -1}, // diagonal up-left
		{1, -1},  // diagonal down-left
		{-1, 1},  // diagonal up-right
	}

	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			for _, dir := range directions {
				if checkWord(grid, word, x, y, dir.dx, dir.dy) {
					count++
				}
			}
		}
	}
	return count
}

func checkWord(grid []string, word string, x, y, dx, dy int) bool {
	for i := 0; i < len(word); i++ {
		nx, ny := x+dx*i, y+dy*i
		if nx < 0 || ny < 0 || ny >= len(grid) || nx >= len(grid[ny]) || grid[ny][nx] != word[i] {
			return false
		}
	}
	return true
}

func fs2(input io.Reader) int {
	grid := aoc.ReaderToStrings(input)
	return countXMASOccurrences(grid)
}

func countXMASOccurrences(grid []string) int {
	count := 0
	for y := 0; y < len(grid)-2; y++ {
		for x := 0; x < len(grid[y])-2; x++ {
			if checkXMAS(grid, x, y) {
				count++
			}
		}
	}
	return count
}

func checkXMAS(grid []string, x, y int) bool {
	patterns := [][]string{
		{"M.S", ".A.", "M.S"},
		{"S.S", ".A.", "M.M"},
		{"S.M", ".A.", "S.M"},
		{"M.M", ".A.", "S.S"},
	}

	for _, pattern := range patterns {
		if matchPattern(grid, x, y, pattern) {
			return true
		}
	}
	return false
}

func matchPattern(grid []string, x, y int, pattern []string) bool {
	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 3; dx++ {
			if pattern[dy][dx] != '.' && grid[y+dy][x+dx] != pattern[dy][dx] {
				return false
			}
		}
	}
	return true
}

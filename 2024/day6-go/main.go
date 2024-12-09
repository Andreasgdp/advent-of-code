package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

type Cell struct {
	empty            bool
	obstruction      bool
	isTrail          bool
	extraObstruction bool
}

func getCellGrid(input io.Reader) ([][]Cell, [2]int) {
	scanner := bufio.NewScanner(input)
	cellGrid := make([][]Cell, 0)
	startingPos := [2]int{-1, -1}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Cell, len(line))
		for x, char := range line {
			switch char {
			case '.':
				row[x] = Cell{empty: true}
			case '#':
				row[x] = Cell{obstruction: true}
			case '^':
				row[x] = Cell{empty: true}
				startingPos = [2]int{x, y}
			default:
				panic("Invalid character")
			}
		}
		cellGrid = append(cellGrid, row)
		y++
	}

	if startingPos[0] == -1 || startingPos[1] == -1 {
		panic("No starting position found")
	}

	return cellGrid, startingPos
}

// Goes through the guard trail, from starting pos to out of bounds.
// Guard starts moving up and then turns right on every obstruction.
// Returns the number of empty cells visited.
func getGuardTrailCount(cellGrid [][]Cell, startingPos [2]int) int {
	x, y := startingPos[0], startingPos[1]
	direction := [2]int{0, -1}

	for {
		// Move the guard
		x += direction[0]
		y += direction[1]

		// Check if the guard is out of bounds
		if y < 0 || y >= len(cellGrid) || x < 0 || x >= len(cellGrid[0]) {
			fmt.Println("out of bounds")
			fmt.Println(x, y)
			break
		}

		// Check if the guard is on an empty cell
		if cellGrid[y][x].empty {
			cellGrid[y][x].isTrail = true
		}

		// Check if the guard needs to turn right
		if cellGrid[y][x].obstruction {
			x -= direction[0]
			y -= direction[1]
			if direction[1] == -1 {
				direction = [2]int{1, 0}
			} else if direction[0] == 1 {
				direction = [2]int{0, 1}
			} else if direction[1] == 1 {
				direction = [2]int{-1, 0}
			} else if direction[0] == -1 {
				direction = [2]int{0, -1}
			}
		}
	}

	// drawCellGrid(cellGrid, startingPos)

	// Count the number cell.isTrail
	count := 0
	for _, row := range cellGrid {
		for _, cell := range row {
			if cell.isTrail {
				count++
			}
		}
	}

	// gotta cound the starting position as well
	return count + 1
}

func drawCellGrid(cellGrid [][]Cell, startingPos [2]int) {
	// clear the screen
	fmt.Print("\033[H\033[2J")
	for y, row := range cellGrid {
		for x, cell := range row {
			if x == startingPos[0] && y == startingPos[1] {
				fmt.Print("^")
			} else {
				if cell.obstruction {
					fmt.Print("#")
				} else if cell.extraObstruction {
					fmt.Print("O")
				} else if cell.isTrail {
					fmt.Print("X")
				} else if cell.empty {
					fmt.Print(".")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
	// wait a little to simulate animation
	time.Sleep(100 * time.Millisecond)
}

func fs1(input io.Reader) int {
	cellGrid, startingPos := getCellGrid(input)

	return getGuardTrailCount(cellGrid, startingPos)
}

// needs to count the number of different placements for a single obstruction to be placed,
// that would cause the guard to get stuck in a loop.
func countObstructionPlacementsForGuardToGetStuckInALoop(cellGrid [][]Cell, startingPos [2]int) int {
	count := 0

	for y, row := range cellGrid {
		for x, cell := range row {
			if cell.empty {
				// Place an obstruction
				cellGrid[y][x].obstruction = true
				cellGrid[y][x].empty = false

				// Check if the guard gets stuck in a loop
				if isGuardStuckInLoop(cellGrid, startingPos) {
					cellGrid[y][x].extraObstruction = true
					count++
				}

				// Remove the obstruction
				cellGrid[y][x].obstruction = false
				cellGrid[y][x].empty = true
			}
		}
	}

	return count
}

func isGuardStuckInLoop(cellGrid [][]Cell, startingPos [2]int) bool {
	x, y := startingPos[0], startingPos[1]
	direction := [2]int{0, -1}
	visited := make(map[[4]int]bool)

	for {
		// Move the guard
		x += direction[0]
		y += direction[1]

		// TODO: why are we getting all these possible placements, there should only be 6 in total for the test input
		drawCellGrid(cellGrid, startingPos)

		// Check if the guard is out of bounds
		if y < 0 || y >= len(cellGrid) || x < 0 || x >= len(cellGrid[0]) {
			return false
		}

		// Check if the guard is on an empty cell
		if cellGrid[y][x].empty {
			visited[[4]int{x, y, direction[0], direction[1]}] = true
		}

		// Check if the guard needs to turn right
		if cellGrid[y][x].obstruction {
			x -= direction[0]
			y -= direction[1]
			if direction[1] == -1 {
				direction = [2]int{1, 0}
			} else if direction[0] == 1 {
				direction = [2]int{0, 1}
			} else if direction[1] == 1 {
				direction = [2]int{-1, 0}
			} else if direction[0] == -1 {
				direction = [2]int{0, -1}
			}
		}

		// Check if the guard is stuck in a loop
		if visited[[4]int{x, y, direction[0], direction[1]}] {
			return true
		}
	}
}

func fs2(input io.Reader) int {
	cellGrid, startingPos := getCellGrid(input)
	return countObstructionPlacementsForGuardToGetStuckInALoop(cellGrid, startingPos)
}

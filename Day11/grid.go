package main

import (
	"fmt"
	"strings"
)

func (grid Grid) applyRules() (Grid, bool) {
	nextGrid := grid.copyGrid()
	hasChanged := false

	for i, row := range(grid) {
		for j, square := range(row) {
			switch rune(square) {
			case 'L':
				if grid.countAdjacentAndSelf(i, j, '#') == 0 {
					hasChanged = true
					nextGrid[i][j] = '#'
				}
			case '#':
				if grid.countAdjacentAndSelf(i, j, '#') >= 5 {
					hasChanged = true
					nextGrid[i][j] = 'L'
				}
			}
		}
	}

	return nextGrid, hasChanged
}

func (grid Grid) countAdjacentAndSelf(y int, x int, charToCount rune) int {
	minRow := max(y - 1, 0)
	maxRow := min(y + 1, len(grid) - 1)
	minColumn := max(x - 1, 0)
	maxColumn := min(x + 1, len(grid[0]) - 1)
	count := 0

	for _, row := range(grid[minRow:maxRow + 1]) {
		for _, char := range(row[minColumn:maxColumn + 1]) {
			if rune(char) == charToCount {
				count++
			}
		}
	}

	return count
}

func (template Grid) copyGrid() Grid {
	nextGrid := make(Grid, len(template))
	for i, _ := range(template) {
		nextGrid[i] = make([]byte, len(template))
		copy(nextGrid[i], template[i])
	}

	return nextGrid
}

func (grid Grid) print() {
	for i, row := range(grid) {
		fmt.Println(i, string(row))
	}
}

func (grid Grid) countOccupied() int {
	count := 0
	for _, row := range(grid) {
		count += strings.Count(string(row), "#")
	}

	return count
}
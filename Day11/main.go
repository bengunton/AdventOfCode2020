package main

import(
	"os"
	"bufio"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Grid [][]byte

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make(Grid, 0)
	
	for i := 0; scanner.Scan(); i++ {
		bytes := scanner.Bytes()
		grid = append(grid, make([]byte, len(bytes)))
		copy(grid[i], bytes)
	}
	check(scanner.Err())

	isUnstable := true
	for isUnstable {
		grid, isUnstable = grid.applyRules()
	}

	grid, _ = grid.applyRules()

	fmt.Printf("There are %d occupied seats\n", grid.countOccupied())
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
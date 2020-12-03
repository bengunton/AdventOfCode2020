package main 

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

type trajectory struct {
	xOffset int
	yOffset int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputMap := make([]string, 0)


	for scanner.Scan() {
		inputMap = append(inputMap, scanner.Text())
	}

	trees3_1 := countTrees(inputMap, trajectory{3, 1})

	fmt.Printf("There are %d trees on the way! (for trajectory: (3,1))\n", trees3_1)

	trees1_1 := countTrees(inputMap, trajectory{1, 1})
	trees5_1 := countTrees(inputMap, trajectory{5, 1})
	trees7_1 := countTrees(inputMap, trajectory{7, 1})
	trees1_2 := countTrees(inputMap, trajectory{1, 2})

	fmt.Printf("Number of trees on all routes multiplied: %d!\n", trees1_1 * trees1_2 * trees3_1 * trees5_1 * trees7_1)
}

func countTrees(inputMap []string, t trajectory) int {
	trees := 0
	pos := position{0,0}
	mapWidth := len(inputMap[0])

	for (pos.y != (len(inputMap) - 1)) {
		if inputMap[pos.y][pos.x] == byte('#') {
			trees++
		}

		pos.move(t.xOffset, t.yOffset, mapWidth)
	}

	if inputMap[pos.y][pos.x] == byte('#') {
		trees++
	}

	return trees
}

func (p *position) move(xOffset int, yOffset int, maxX int) {
	p.x = (p.x + xOffset) % maxX
	p.y += yOffset
}
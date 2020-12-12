package main

import (
	"fmt"
	"os"
	"sort"
	"bufio"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	joltages := make([]int, 1)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		check(err)
		joltages = append(joltages, num)
	}

	sort.Ints(joltages)
	joltages = append(joltages, joltages[len(joltages) - 1] + 3)

	oneJoltJumps, threeJoltJumps := countJoltJumps(joltages)

	fmt.Printf("%d 1-jolt jumps and %d 3-jolt jumps\n", oneJoltJumps, threeJoltJumps)
	fmt.Printf("Product is %d ", oneJoltJumps * threeJoltJumps)
}

func countJoltJumps(joltages []int) (int, int) {
	oneJoltJumps, threeJoltJumps := 0, 0
	for i, n := range(joltages) {
		if i == (len(joltages) - 1) {
			break
		}
		difference := joltages[i+1] - n
		switch difference {
		case 1:
			oneJoltJumps++
		case 3:
			threeJoltJumps++
		}
	}

	return oneJoltJumps, threeJoltJumps
}
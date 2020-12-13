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
var memo map[int]int = make(map[int]int)

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// start with the wall (0)
	joltages := make([]int, 1)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		check(err)
		joltages = append(joltages, num)
	}

	sort.Ints(joltages)
	// append the device (max + 3)
	joltages = append(joltages, joltages[len(joltages) - 1] + 3)

	oneJoltJumps, threeJoltJumps := countJoltJumps(joltages)

	fmt.Printf("%d 1-jolt jumps and %d 3-jolt jumps\n", oneJoltJumps, threeJoltJumps)
	fmt.Printf("Product is %d\n", oneJoltJumps * threeJoltJumps)

	fmt.Printf("There are %d joltage combinations to charge the device\n", countCombinations(0, joltages[:(len(joltages) - 1)]))
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


func countCombinations(startJolt int, joltages []int) int {
	// save recalculating when we already know how many branches there are from this point
	if (joltages[0] - startJolt) > 3 {
		return 0
	}
	if (memo[joltages[0]] != 0) {
		return memo[joltages[0]]
	}
	if (len(joltages) == 1) {
		return 1
	}

	sum := 0
	switch {
	case len(joltages) > 3:
		sum += countCombinations(joltages[0], joltages[3:])
		fallthrough
	case len(joltages) > 2:
		sum += countCombinations(joltages[0], joltages[2:])
		fallthrough
	case len(joltages) > 1:
		sum += countCombinations(joltages[0], joltages[1:])
	}

	// save value once we know it the first time
	if memo[joltages[0]] == 0 {
		memo[joltages[0]] = sum
	}
	return sum
}


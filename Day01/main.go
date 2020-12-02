package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("input")
	check(err)

	expenses, err := parseInputs(string(file))
	check(err)

	const target = 2020
	x, y, err := findPairSummingToN(expenses, target)
	check(err)

	fmt.Printf("%d + %d = %d\n", x, y, target)
	fmt.Printf("%d * %d = %d\n", x, y, x*y)

	x, y, z, err := findTripleSummingToN(expenses, target)
	check(err)

	fmt.Printf("%d + %d + %d = %d\n", x, y, z, target)
	fmt.Printf("%d * %d * %d = %d\n", x, y, z, x*y*z)
}

func parseInputs(fileContents string) ([]int, error) {
	trimmedContents := strings.TrimSpace(fileContents)
	parsedStrings := strings.Split(trimmedContents, "\r\n")

	var parsedInts = make([]int, len(parsedStrings), len(parsedStrings))
	for i, s := range parsedStrings {
		parsedInt, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		parsedInts[i] = parsedInt
	}

	return parsedInts, nil
}

func findPairSummingToN(candidates []int, target int) (int, int, error) {
	for index, lhs := range candidates {
		for _, rhs := range candidates[index:] {
			if lhs + rhs == target {
				return lhs, rhs, nil
			}
		}
	}

	return -1, -1, fmt.Errorf("No match found for target: %d", target)
}

func findTripleSummingToN(candidates []int, target int) (int, int, int, error) {
	for xIndex, x := range candidates {
		for yIndex, y := range candidates[xIndex:] {
			for _, z := range candidates[yIndex:] {
				if x + y + z == target {
					return x, y, z, nil
				}
			}
		}
	}

	return -1, -1, -1, fmt.Errorf("No match found for target: %d", target)
}
package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const BUFFER_SIZE = 25

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buffer := make([]int, BUFFER_SIZE, BUFFER_SIZE)
	fullList := make([]int, 0)
	var target int

	for i := 0; scanner.Scan(); i++ {
		num, err := strconv.Atoi(scanner.Text())
		fullList = append(fullList, num)
		check(err)
		
		if (i >= BUFFER_SIZE) {
			if !isValidEntry(num, buffer) {
				target = num
				fmt.Printf("Invalid entry: %d\n", num)
			}
		}

		buffer[i % BUFFER_SIZE] = num

	}

	for i, _ := range(fullList) {
		foundSum, ints := checkForContinuousSumToN(fullList[i:], target)
		if foundSum {
			fmt.Println("Contiguous ints are", ints)
			printMinAndMax(ints)
			break;
		}
	}
}

func isValidEntry(target int, buffer []int) bool {
	for i, x := range(buffer) {
		for _, y := range(buffer[i:]) {
			if x + y == target {
				return true
			}
		}
	}

	return false
}

func checkForContinuousSumToN(ints []int, target int) (bool, []int) {
	numsToSum := make([]int, 0)
	for _, num := range(ints) {
		target -= num
		numsToSum = append(numsToSum, num)

		if (target  < 0) {
			break;
		}
		if (target == 0 && len(numsToSum) >= 2) {
			return true, numsToSum
		}
	}

	return false, nil
}

func printMinAndMax(ints []int) {
	min, max := ints[0], ints[0]
	for _, n := range(ints) {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	fmt.Printf("Min is %d, max is %d, sum is %d\n", min, max, min + max)
}

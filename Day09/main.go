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

	for i := 0; scanner.Scan(); i++ {
		num, err := strconv.Atoi(scanner.Text())
		check(err)
		
		if (i >= BUFFER_SIZE) {
			reportIfInvalid(num, buffer)
		}

		buffer[i % BUFFER_SIZE] = num

	}
}

func reportIfInvalid(num int, buffer []int) {
	if !isValidEntry(num, buffer) {
		fmt.Printf("Invalid entry: %d\n", num)
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
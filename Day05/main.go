package main

import (
	"fmt"
	"os"
	"bufio"
)

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

	var max uint = 0

	for scanner.Scan() {
		boardingPass := scanner.Text()
		id := calculateId(boardingPass)
		if id > max {
			max = id
		}
	}

	fmt.Printf("Max id is %d\n", max)
}

func calculateId(boardingPass string) uint {
	sum := uint(0)
	for i := 0; i < 7; i++ {
		sum = sum << 1
		if boardingPass[i] == 'B' {
			sum++
		}
	}

	for i := 7; i < 10; i++ {
		sum = sum << 1
		if boardingPass[i] == 'R' {
			sum++
		}
	}
	return sum
}
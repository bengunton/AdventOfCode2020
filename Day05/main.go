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
	foundPasses := make(map[uint] bool)	

	for scanner.Scan() {
		boardingPass := scanner.Text()
		id := calculateId(boardingPass)

		if id > max {
			max = id
		}

		foundPasses[id] = true
	}

	for id, _ := range(foundPasses) {
		if (!foundPasses[id + 1] && foundPasses[id + 2]) {
			fmt.Println("Found empty space ", (id + 1))
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
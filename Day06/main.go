package main

import(
	"fmt"
	"os"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var GroupRegex = "([a-z]+\r\n)+\r\n"

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	groups := make([][]string, 0)
	groups = append(groups, make([]string, 0))
	groupNum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		response := scanner.Text()
		if response == "" {
			groupNum++;
			groups = append(groups, make([]string, 0))
		} else {
			groups[groupNum] = append(groups[groupNum], response)
		}
		
	}

	sum := 0
	for _, group := range(groups) {
		sum += sumGroup(group)
	}
	fmt.Printf("Sum is %d\n", sum)
}

func sumGroup(group []string) int {
	answered := make([]int, 26)
	for _, entry := range(group) {
		for _, char := range(entry) {
			index := int(char) - int('a')
			answered[index]++
		}
	}

	return countWhereEqual(answered, len(group))
}

func countWhereEqual(ints []int, target int) int {
	sum := 0

	for _, entry := range(ints) {
		if entry == target {
			sum++
		}
	}

	return sum
}

func countTrue(bools []bool) int {
	sum := 0

	for _, entry := range(bools) {
		if entry {
			sum++
		}
	}

	return sum
}
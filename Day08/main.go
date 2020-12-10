package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func check (err error) {
	if err != nil {
		panic(err)
	}
}

type instruction struct {
	code string
	offset int 
}

var instructions = make([]instruction, 0)

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		components := strings.Split(row, " ")

		offset := parseOffset(components[1])
		
		instructions = append(instructions, instruction{components[0], offset})
	}

	seenInstructions := make(map[int]bool)

	i := 0
	acc := 0
	for seenInstructions[i] != true {
		seenInstructions[i] = true
		i, acc = processInstruction(i, acc)
	}

	fmt.Printf("The accumulator is: %d\n", acc)

	// for _, inst := range(instructions) {
	// 	fmt.Println(inst)

	// }
}

func processInstruction(index int, acc int) (int, int) {
	inst := instructions[index]

	switch (inst.code) {
	case "nop":
		index++
	case "acc":
		acc += inst.offset
		index++
	case "jmp":
		index += inst.offset
	default:
		panic(inst.code)
	}

	return index, acc
}

func parseOffset(offset string) int {
	sign := offset[0]
	value, err := strconv.Atoi(offset[1:])
	check(err)

	if sign == '-' {
		return -1 * value
	}
	return value
}
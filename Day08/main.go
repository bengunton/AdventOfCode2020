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

	fixedInstructions := make([]instruction, len(instructions))

	for i, _ := range(instructions) {
		_ = copy(fixedInstructions, instructions)

		if flipInstruction(i, fixedInstructions) {
			acc, isFixed := runProgram(fixedInstructions)

			if (isFixed) {
				fmt.Printf("The accumulator is: %d\n", acc)
				break
			}
		}
	}

}

func printInstructions(is []instruction) {
	for _, inst := range(is) {
		fmt.Println(inst)

	}
}

func flipInstruction(index int, is []instruction) bool {
	inst := is[index]

	switch inst.code {
	case "nop":
		is[index] = instruction{"jmp", inst.offset}
		return true
	case "jmp":
		is[index] = instruction{"nop", inst.offset}
		return true
	default:
		return false 
	}

}

func runProgram(program []instruction) (int, bool) {
	seenInstructions := make(map[int]bool)

	i := 0
	acc := 0
	for seenInstructions[i] != true {
		seenInstructions[i] = true
		i, acc = processInstruction(i, acc, program)

		if i == (len(program)) {
			return acc, true
		}
	}

	return acc, false
}

func processInstruction(index int, acc int, program []instruction) (int, int) {
	inst := program[index]

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
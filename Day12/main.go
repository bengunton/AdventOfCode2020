package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Direction int

const (
	North Direction = 0
	East Direction = 90
	South Direction = 180
	West Direction = 270
)

type Boat struct {
	direction Direction
	eastOffset int
	northOffset int
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	boat := Boat{East, 0, 0}

	for scanner.Scan() {
		instruction := scanner.Text()
		boat.move(instruction)
	}

	fmt.Printf("Manhatten distance is %d\n", boat.manhattenDistance())
}

func (boat *Boat) move(instruction string) {
	code := instruction[0]
	amount, err := strconv.Atoi(instruction[1:])
	check(err)

	switch code {
	case 'F':
		boat.moveForward(amount)
	case 'N': 
		boat.northOffset += amount
	case 'S': 
		boat.northOffset -= amount
	case 'E': 
		boat.eastOffset += amount
	case 'W': 
		boat.eastOffset -= amount
	case 'L':
		boat.turnRight(-amount)
	case 'R':
		boat.turnRight(amount)
	}
}

func (boat *Boat) moveForward(amount int) {
	switch (boat.direction + 360) % 360 {
	case 0: 
		boat.northOffset += amount
	case 180: 
		boat.northOffset -= amount
	case 90: 
		boat.eastOffset += amount
	case 270: 
		boat.eastOffset -= amount
	}
}

func (boat *Boat) turnRight(amount int) {
	boat.direction = Direction((int(boat.direction) + amount) % 360)
}

func (boat *Boat) manhattenDistance() int {
	return abs(boat.eastOffset) + abs(boat.northOffset)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
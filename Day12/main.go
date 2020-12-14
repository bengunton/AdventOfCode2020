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

type Direction int64

const (
	North Direction = 0
	East Direction = 90
	South Direction = 180
	West Direction = 270
)

type Boat struct {
	direction Direction
	eastOffset int64
	northOffset int64
}

type Waypoint struct {
	eastOffset int64
	northOffset int64
}

func main() {
	distance := runFile("input")
	fmt.Printf("Manhatten distance is %d\n", distance)
}

func runFile(fileName string) int64 {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	boat := Boat{North, 0, 0}
	wp := Waypoint{10, 1}

	for scanner.Scan() {
		instruction := scanner.Text()
		wp = boat.moveWaypoint64(instruction, wp)
	}

	return boat.manhattenDistance()
}

func (boat *Boat) move(instruction string) {
	code := instruction[0]
	amounti, err := strconv.Atoi(instruction[1:])
	check(err)
	amount := int64(amounti)

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

func (boat *Boat) moveForward(amount int64) {
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

func (boat *Boat) moveForwardToWaypoint64(eastOff, northOff int64) {
	switch (boat.direction + 360) % 360 {
	case 0: 
		boat.northOffset += northOff
		boat.eastOffset += eastOff
	case 180: 
		boat.northOffset -= northOff
		boat.eastOffset -= eastOff
	case 90: 
		boat.eastOffset += northOff
		boat.northOffset -= eastOff
	case 270: 
		boat.eastOffset -= northOff
		boat.northOffset += eastOff
	}
}

func (boat *Boat) moveWaypoint64(instruction string, wp Waypoint) Waypoint {
	code := instruction[0]
	amounti, err := strconv.Atoi(instruction[1:])
	check(err)
	amount := int64(amounti)

	switch code {
	case 'F':
		for i := int64(0); i < amount; i++ {
			boat.moveForwardToWaypoint64(wp.eastOffset, wp.northOffset)
		}
	case 'N': 
		switch boat.direction {
		case North:
			wp.northOffset += amount
		case East:
			wp.eastOffset -= amount
		case South:
			wp.northOffset -= amount
		case West:
			wp.eastOffset += amount
		}
	case 'S': 
		switch boat.direction {
		case 0:
			wp.northOffset -= amount
		case 90:
			wp.eastOffset += amount
		case 180:
			wp.northOffset += amount
		case 270:
			wp.eastOffset -= amount
		}
	case 'E': 
		switch boat.direction {
		case 0:
			wp.eastOffset += amount
		case 90:
			wp.northOffset += amount
		case 180:
			wp.eastOffset -= amount
		case 270:
			wp.northOffset -= amount
		}
	case 'W': 
		switch boat.direction {
		case 0:
			wp.eastOffset -= amount
		case 90:
			wp.northOffset -= amount
		case 180:
			wp.eastOffset += amount
		case 270:
			wp.northOffset += amount
		}
	case 'L':
		boat.turnRight(-amount)
	case 'R':
		boat.turnRight(amount)
	}

	return wp
}

func (boat *Boat) turnRight(amount int64) {
	boat.direction = Direction((int64(boat.direction) + amount + 360) % 360)
}

func (boat *Boat) manhattenDistance() int64 {
	return abs(boat.eastOffset) + abs(boat.northOffset)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
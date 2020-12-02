package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

type passwordCheck struct {
	x int
	y int
	letterToCheck string
	password string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	passwordsToCheck := make([]passwordCheck, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwordCheck, err := parsePasswordCheck(scanner.Text())
		check(err)

		passwordsToCheck = append(passwordsToCheck, passwordCheck)
    }

	validCount := 0
	for _, password := range(passwordsToCheck) {
		if password.isValid (){
			validCount++
		}
	}
	
	fmt.Printf("Valid passwords: %d\n", validCount)
}

func parsePasswordCheck(input string) (passwordCheck, error) {
	s := strings.SplitN(input, "-", 2)
	x, err := strconv.Atoi(s[0])
	if err != nil {
		return passwordCheck{}, err
	}

	s = strings.SplitN(s[1], " ", 2)
	y, err := strconv.Atoi(s[0])
	if err != nil {
		return passwordCheck{}, err
	}
	
	s = strings.SplitN(s[1], ":", 2)
	characterToCheck := strings.TrimSpace(s[0])
	password := strings.TrimSpace(s[1])

	return passwordCheck{x, y, characterToCheck, password}, nil
}

func (p *passwordCheck) isValid() bool {
	count := strings.Count(p.password, p.letterToCheck)
	return count >= p.x && count <= p.y
}
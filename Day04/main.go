package main

import (
	"io/ioutil"
	"regexp"
	"fmt"
)

func check(e error)  {
	if (e != nil) {
		panic(e)
	}
}

const PotentiallyCompleteRecordRegex = "([a-zA-Z0-9#]*:[a-zA-Z0-9#]*( |\r\n)){7,8}"
const KeyRegex = "[a-zA-Z0-9#]*:"

func main() {
	input, err := ioutil.ReadFile("./input")
	check(err)

	stringInput := string(input)
	potentialRecords := getPotentiallyCompleteRecords(stringInput)
	validRecords := validateRecords(potentialRecords)

	fmt.Printf("There are %d valid records\n", len(validRecords))
}

func getPotentiallyCompleteRecords(input string) []string {
	regex, err := regexp.Compile(PotentiallyCompleteRecordRegex)
	// regex, err := regexp.Compile("cid:[[:alnum:]]* ")
	check(err)

	matches := regex.FindAllString(input, -1)
	return matches
}

func validateRecords(records []string) []string {
	var validRecords []string

	for _, record := range(records) {
		if (isValidRecord(record)) {
			validRecords = append(validRecords, record)
		}
	}

	return validRecords
}

func isValidRecord(record string) bool {
	regex, err := regexp.Compile(KeyRegex)
	check(err)

	keys := regex.FindAllString(record, -1)
	if len(keys) == 8 {
		return true
	}

	switch len(keys) {
	case 8:
		return true
	case 7:
		for _, key := range(keys) {
			if key == "cid:" {
				return false
			}
		}
		return true
	default:
		return false
	}
}
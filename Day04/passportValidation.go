package main

import (
	"strconv"
	"regexp"
	"strings"
)

var PassportValidations = map[string]func(string) bool {
	"byr": validateBirthYear,
	"iyr": validateIssueYear,
	"eyr": validateExpirationYear,
	"hgt": validateHeight,
	"hcl": validateHairColour,
	"ecl": validateEyeColour,
	"pid": validatePassportId,
}
var RequiredKeys = []string {
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func validateBirthYear(input string) bool {
	return validateRange(input, 1920, 2002)
}

func validateIssueYear(input string) bool {
	return validateRange(input, 2010, 2020)
}

func validateExpirationYear(input string) bool {
	return validateRange(input, 2020, 2030)
}

func validateHeight(input string) bool {
	regex, _ := regexp.Compile("^([0-9]*)(in|cm)$")
	matches := regex.FindStringSubmatch(strings.TrimSpace(input))
	if len(matches) != 3 {
		return false
	}

	switch matches[2] {
	case "in":
		return validateRange(matches[1], 59, 76)
	case "cm":
		return validateRange(matches[1], 150, 193)
	default:
		return false
	}
}

func validateHairColour(input string) bool {
	regex, _ := regexp.Compile("^#[0-9a-f]{6}$")
	return regex.MatchString(strings.TrimSpace(input))
}

func validateEyeColour(input string) bool {
	regex, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	return regex.MatchString(strings.TrimSpace(input))
}

func validatePassportId(input string) bool {
	regex, _ := regexp.Compile("^[0-9]{9}$")
	return regex.MatchString(strings.TrimSpace(input))
}

func validateRange(input string, min int, max int) bool {
	date, err := strconv.Atoi(input)
	if (err != nil) {
		return false
	}

	return date >= min && date <= max 
}
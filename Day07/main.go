package main

import(
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var RuleRegex = "^([a-z ]+) bags contain (?:((?:[0-9]+ [a-z ]+ bags?,? ?)+)|no other bags)\\."
var ChildRegex = "[0-9]+ ((?:[a-z ])+) bags?"

var containedIn = make(map[string][]string)

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		regex, err := regexp.Compile(RuleRegex)
		check(err)

		matches := regex.FindStringSubmatch(row)
		updateContainedIn(matches[1], strings.Split(matches[2], ","))
	}

	// for key, values := range(containedIn) {
	// 	fmt.Printf("%s: %v\n", key, values)
	// }

	foundParents := make(map[string]bool)

	fmt.Printf("There are %d parents", countParents("shiny gold", foundParents))
}

func updateContainedIn(subject string, children []string) {

	regex := regexp.MustCompile(ChildRegex)
	for _, child := range(children) {
		matches := regex.FindStringSubmatch(child)
		if matches == nil {
			continue
		}
	
		childName := matches[1]

		if containedIn[childName] == nil {
			containedIn[childName] = make([]string, 0)
		}

		containedIn[childName] = append(containedIn[childName], subject)
	}
}

func countParents(target string, foundParents map[string]bool) int {
	parents := containedIn[target]
	if len(parents) == 0 {
		return 0
	} 

	sum := 0
	for _, parent := range(parents) {

		if (foundParents[parent]) {
			continue
		}

		foundParents[parent] = true
		sum += 1 + countParents(parent, foundParents)
	}
	return sum
}
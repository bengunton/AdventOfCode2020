package main

import(
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type parent struct {
	count int
	name string
}

var RuleRegex = "^([a-z ]+) bags contain (?:((?:[0-9]+ [a-z ]+ bags?,? ?)+)|no other bags)\\."
var ChildRegex = "([0-9]+) ((?:[a-z ])+) bags?"

var containedIn = make(map[string][]string)
var contains = make(map[string][]parent)

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
		updateContains(matches[1], strings.Split(matches[2], ","))
	}

	foundParents := make(map[string]bool)

	fmt.Printf("There are %d parents\n", countParents("shiny gold", foundParents))

	fmt.Printf("There are %d boxes in shiny gold", countContaining("shiny gold") - 1)
}

func updateContainedIn(subject string, children []string) {

	regex := regexp.MustCompile(ChildRegex)
	for _, child := range(children) {
		matches := regex.FindStringSubmatch(child)
		if matches == nil {
			continue
		}
	
		childName := matches[2]

		if containedIn[childName] == nil {
			containedIn[childName] = make([]string, 0)
		}

		containedIn[childName] = append(containedIn[childName], subject)
	}
}

func updateContains(subject string, children []string) {

	regex := regexp.MustCompile(ChildRegex)
	for _, child := range(children) {
		matches := regex.FindStringSubmatch(child)
		if matches == nil {
			contains[subject] = make([]parent, 0)
			continue
		}
	
		count, err := strconv.Atoi(matches[1])
		check(err)

		childName := matches[2]

		if contains[subject] == nil {
			contains[subject] = make([]parent, 0)
		}

		contains[subject] = append(contains[subject], parent{count, childName})
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

func countContaining(target string) int {
	children := contains[target]

	sum := 1
	for _, child := range(children) {
		sum += child.count * countContaining(child.name)
	}
	return sum
}
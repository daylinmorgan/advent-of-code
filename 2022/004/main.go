package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitFn(c rune) bool {
	return c == '\n'
}

func readFileLines(filename string) []string {

	dat, err := os.ReadFile(filename)
	check(err)
	return strings.FieldsFunc(string(dat), splitFn)
}

func expandAssignment(assignment string) []int {
	ends := strings.Split(assignment, "-")
	start, err := strconv.Atoi(ends[0])
	check(err)
	finish, err := strconv.Atoi(ends[1])
	check(err)

	var sections []int
	for i := start; i <= finish; i++ {
		sections = append(sections, i)
	}

	return sections
}

// func compareAssignments(assignments []map[string]int) bool {
func compareAssignments(assignments [][]int) bool {
	var allSections []int
	var maxSize int
	var totalSize int
	for _, assignment := range assignments {
		if len(assignment) >= maxSize {
			// maxSize = len(assignment)
			totalSize += len(assignment)
		}
		allSections = append(allSections, assignment...)
	}
	allSections = dedupe(allSections)

	return len(allSections) != totalSize
}

// copy paste :/
func dedupe(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	allAssignments := readFileLines("input.txt")
	// allAssignments := readFileLines("input-head.txt")
	var overlappedAssignments int

	for _, pair := range allAssignments {
		assignments := strings.Split(pair, ",")
		var assignmentsRange [][]int
		for _, assignment := range assignments {
			assignmentsRange = append(assignmentsRange, expandAssignment(assignment))
		}
		if compareAssignments(assignmentsRange) {
			overlappedAssignments++
		}
	}
	fmt.Println("Total overlapped assignments:", overlappedAssignments)
}

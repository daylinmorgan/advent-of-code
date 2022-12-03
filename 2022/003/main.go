package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileLines(filename string) []string {

	splitFn := func(c rune) bool {
		return c == '\n'
	}
	dat, err := os.ReadFile(filename)
	check(err)
	return strings.FieldsFunc(string(dat), splitFn)
}

func getSharedItem(bag string) rune {
	size := len(bag)
	comp1, comp2 := bag[:size/2], bag[size/2:]
	var sharedType rune
	for _, c1 := range comp1 {
		for _, c2 := range comp2 {
			if c1 == c2 {
				sharedType = c1
			}
		}
	}
	return sharedType

}

func priorityKey() map[rune]int {

	priority := make(map[rune]int)
	priorityCounter := 1

	// iterate through a to z
	for r := 'a'; r <= 'z'; r++ {
		priority[r] = priorityCounter
		priority[unicode.ToUpper(r)] = priorityCounter + 26
		priorityCounter++
	}
	return priority
}

func main() {
	priorities := priorityKey()
	prioritySum := 0

	bags := readFileLines("input.txt")

	for _, bag := range bags {
		prioritySum += priorities[getSharedItem(bag)]
	}
	fmt.Println("Sum of priorities is", prioritySum)
}

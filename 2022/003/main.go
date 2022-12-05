package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// https://stackoverflow.com/a/69477128
func chunkSlice(items []string, chunkSize int32) (chunks [][]string) {
	//While there are more items remaining than chunkSize...
	for chunkSize < int32(len(items)) {
		//We take a slice of size chunkSize from the items array and append it to the new array
		chunks = append(chunks, items[0:chunkSize])
		//Then we remove those elements from the items array
		items = items[chunkSize:]
	}
	//Finally we append the remaining items to the new array and return it
	return append(chunks, items)
}

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

// part one
//
//	func getSharedItem(bag string) rune {
//		size := len(bag)
//		comp1, comp2 := bag[:size/2], bag[size/2:]
//		var sharedType rune
//		for _, c1 := range comp1 {
//			for _, c2 := range comp2 {
//				if c1 == c2 {
//					sharedType = c1
//				}
//			}
//		}
//		return sharedType
//	}

func getSharedItem(bags []string) rune {
	// size := len(bag)
	// comp1, comp2 := bag[:size/2], bag[size/2:]
	var sharedType rune
	for _, c1 := range bags[0] {
		for _, c2 := range bags[1] {
			for _, c3 := range bags[2] {
				if c1 == c2 && c2 == c3 {
					sharedType = c1
				}
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

	// bags := readFileLines("input-head.txt")
	bags := readFileLines("input.txt")
	elfGroups := chunkSlice(bags, 3)

	for _, bags := range elfGroups {
		prioritySum += priorities[getSharedItem(bags)]
	}
	fmt.Println("Sum of priorities is", prioritySum)
}

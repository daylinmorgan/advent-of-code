package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	// "github.com/daylinmorgan/advent-of-code/cast"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	return traverseBuffer(input, 4)
}

func part2(input string) int {
	return traverseBuffer(input, 14)
}

func traverseBuffer(input string, packetSize int) int {
	for i := 0; i < (len(input) - packetSize); i++ {
		if checkUniq(input[i : i+packetSize]) {
			return i + packetSize
		} else {
			continue
		}
	}
	panic("no marker found")
}

// check whether string containing characters are all unique
func checkUniq(chars string) bool {
	charKey := make(map[rune]bool)
	for _, r := range chars {
		if _, ok := charKey[r]; ok {
			return false
		} else {
			charKey[r] = true
		}

	}
	return true
}

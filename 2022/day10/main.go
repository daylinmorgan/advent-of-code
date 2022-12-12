package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
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
	instructions := parseInput(input)
	cycles := []int{1}
	for _, instruction := range instructions {
		cycles = executeInstruction(cycles, instruction)
	}
	signalStrength := 0
	for _, i := range []int{20, 60, 100, 140, 180, 220} {
		signalStrength += i * cycles[i-1]
	}
	return signalStrength
}

func part2(input string) string {
	instructions := parseInput(input)
	cycles := []int{1}
	for _, instruction := range instructions {
		cycles = executeInstruction(cycles, instruction)
	}
	var crt [][]string
	crt = renderCRT(crt, cycles)
	display := displayCRT(crt)
	return display
}

func parseInput(input string) (instructions []map[string]string) {
	// since example is embedded in main_test.go
	input = strings.TrimRight(input, "\n")

	for _, line := range strings.Split(input, "\n") {
		if line == "noop" {
			instructions = append(instructions, map[string]string{"program": "noop", "value": ""})
		} else if strings.HasPrefix(line, "addx") {
			instructions = append(instructions, map[string]string{"program": "addx", "value": strings.Split(line, " ")[1]})
		}
	}
	return instructions
}

func executeInstruction(cycles []int, instruction map[string]string) []int {
	if instruction["program"] == "noop" {
		cycles = append(cycles, cycles[len(cycles)-1])
	} else {
		v, err := strconv.Atoi(instruction["value"])
		if err != nil {
			panic(err)
		}
		cycles = append(cycles, cycles[len(cycles)-1])
		cycles = append(cycles, cycles[len(cycles)-1]+v)
	}

	return cycles
}

func renderCRT(crt [][]string, cycles []int) [][]string {
	rowNum := 0
	var row []string
	var x int
	for i, register := range cycles {
		x = i - (40 * rowNum)
		if x-1 <= register && register <= x+1 {
			row = append(row, "#")
		} else {
			row = append(row, ".")
		}
		if len(row) == 40 {
			crt = append(crt, row)
			rowNum++
			row = nil
		}
	}
	return crt
}

func displayCRT(crt [][]string) string {
	display := "\n"
	var toDisplay string
	for row, chars := range crt {
		toDisplay = ""
		for _, c := range chars {
			toDisplay += c
		}
		display += fmt.Sprintf("Row %d --> %s\n", row, toDisplay)
	}
	return display
}

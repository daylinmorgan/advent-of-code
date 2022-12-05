package main

import (
	"fmt"
	"os"
	"regexp"
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

func parseCargo(cargo string) map[int][]string {
	start := make(map[int][]string)
	cargoRows := strings.FieldsFunc(cargo, splitFn)
	numRows := len(cargoRows)
	revCargoRows := make([]string, numRows)
	for i, row := range cargoRows {
		revCargoRows[numRows-i-1] = row
	}
	for i, c := range revCargoRows[0] {
		if string(c) != " " {
			for _, row := range revCargoRows[1:] {
				if i > len(row) || row[i] == ' ' {
					continue
				}
				// is strconv necessary here?
				stack, err := strconv.Atoi(string(c))
				check(err)
				if _, ok := start[stack]; !ok {
					start[stack] = []string{string(row[i])}
				} else {
					start[stack] = append(start[stack], string(row[i]))
				}
			}
		}
	}
	return start
}

func parseMove(s string) map[string]int {
	re := regexp.MustCompile(`move (?P<num>\d+) from (?P<start>\d) to (?P<end>\d)`)
	match := re.FindStringSubmatch(s)
	moveMap := make(map[string]int)
	for i, name := range re.SubexpNames()[1:] {
		intVal, err := strconv.Atoi(match[i+1])
		check(err)
		moveMap[name] = intVal
	}
	return moveMap
}

func parseMoves(movesDef string) []map[string]int {
	var moves []map[string]int
	splitMoves := strings.FieldsFunc(movesDef, splitFn)
	for _, move := range splitMoves {
		moves = append(moves, parseMove(move))
	}
	return moves
}

func moveCrates(move map[string]int, cargo map[int][]string) map[int][]string {
	startStack, endStack, numCrates := move["start"], move["end"], move["num"]
	for i := 0; i < numCrates; i++ {
		sizeStartStack := len(cargo[startStack])
		cargo[endStack] = append(cargo[endStack], cargo[startStack][sizeStartStack-1])
		cargo[startStack] = cargo[startStack][:sizeStartStack-1]
	}
	return cargo
}

func readFile(filename string) (map[int][]string, []map[string]int) {

	dat, err := os.ReadFile(filename)
	check(err)
	s := strings.Split(string(dat), "\n\n")
	cargo := parseCargo(s[0])
	moves := parseMoves(s[1])
	return cargo, moves
}

func main() {
	// cargo, moves := readFile("example.txt")
	cargo, moves := readFile("input.txt")
	for _, move := range moves {
		cargo = moveCrates(move, cargo)
	}

	var message string
	msgchars := make([]string, len(cargo))
	for stack, crates := range cargo {
		fmt.Printf("stack:%d, top crate: %s\n", stack, crates[len(crates)-1])
		msgchars[stack-1] = crates[len(crates)-1]
	}
	for _, char := range msgchars {
		message += char
	}
	fmt.Println("Message:", message)
}

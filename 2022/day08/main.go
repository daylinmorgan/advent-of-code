package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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
	trees := parseInput(input)

	numTrees := forEachInnerTree(trees, findVisibleTrees)

	// include edge trees
	numTrees += (len(trees) + (len(trees) - 2)) * 2

	return numTrees
}

func part2(input string) int {
	trees := parseInput(input)
	bestScore := forEachInnerTree(trees, calculateScenicScore)
	return bestScore
}

func createRowArray(row string) []uint8 {
	a := make([]uint8, len(row))
	for i, numStr := range row {
		numInt, err := strconv.ParseUint(string(numStr), 10, 8)
		if err != nil {
			panic(err)
		}
		a[i] = uint8(numInt)
	}
	return a

}

func parseInput(input string) [][]uint8 {
	rows := strings.Split(input, "\n")
	a := make([][]uint8, len(rows))
	for i, row := range rows {
		a[i] = createRowArray(row)
	}
	return a
}

func forEachInnerTree(trees [][]uint8, f func(trees [][]uint8, x int, y int, score *int)) int {
	// score could be scenic score or num of trees
	score := 0
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees)-1; j++ {
			f(trees, i, j, &score)
		}
	}

	return score
}

// I can't think of anythin better than this
// make each slice start from x,y and go towards edge
func sliceTrees(trees [][]uint8, x int, y int, direction string) []uint8 {
	var newSlice []uint8
	switch direction {
	case "top":
		for i := x - 1; i >= 0; i-- {
			newSlice = append(newSlice, trees[i][y])
		}
	case "bottom":
		for i := x + 1; i < len(trees); i++ {
			newSlice = append(newSlice, trees[i][y])
		}
	case "left":
		for j := y - 1; j >= 0; j-- {
			newSlice = append(newSlice, trees[x][j])
		}
	case "right":
		for j := y + 1; j < len(trees); j++ {
			newSlice = append(newSlice, trees[x][j])
		}
	default:
		panic("Unknown direction")
	}
	return newSlice
}

var directions = [4]string{"left", "right", "top", "bottom"}

func isTallest(sliceTrees []uint8, h uint8) bool {
	isTaller := 0
	for _, t := range sliceTrees {
		if t >= h {
			isTaller++
		}
	}
	return isTaller == 0
}

func findVisibleTrees(trees [][]uint8, x int, y int, score *int) {
	notVisible := 0
	for _, d := range directions {
		if !isTallest(sliceTrees(trees, x, y, d), trees[x][y]) {
			notVisible++
		}
	}

	if notVisible < 4 {
		*score = *score + 1
	}
}

func calculateScenicScore(trees [][]uint8, x int, y int, bestScore *int) {
	var score int
	for _, d := range directions {
		s := calculateScenicScoreOneWay(sliceTrees(trees, x, y, d), trees[x][y])
		if score == 0 {
			score = s
		} else {
			score *= s
		}
	}
	if score > *bestScore {
		*bestScore = score
	}
}

func calculateScenicScoreOneWay(s []uint8, h uint8) int {
	var score int

	// it can always see the closest tree
	if len(s) == 1 {
		return 1
	}
	for _, t := range s {
		score++
		if t >= h {
			break
		}
	}
	return score
}

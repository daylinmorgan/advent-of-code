package main

import (
	"fmt"
	"os"
	"strings"
)

// need each value as a map?
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// A/X -> Rock
// B/Y -> Paper
// C/Z -> Scissors
//
// Part One Scoring
// func scoreRound(myPlay string, oppPlay string) int {
// 	scoreKey := map[string]string{"X": "C", "Z": "B", "Y": "A"}
// 	scorer := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
//
// 	if scorer[myPlay] == scorer[oppPlay] {
// 		return scorer[myPlay] + 3
// 	} else if scoreKey[myPlay] == oppPlay {
// 		return scorer[myPlay] + 6
// 	} else {
// 		return scorer[myPlay]
// 	}
// }

// X -> lose
// Y -> draw
// Z -> win

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func scoreRound(outcome string, play string) int {
	var score int
	loseKey := map[string]string{"A": "C", "B": "A", "C": "B"}
	winKey := reverseMap(loseKey)
	scorer := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

	switch outcome {
	case "X":
		score = scorer[loseKey[play]]
	case "Y":
		score = scorer[play] + 3
	case "Z":
		score = scorer[winKey[play]] + 6
	}
	return score
}

func main() {

	var score int

	splitFn := func(c rune) bool {
		return c == '\n'
	}

	fmt.Println("Rock-Paper-Scissors!")

	dat, err := os.ReadFile("input")
	check(err)

	rounds := strings.FieldsFunc(string(dat), splitFn)

	for _, line := range rounds {
		plays := strings.Split(line, " ")
		score = score + scoreRound(plays[1], plays[0])
	}
	fmt.Println("My final score ->", score)
}

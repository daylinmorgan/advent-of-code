package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func part1(input string) int {
	packetPairs := parseInput(input)
	correct := 0
	for i, pair := range packetPairs {
		if compare(pair[0], pair[1]) <= 0 {
			correct += i + 1
		}
	}

	return correct
}

func part2(input string) int {
	packetPairs := parseInput(input)
	packets := []any{}
	for _, pairs := range packetPairs {
		packets = append(packets, pairs...)
	}
	packets = append(packets, []any{[]any{2.0}}, []any{[]any{6.0}})
	sort.Slice(packets, func(i, j int) bool { return compare(packets[i], packets[j]) <= 0 })

	decoder := 1
	for i, p := range packets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			decoder *= i + 1
		}
	}
	return decoder
}

func extractPackets(packets string) []any {
	var a, b any
	s := strings.Split(packets, "\n")
	json.Unmarshal([]byte(s[0]), &a)
	json.Unmarshal([]byte(s[1]), &b)
	return []any{a, b}
}

func compare(a any, b any) int {
	var ax, bx []any
	var singleA, singleB bool

	switch a.(type) {
	case float64:
		ax, singleA = []any{a}, true
	case []any, []float64:
		ax = a.([]any) // converting to any?
	}

	switch b.(type) {
	case float64:
		bx, singleB = []any{b}, true
	case []any, []float64:
		bx = b.([]any) // converting to any?
	}

	if singleA && singleB {
		return int(ax[0].(float64) - bx[0].(float64))
	}

	for i := 0; i < len(ax) && i < len(bx); i++ {
		if c := compare(ax[i], bx[i]); c != 0 {
			return c
		}
	}

	return len(ax) - len(bx)
}

func parseInput(input string) (packetPairs [][]any) {
	for _, packets := range strings.Split(input, "\n\n") {
		packetPairs = append(packetPairs, extractPackets(packets))
	}
	return packetPairs
}

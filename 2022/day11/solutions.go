package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part1(input string) int {
	reliefFactor, rounds := 3, 20
	monkeys := parseInput(input)
	ledger := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for idx := range monkeys {
			ledger[idx] = ledger[idx] + len(monkeys[idx].items)
			monkeys = throwItems(monkeys, idx, reliefFactor)
		}
	}
	sort.Ints(ledger)
	return ledger[len(ledger)-1] * ledger[len(ledger)-2]

}

func part2(input string) int {
	reliefFactor, rounds := 1, 10000

	monkeys := parseInput(input)
	ledger := make([]int, len(monkeys))

	for i := 0; i < rounds; i++ {
		for idx := range monkeys {
			ledger[idx] = ledger[idx] + len(monkeys[idx].items)
			monkeys = throwItems(monkeys, idx, reliefFactor)
		}
	}

	sort.Ints(ledger)
	return ledger[len(ledger)-1] * ledger[len(ledger)-2]
}

func parseInput(input string) (monkeys []Monkey) {
	for _, monkeyDef := range strings.Split(input, "\n\n") {
		monkeys = append(monkeys, newMonkey(monkeyDef))
	}
	return monkeys
}

type Monkey struct {
	items       []int
	operation   Operation
	test        int
	conditional map[bool]string
}

// add modfactor since part2 has big old numbers
func (m *Monkey) turn(reliefFactor int, modfactor int) (throws []map[string]int) {
	if len(m.items) == 0 {
		return
	}

	var newItem, newMonkey int
	var err error

	for _, item := range m.items {
		if reliefFactor == 1 {
			newItem = m.operation.calc(item)
			newItem = newItem % modfactor
		} else {
			newItem = m.operation.calc(item) / reliefFactor
		}
		newMonkey, err = strconv.Atoi(m.conditional[(newItem%m.test) == 0])
		if err != nil {
			panic(err)
		}
		throws = append(throws, map[string]int{"item": newItem, "monkey": newMonkey})
	}

	return throws
}

type Operation struct {
	left   string
	symbol string
	right  string
}

func (o *Operation) calc(item int) int {
	var left, right, newItem int
	var err error

	if o.left == "old" {
		left = item
	} else {
		left, err = strconv.Atoi(o.left)
		if err != nil {
			panic(err)
		}
	}
	if o.right == "old" {
		right = item
	} else {
		right, err = strconv.Atoi(o.right)
		if err != nil {
			panic(err)
		}
	}
	switch o.symbol {
	case "*":
		newItem = left * right
	case "+":
		newItem = left + right
	default:
		panic(fmt.Sprint("unknown operation", o.symbol))
	}
	return newItem
}

func newMonkey(monkeyDef string) Monkey {
	lines := strings.Split(monkeyDef, "\n")
	return Monkey{items: newItems(lines[1]), operation: newOperation(lines[2]), test: parseTest(lines[3]), conditional: parseConditional(lines[4], lines[5])}
}

func newOperation(operationLine string) Operation {
	s := strings.Split(strings.Split(operationLine, "new = ")[1], " ")
	return Operation{left: s[0], symbol: s[1], right: s[2]}
}

func newItems(itemLine string) []int {
	items := strings.Split(strings.Split(itemLine, ":")[1], ",")
	intItems := []int{}
	for _, item := range items {
		intItem, err := strconv.Atoi(strings.TrimSpace(item))
		if err != nil {
			panic(err)
		}
		intItems = append(intItems, intItem)
	}
	return intItems
}

func parseTest(testLine string) int {
	s := strings.Replace(strings.TrimSpace(testLine), "Test: divisible by ", "", 1)
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func parseConditional(trueString string, falseString string) map[bool]string {
	conditional := make(map[bool]string)
	conditional[true] = strings.Replace(strings.TrimSpace(trueString), "If true: throw to monkey ", "", 1)
	conditional[false] = strings.Replace(strings.TrimSpace(falseString), "If false: throw to monkey ", "", 1)
	return conditional
}

func throwItems(monkeys []Monkey, idx int, reliefFactor int) []Monkey {
	modfactor := 1
	for _, m := range monkeys {
		modfactor *= m.test
	}
	throws := monkeys[idx].turn(reliefFactor, modfactor)
	if throws != nil {
		for _, t := range throws {
			monkeys[t["monkey"]].items = append(monkeys[t["monkey"]].items, t["item"])
		}
		monkeys[idx].items = []int{}
	}

	return monkeys
}

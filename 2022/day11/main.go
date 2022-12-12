package main

import (
	_ "embed"
	"flag"
	"fmt"
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

//
// func part1(input string) int {
// 	monkeys := parseInput(input)
// 	ledger := make([]int, len(monkeys))
// 	for i := 0; i < 20; i++ {
// 		for idx := range monkeys {
// 			ledger[idx] = ledger[idx] + len(monkeys[idx].items)
// 			monkeys = throwItems(monkeys, idx)
// 		}
// 		fmt.Println("round", i)
// 		for j, m := range monkeys {
// 			fmt.Printf("Monkey %d: %s\n", j, fmt.Sprint(m.items))
// 		}
// 	}
// 	for i, m := range monkeys {
// 		fmt.Printf("Monkey %d: %s\n", i, fmt.Sprint(m.items))
// 	}
// 	fmt.Println("ledger", ledger)
// 	sort.Ints(ledger)
// 	return ledger[len(ledger)-1] * ledger[len(ledger)-2]
// }
//
// func part2(input string) int {
// 	return 0
// }
//
// func parseInput(input string) (monkeys []Monkey) {
// 	for _, monkeyDef := range strings.Split(input, "\n\n") {
// 		monkeys = append(monkeys, newMonkey(monkeyDef))
// 	}
// 	return monkeys
// }
//
// type Monkey struct {
// 	items       []int
// 	operation   Operation
// 	test        int
// 	conditional map[bool]string
// }
//
// func (m *Monkey) turn() (throws []map[string]int) {
// 	if len(m.items) == 0 {
// 		return
// 	}
//
// 	var newItem, newMonkey int
// 	var err error
//
// 	for _, item := range m.items {
// 		newItem = m.operation.calc(item) / 3
// 		newMonkey, err = strconv.Atoi(m.conditional[(newItem%m.test) == 0])
// 		if err != nil {
// 			panic(err)
// 		}
// 		throws = append(throws, map[string]int{"item": newItem, "monkey": newMonkey})
// 	}
//
// 	return throws
// }
//
// type Operation struct {
// 	left   string
// 	symbol string
// 	right  string
// }
//
// func (o *Operation) calc(item int) int {
// 	var left, right, newItem int
// 	var err error
//
// 	if o.left == "old" {
// 		left = item
// 	} else {
// 		left, err = strconv.Atoi(o.left)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	if o.right == "old" {
// 		right = item
// 	} else {
// 		right, err = strconv.Atoi(o.right)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	switch o.symbol {
// 	case "*":
// 		newItem = left * right
// 	case "+":
// 		newItem = left + right
// 	default:
// 		panic(fmt.Sprint("unknown operation", o.symbol))
// 	}
// 	return newItem
// }
//
// func newMonkey(monkeyDef string) Monkey {
// 	lines := strings.Split(monkeyDef, "\n")
// 	return Monkey{items: newItems(lines[1]), operation: newOperation(lines[2]), test: parseTest(lines[3]), conditional: parseConditional(lines[4], lines[5])}
// }
//
// func newOperation(operationLine string) Operation {
// 	s := strings.Split(strings.Split(operationLine, "new = ")[1], " ")
// 	return Operation{left: s[0], symbol: s[1], right: s[2]}
// }
//
// func newItems(itemLine string) []int {
// 	items := strings.Split(strings.Split(itemLine, ":")[1], ",")
// 	intItems := []int{}
// 	for _, item := range items {
// 		intItem, err := strconv.Atoi(strings.TrimSpace(item))
// 		if err != nil {
// 			panic(err)
// 		}
// 		intItems = append(intItems, intItem)
// 	}
// 	return intItems
// }
//
// func parseTest(testLine string) int {
// 	s := strings.Replace(strings.TrimSpace(testLine), "Test: divisible by ", "", 1)
// 	v, err := strconv.Atoi(s)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return v
// }
//
// func parseConditional(trueString string, falseString string) map[bool]string {
// 	conditional := make(map[bool]string)
// 	conditional[true] = strings.Replace(strings.TrimSpace(trueString), "If true: throw to monkey ", "", 1)
// 	conditional[false] = strings.Replace(strings.TrimSpace(falseString), "If false: throw to monkey ", "", 1)
// 	return conditional
// }
//
// func throwItems(monkeys []Monkey, idx int) []Monkey {
//
// 	throws := monkeys[idx].turn()
// 	if throws != nil {
// 		for _, t := range throws {
// 			monkeys[t["monkey"]].items = append(monkeys[t["monkey"]].items, t["item"])
// 		}
// 		monkeys[idx].items = []int{}
// 	}
//
// 	return monkeys
// }

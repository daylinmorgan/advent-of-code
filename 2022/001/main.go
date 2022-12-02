package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// todo: naming conventions?
func arr_str_to_int(arr []string) []int {
	var newarr = []int{}
	for _, i := range arr {
		j, err := strconv.Atoi(i)
		check(err)
		newarr = append(newarr, j)
	}
	return newarr
}

func sum(numarray []int) int {
	var arrSum int
	for i := 0; i < len(numarray); i++ {
		arrSum = arrSum + numarray[i]
	}
	return arrSum
}

func main() {
	maxCal := 0
	// dat, err := os.ReadFile("simple-input.txt")
	dat, err := os.ReadFile("input")
	check(err)

	// split into groups based on blank new line
	var calories = []int{}
	for _, arr := range strings.Split(string(dat), "\n\n") {
		sum := sum(arr_str_to_int(strings.Fields(arr)))
		calories = append(calories, sum)
		if sum > maxCal {
			maxCal = sum
		}
	}
	fmt.Printf("max calories are %d\n", maxCal)

	// sort.Ints(calories)
	// copy pasted .... https://stackoverflow.com/a/18343326
	// sort.IntSlice does what?
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	total := 0
	for i := 0; i < 3; i++ {
		total = total + calories[i]
	}
	fmt.Println("total of top 3 elves ->", total)
}

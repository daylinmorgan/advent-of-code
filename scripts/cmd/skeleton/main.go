package main

import (
	"flag"
	"time"

	"github.com/daylinmorgan/advent-of-code/scripts/skeleton"
)

func main() {
	today := time.Now()
	day := flag.Int("day", today.Day(), "day number to fetch, 1-25")
	year := flag.Int("year", today.Year(), "AOC year")
	flag.Parse()
	skeleton.Run(*day, *year)
}

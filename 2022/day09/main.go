package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

	rope := makeRope(2)
	path := startPath()
	moves := parseInput(input)

	for _, move := range moves {
		moveRope(&rope, move, &path)
	}
	return len(path.set)
}

func part2(input string) int {
	rope := makeRope(10)
	path := startPath()
	moves := parseInput(input)

	for _, move := range moves {
		moveRope(&rope, move, &path)
	}
	return len(path.set)

}

type move struct {
	direction string
	distance  int
}

type Coord struct {
	x int
	y int
}

type Rope struct {
	knots []*Coord
}

func makeRope(n int) (r Rope) {
	for i := 0; i < n; i++ {
		r.knots = append(r.knots, &Coord{x: 0, y: 0})
	}
	return r
}

type Path struct {
	points [][2]int
	set    map[[2]int]bool
}

func startPath() Path {
	p1 := [2]int{0, 0}
	m := make(map[[2]int]bool)
	m[p1] = true
	return Path{points: [][2]int{p1}, set: m}
}

func (p *Path) addPoint(pt [2]int) {
	p.points = append(p.points, pt)
	if p.set[pt] {
		return // Already in the map
	}
	p.set[pt] = true
}

func moveRope(r *Rope, m move, p *Path) {
	for i := 0; i < m.distance; i++ {
		moveCoordOnce(r.knots[0], m.direction, 1)
		for j := 0; j < len(r.knots)-1; j++ {
			moveTail(r.knots[j], r.knots[j+1], p, m.direction)
		}
		tail := r.knots[len(r.knots)-1]
		p.addPoint([2]int{tail.x, tail.y})
	}
}

func signedOne(i int) int {
	if i > 0 {
		return 1
	} else {
		return -1
	}
}

func moveTail(h *Coord, t *Coord, p *Path, dir string) {
	dx := h.x - t.x
	dy := h.y - t.y
	diagonal := math.Sqrt(float64(dx*dx + dy*dy))

	// move diagonally
	if diagonal > math.Sqrt(2) && (Abs(dx) == 1 || Abs(dy) == 1) {
		if Abs(dx) > Abs(dy) {
			t.x = t.x + signedOne(dx)
			t.y = h.y
		} else {
			t.x = h.x
			t.y = t.y + signedOne(dy)
		}
	} else if Abs(dx) > 1 && Abs(dy) == 0 {
		t.x = t.x + signedOne(dx)
	} else if Abs(dx) == 0 && Abs(dy) > 1 {
		t.y = t.y + signedOne(dy)
	} else if Abs(dx) <= 1 && Abs(dy) <= 1 {
	} else if Abs(dx) == Abs(dy) {
		t.x = t.x + signedOne(dx)
		t.y = t.y + signedOne(dy)
	} else {
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveCoordOnce(c *Coord, dir string, dist int) {
	switch dir {
	case "R":
		(*c).x = (*c).x + dist
	case "L":
		(*c).x = (*c).x - dist
	case "U":
		(*c).y = (*c).y + dist
	case "D":
		(*c).y = (*c).y - dist
	}
}

func parseInput(input string) (moves []move) {
	for _, line := range strings.Split(input, "\n") {
		s := strings.Split(line, " ")
		dist, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		moves = append(moves, move{direction: s[0], distance: dist})
	}
	return moves
}

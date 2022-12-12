package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

var part2examplans string = `
Row 0 --> ##..##..##..##..##..##..##..##..##..##..
Row 1 --> ###...###...###...###...###...###...###.
Row 2 --> ####....####....####....####....####....
Row 3 --> #####.....#####.....#####.....#####.....
Row 4 --> ######......######......######......####
Row 5 --> #######.......#######.......#######.....
`

var part2actualans string = `
Row 0 --> ###..#..#.###....##.###..###..#.....##..
Row 1 --> #..#.#.#..#..#....#.#..#.#..#.#....#..#.
Row 2 --> #..#.##...#..#....#.###..#..#.#....#..#.
Row 3 --> ###..#.#..###.....#.#..#.###..#....####.
Row 4 --> #.#..#.#..#....#..#.#..#.#....#....#..#.
Row 5 --> #..#.#..#.#.....##..###..#....####.#..#.
`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  13140,
		},
		{
			name:  "actual",
			input: input,
			want:  15120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: example,
			want:  part2examplans,
		},
		{
			name:  "actual",
			input: input,
			want:  part2actualans,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

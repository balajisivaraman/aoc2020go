package solutions

import (
	"testing"
)

func TestDay03Part1(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	want := 7
	actual := Day03Part1(input)
	if !(actual == want) {
		t.Errorf("Day02: Part 3 = %d, want %d", actual, want)
	}
}

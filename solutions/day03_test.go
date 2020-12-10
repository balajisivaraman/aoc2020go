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
		t.Errorf("Day03: Part 1 = %d, want %d", actual, want)
	}
}

func TestDay03Part2(t *testing.T) {
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
	want := 336
	actual := Day03Part2(input)
	if !(actual == want) {
		t.Errorf("Day03: Part 2 = %d, want %d", actual, want)
	}
}

package solutions

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	want := 514579
	actual := Day01Part1(input)
	if !(actual == want) {
		t.Errorf("Day01: Part 1(%d) = %d, want %d", input, actual, want)
	}
}

func TestDay01Part2(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	want := 241861950
	actual := Day01Part2(input)
	if !(actual == want) {
		t.Errorf("Day01: Part 2(%d) = %d, want %d", input, actual, want)
	}
}

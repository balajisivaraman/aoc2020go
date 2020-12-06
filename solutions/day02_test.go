package solutions

import (
	"testing"
)

func TestDay02Part1(t *testing.T) {
	input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	want := 2
	actual := Day02Part1(input)
	if !(actual == want) {
		t.Errorf("Day02: Part 1(%s) = %d, want %d", input, actual, want)
	}
}

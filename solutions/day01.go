package solutions

import (
	"github.com/emirpasic/gods/sets/hashset"
)

func Day01Part1(input []int) int {
	result := 0
	mappedSet := hashset.New()
	for _, val := range input {
		mappedSet.Add(val)
		var rem = 2020 - val
		if mappedSet.Contains(rem) {
			result = rem * val
		}
	}
	return result
}

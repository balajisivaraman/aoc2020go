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
			break
		}
	}
	return result
}

func Day01Part2(input []int) int {
	result := 0
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					result = input[i] * input[j] * input[k]
					break
				}
			}
		}
	}
	return result
}

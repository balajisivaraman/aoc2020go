package main

import (
	"fmt"
	"github.com/balajisivaraman/aoc2020go/solutions"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readLines(filepath string) []int {
	content, _ := ioutil.ReadFile(filepath)
	lines := strings.Split(string(content[:]), "\n")
	input := []int{}
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		input = append(input, num)
	}
	// below hack is needed because read lines apparently inserts
	// a blank at the end
	return input[:len(input)-1]
}

func main() {
	var result int
	day := os.Args[1]
	if day == "day01" {
		input := readLines("input/day01.txt")
		result = solutions.Day01Part1(input)
		fmt.Println("Day01, Part1 : ", result)
		result = solutions.Day01Part2(input)
		fmt.Println("Day01, Part2 : ", result)
	}
}

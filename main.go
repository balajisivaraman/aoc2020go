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
	return input
}

func main() {
	var result int
	day := os.Args[1]
	if day == "day01" {
		input := readLines("input/day01.txt")
		result = solutions.Day01Part1(input)
		fmt.Printf("Day01, Part1 : %d", result)
	}
}

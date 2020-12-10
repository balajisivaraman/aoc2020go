package main

import (
	"fmt"
	"github.com/balajisivaraman/aoc2020go/solutions"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readLinesAsStrings(filepath string) []string {
	content, _ := ioutil.ReadFile(filepath)
	lines := strings.Split(string(content[:]), "\n")
	// below hack is needed because read lines apparently inserts
	// a blank at the end
	return lines[:len(lines)-1]
}

func readLinesAsInts(filepath string) []int {
	lines := readLinesAsStrings(filepath)
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
		input := readLinesAsInts("input/day01.txt")
		result = solutions.Day01Part1(input)
		fmt.Println("Day01, Part1 : ", result)
		result = solutions.Day01Part2(input)
		fmt.Println("Day01, Part2 : ", result)
	}
	if day == "day02" {
		input := readLinesAsStrings("input/day02.txt")
		result = solutions.Day02Part1(input)
		fmt.Println("Day02, Part1 : ", result)
		result = solutions.Day02Part2(input)
		fmt.Println("Day02, Part2 : ", result)
	}
	if day == "day03" {
		input := readLinesAsStrings("input/day03.txt")
		result = solutions.Day03Part1(input)
		fmt.Println("Day03, Part1 : ", result)
		result = solutions.Day03Part2(input)
		fmt.Println("Day03, Part2 : ", result)
	}
}

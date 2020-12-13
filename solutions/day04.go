package solutions

import (
	"strconv"
	"strings"
)

type Passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColour     string
	eyeColour      string
	passportId     string
	countryId      string
}

func (passport *Passport) isValidPerOldMethod() bool {
	return passport.birthYear != "" &&
		passport.issueYear != "" &&
		passport.expirationYear != "" &&
		passport.height != "" &&
		passport.hairColour != "" &&
		passport.eyeColour != "" &&
		passport.passportId != ""
}

func parseStringToIntIgnoringError(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func parseIntoMap(parsedPassport map[string]string, parsedLine []string) {
	for _, item := range parsedLine {
		splitByColon := strings.Split(item, ":")
		parsedPassport[splitByColon[0]] = splitByColon[1]
	}
}

func parseInputToMap(input []string) []map[string]string {
	result := make([]map[string]string, 0)
	parsedPassport := make(map[string]string)
	counter := 0
	idx := 0
	for {
		if idx == len(input) {
			result = append(result, parsedPassport)
			break
		}
		line := input[idx]
		if line == "" {
			result = append(result, parsedPassport)
			parsedPassport = make(map[string]string)
			counter++
		} else {
			splitBySpaces := strings.Split(input[idx], " ")
			parseIntoMap(parsedPassport, splitBySpaces)
		}
		idx++
	}
	return result
}

func parseInput(input []string) []Passport {
	result := make([]Passport, 0)
	parsedMaps := parseInputToMap(input)
	for _, passportMap := range parsedMaps {
		passport := Passport{
			birthYear:      passportMap["byr"],
			issueYear:      passportMap["iyr"],
			expirationYear: passportMap["eyr"],
			height:         passportMap["hgt"],
			hairColour:     passportMap["hcl"],
			eyeColour:      passportMap["ecl"],
			passportId:     passportMap["pid"],
			countryId:      passportMap["cid"],
		}
		result = append(result, passport)
	}
	return result
}

func Day04Part1(input []string) int {
	passports := parseInput(input)
	validPassportCount := 0
	for _, passport := range passports {
		if passport.isValidPerOldMethod() {
			validPassportCount++
		}
	}
	return validPassportCount
}

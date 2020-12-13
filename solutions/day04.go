package solutions

import (
	"github.com/balajisivaraman/aoc2020go/passport"
	"strings"
)

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

func parseInput(input []string) []passport.Passport {
	result := make([]passport.Passport, 0)
	parsedMaps := parseInputToMap(input)
	for _, passportMap := range parsedMaps {
		passport := passport.Passport{
			BirthYear:      passportMap["byr"],
			IssueYear:      passportMap["iyr"],
			ExpirationYear: passportMap["eyr"],
			Height:         passportMap["hgt"],
			HairColour:     passportMap["hcl"],
			EyeColour:      passportMap["ecl"],
			PassportId:     passportMap["pid"],
			CountryId:      passportMap["cid"],
		}
		result = append(result, passport)
	}
	return result
}

func Day04Part1(input []string) int {
	passports := parseInput(input)
	validPassportCount := 0
	for _, passport := range passports {
		if passport.IsValidPerOldMethod() {
			validPassportCount++
		}
	}
	return validPassportCount
}

func Day04Part2(input []string) int {
	passports := parseInput(input)
	validPassportCount := 0
	for _, passport := range passports {
		if passport.IsValidPerNewMethod() {
			validPassportCount++
		}
	}
	return validPassportCount
}

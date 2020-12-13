package solutions

import (
	"regexp"
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

func isValidYear(strYear string, startYear int, endYear int) bool {
	if len(strYear) != 4 {
		return false
	}
	year := parseStringToIntIgnoringError(strYear)
	return year >= startYear && year <= endYear
}

func (passport *Passport) isBirthYearValid() bool {
	return isValidYear(passport.birthYear, 1920, 2002)
}

func (passport *Passport) isIssueYearValid() bool {
	return isValidYear(passport.issueYear, 2010, 2020)
}

func (passport *Passport) isExpirationYearValid() bool {
	return isValidYear(passport.expirationYear, 2020, 2030)
}

func (passport *Passport) isHeightValid() bool {
	height := passport.height
	nums, _ := regexp.Compile("[0-9]*")
	if strings.HasSuffix(height, "in") {
		heightInInches := parseStringToIntIgnoringError(nums.FindString(height))
		return heightInInches >= 59 && heightInInches <= 76
	}
	if strings.HasSuffix(height, "cm") {
		heightInCentimeters := parseStringToIntIgnoringError(nums.FindString(height))
		return heightInCentimeters >= 150 && heightInCentimeters <= 193
	}
	return false
}

func (passport *Passport) isHairColourValid() bool {
	hairColour := passport.hairColour
	if !strings.HasPrefix(hairColour, "#") {
		return false
	}
	colour, _ := regexp.Compile("[0-9a-fA-F]{6}")
	colourWithoutHash := string(hairColour[1:])
	parsed := colour.FindString(colourWithoutHash)
	if parsed == "" {
		return false
	}
	return true
}

func (passport *Passport) isEyeColourValid() bool {
	validEyeColours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, validColour := range validEyeColours {
		if validColour == passport.eyeColour {
			return true
		}
	}
	return false
}

func (passport *Passport) isPassportIdValid() bool {
	passportId := passport.passportId
	nums, _ := regexp.Compile("[0-9]*")
	return len(nums.FindString(passportId)) == 9
}

func (passport *Passport) isValidPerNewMethod() bool {
	return passport.isBirthYearValid() &&
		passport.isIssueYearValid() &&
		passport.isExpirationYearValid() &&
		passport.isHeightValid() &&
		passport.isHairColourValid() &&
		passport.isEyeColourValid() &&
		passport.isPassportIdValid()
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

func Day04Part2(input []string) int {
	passports := parseInput(input)
	validPassportCount := 0
	for _, passport := range passports {
		if passport.isValidPerNewMethod() {
			validPassportCount++
		}
	}
	return validPassportCount
}

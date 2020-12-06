package solutions

import (
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	firstNumber            int
	secondNumber           int
	characterThatMustOccur string
	password               string
}

func Day02Part1(input []string) int {
	return countAsPerPolicy(input, isValidAsPerOldPolicy)
}

func Day02Part2(input []string) int {
	return countAsPerPolicy(input, isValidAsPerNewPolicy)
}

func countAsPerPolicy(input []string, policy func(*PasswordPolicy) bool) int {
	var count int
	for _, line := range input {
		passwordPolicy := getPasswordPolicy(line)
		if policy(passwordPolicy) {
			count++
		}
	}
	return count
}

func isValidAsPerNewPolicy(passwordPolicy *PasswordPolicy) bool {
	characterAtFirstPosition := string(passwordPolicy.password[passwordPolicy.firstNumber-1])
	characterAtSecondPosition := string(passwordPolicy.password[passwordPolicy.secondNumber-1])
	return (characterAtFirstPosition == passwordPolicy.characterThatMustOccur && characterAtSecondPosition != passwordPolicy.characterThatMustOccur) || (characterAtFirstPosition != passwordPolicy.characterThatMustOccur && characterAtSecondPosition == passwordPolicy.characterThatMustOccur)
}

func isValidAsPerOldPolicy(passwordPolicy *PasswordPolicy) bool {
	occurrences := countChars(passwordPolicy.password, passwordPolicy.characterThatMustOccur)
	minOccurrences := passwordPolicy.firstNumber
	maxOccurrences := passwordPolicy.secondNumber
	return occurrences >= minOccurrences && occurrences <= maxOccurrences
}

func getPasswordPolicy(str string) *PasswordPolicy {
	splitBySpaces := strings.Split(str, " ")
	occurrencesPortion := splitBySpaces[0]
	splitByDash := strings.Split(occurrencesPortion, "-")
	minOccurrences, _ := strconv.Atoi(splitByDash[0])
	maxOccurrences, _ := strconv.Atoi(splitByDash[1])
	characterThatMustOccur := string(splitBySpaces[1][0])
	password := splitBySpaces[2]
	passwordPolicy := PasswordPolicy{
		firstNumber:            minOccurrences,
		secondNumber:           maxOccurrences,
		characterThatMustOccur: characterThatMustOccur,
		password:               password,
	}
	return &passwordPolicy
}

func countChars(str string, matcher string) int {
	var count int
	for _, s := range str {
		if string(s) == matcher {
			count++
		}
	}
	return count
}

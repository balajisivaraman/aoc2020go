package solutions

import (
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	minOccurrences         int
	maxOccurrences         int
	characterThatMustOccur string
	password               string
}

func Day02Part1(input []string) int {
	var count int
	for _, line := range input {
		passwordPolicy := getPasswordPolicy(line)
		if isValid(passwordPolicy) {
			count++
		}
	}
	return count
}

func isValid(passwordPolicy *PasswordPolicy) bool {
	occurrences := countChars(passwordPolicy.password, passwordPolicy.characterThatMustOccur)
	return occurrences >= passwordPolicy.minOccurrences && occurrences <= passwordPolicy.maxOccurrences
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
		minOccurrences:         minOccurrences,
		maxOccurrences:         maxOccurrences,
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

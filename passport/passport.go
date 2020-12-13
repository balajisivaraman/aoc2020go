package passport

import (
	"github.com/balajisivaraman/aoc2020go/utils"
	"regexp"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColour     string
	EyeColour      string
	PassportId     string
	CountryId      string
}

func isValidYear(strYear string, startYear int, endYear int) bool {
	if len(strYear) != 4 {
		return false
	}
	year := utils.ParseStringToIntIgnoringError(strYear)
	return year >= startYear && year <= endYear
}

func (passport *Passport) isBirthYearValid() bool {
	return isValidYear(passport.BirthYear, 1920, 2002)
}

func (passport *Passport) isIssueYearValid() bool {
	return isValidYear(passport.IssueYear, 2010, 2020)
}

func (passport *Passport) isExpirationYearValid() bool {
	return isValidYear(passport.ExpirationYear, 2020, 2030)
}

func (passport *Passport) isHeightValid() bool {
	height := passport.Height
	nums, _ := regexp.Compile("[0-9]*")
	if strings.HasSuffix(height, "in") {
		heightInInches := utils.ParseStringToIntIgnoringError(nums.FindString(height))
		return heightInInches >= 59 && heightInInches <= 76
	}
	if strings.HasSuffix(height, "cm") {
		heightInCentimeters := utils.ParseStringToIntIgnoringError(nums.FindString(height))
		return heightInCentimeters >= 150 && heightInCentimeters <= 193
	}
	return false
}

func (passport *Passport) isHairColourValid() bool {
	hairColour := passport.HairColour
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
		if validColour == passport.EyeColour {
			return true
		}
	}
	return false
}

func (passport *Passport) isPassportIdValid() bool {
	passportId := passport.PassportId
	nums, _ := regexp.Compile("[0-9]*")
	return len(nums.FindString(passportId)) == 9
}

func (passport *Passport) IsValidPerNewMethod() bool {
	return passport.isBirthYearValid() &&
		passport.isIssueYearValid() &&
		passport.isExpirationYearValid() &&
		passport.isHeightValid() &&
		passport.isHairColourValid() &&
		passport.isEyeColourValid() &&
		passport.isPassportIdValid()
}

func (passport *Passport) IsValidPerOldMethod() bool {
	return passport.BirthYear != "" &&
		passport.IssueYear != "" &&
		passport.ExpirationYear != "" &&
		passport.Height != "" &&
		passport.HairColour != "" &&
		passport.EyeColour != "" &&
		passport.PassportId != ""
}

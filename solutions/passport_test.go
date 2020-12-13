package solutions

import (
	"testing"
)

func TestBirthYearLessThanFourDigitsInvalid(t *testing.T) {
	input := Passport{
		birthYear:      "",
		issueYear:      "",
		expirationYear: "",
		height:         "",
		hairColour:     "",
		eyeColour:      "",
		passportId:     "",
		countryId:      "",
	}
	actual := input.isBirthYearValid()
	if actual {
		t.Errorf("Birth Year 200 is invalid")
	}
}

func TestBirthYearNotBetween1920And2002Invalid(t *testing.T) {
	input := Passport{
		birthYear: "2003",
	}
	actual := input.isBirthYearValid()
	if actual {
		t.Errorf("Birth Year 2003 is invalid")
	}
}

func TestBirthYearBetween1920And2002Valid(t *testing.T) {
	input := Passport{
		birthYear: "2002",
	}
	actual := input.isBirthYearValid()
	if !(actual) {
		t.Errorf("Birth Year 2002 is valid")
	}
}

func TestIssueYearLessThanFourDigitsInvalid(t *testing.T) {
	input := Passport{
		issueYear: "200",
	}
	actual := input.isIssueYearValid()
	if actual {
		t.Errorf("Issue Year 200 is invalid")
	}
}

func TestIssueYearNotBetween1920And2002Invalid(t *testing.T) {
	input := Passport{
		issueYear: "2009",
	}
	actual := input.isIssueYearValid()
	if actual {
		t.Errorf("Issue Year 2009 is invalid")
	}
}

func TestIssueYearBetween1920And2002Valid(t *testing.T) {
	input := Passport{
		issueYear: "2010",
	}
	actual := input.isIssueYearValid()
	if !(actual) {
		t.Errorf("Issue Year 2010 is valid")
	}
}

func TestExpirationYearLessThanFourDigitsInvalid(t *testing.T) {
	input := Passport{
		expirationYear: "200",
	}
	actual := input.isExpirationYearValid()
	if actual {
		t.Errorf("Expiration Year 200 is invalid")
	}
}

func TestExpirationYearNotBetween1920And2002Invalid(t *testing.T) {
	input := Passport{
		expirationYear: "2019",
	}
	actual := input.isExpirationYearValid()
	if actual {
		t.Errorf("Expiration Year 2019 is invalid")
	}
}

func TestExpirationYearBetween1920And2002Valid(t *testing.T) {
	input := Passport{
		expirationYear: "2020",
	}
	actual := input.isExpirationYearValid()
	if !(actual) {
		t.Errorf("Expiration Year 2021 is valid")
	}
}

func TestHeightWithoutCmOrInInvalid(t *testing.T) {
	input := Passport{
		height: "190",
	}
	actual := input.isHeightValid()
	if actual {
		t.Errorf("Height 190 is invalid")
	}
}

func TestHeightIncorrectCentimetersInvalid(t *testing.T) {
	input := Passport{
		height: "149cm",
	}
	actual := input.isHeightValid()
	if actual {
		t.Errorf("Height 149cm is invalid")
	}
}

func TestHeightCorrectCentimetersValid(t *testing.T) {
	input := Passport{
		height: "150cm",
	}
	actual := input.isHeightValid()
	if !(actual) {
		t.Errorf("Height 150cm is valid")
	}
}

func TestHeightIncorrectInchesInvalid(t *testing.T) {
	input := Passport{
		height: "190in",
	}
	actual := input.isHeightValid()
	if actual {
		t.Errorf("Height 190in is invalid")
	}
}

func TestHeightCorrectInchesValid(t *testing.T) {
	input := Passport{
		height: "59in",
	}
	actual := input.isHeightValid()
	if !(actual) {
		t.Errorf("Height 59in is valid")
	}
}

func TestHairColourWithoutHashInvalid(t *testing.T) {
	input := Passport{
		hairColour: "123abc",
	}
	actual := input.isHairColourValid()
	if actual {
		t.Errorf("Hair Colour 123abc is invalid")
	}
}

func TestHairColourWithHashIncorrectFormatInvalid(t *testing.T) {
	input := Passport{
		hairColour: "#123abz",
	}
	actual := input.isHairColourValid()
	if actual {
		t.Errorf("Hair Colour #123abz is invalid")
	}
}

func TestHairColourWithHashCorrectFormatValid(t *testing.T) {
	input := Passport{
		hairColour: "#123abc",
	}
	actual := input.isHairColourValid()
	if !(actual) {
		t.Errorf("Hair Colour #123abc is invalid")
	}
}

func TestEyeColourNotWithinSetInvalid(t *testing.T) {
	input := Passport{
		eyeColour: "amc",
	}
	actual := input.isEyeColourValid()
	if actual {
		t.Errorf("Hair Colour amc is invalid")
	}
}

func TestEyeColourWithinSetValid(t *testing.T) {
	input := Passport{
		eyeColour: "amb",
	}
	actual := input.isEyeColourValid()
	if !(actual) {
		t.Errorf("Hair Colour amb is valid")
	}
}

func TestPassportIdInvalid(t *testing.T) {
	input := Passport{
		passportId: "0123456789",
	}
	actual := input.isPassportIdValid()
	if actual {
		t.Errorf("Passport ID 0123456789 is invalid")
	}
}

func TestPassportIdValid(t *testing.T) {
	input := Passport{
		passportId: "000000001",
	}
	actual := input.isPassportIdValid()
	if !(actual) {
		t.Errorf("Passport ID 000000001 is valid")
	}
}

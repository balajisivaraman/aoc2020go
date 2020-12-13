package utils

import (
	"strconv"
)

func ParseStringToIntIgnoringError(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

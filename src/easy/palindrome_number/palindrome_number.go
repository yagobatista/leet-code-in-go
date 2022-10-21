package base

import (
	"math"
	"strconv"
)

func IsPalindrome(number int) bool {
	parsedNumber := strconv.Itoa(number)

	digitsCount := len(parsedNumber)

	middle := int(math.Floor(float64((digitsCount) / 2)))

	firstHalfIndex := middle
	secondHalfIndex := middle

	oddCount := digitsCount%2 == 1
	if oddCount {
		firstHalfIndex = middle
		secondHalfIndex = middle + 1
	}

	firstHalf := parsedNumber[:firstHalfIndex]
	secondHalf := parsedNumber[secondHalfIndex:]

	return firstHalf == secondHalf
}

package roman_to_integer

import (
	"fmt"
	"strings"
)

type Roman string

var (
	I Roman = "I"
	V Roman = "V"
	X Roman = "X"
	L Roman = "L"
	C Roman = "C"
	D Roman = "D"
	M Roman = "M"

	IV Roman = "IV"
	IX Roman = "IX"
	XL Roman = "XL"
	XC Roman = "XC"
	CD Roman = "CD"
	CM Roman = "CM"
)

var romanToIntegerMap = map[Roman]int{
	I: 1,
	V: 5,
	X: 10,
	L: 50,
	C: 100,
	D: 500,
	M: 1000,

	IV: 4,
	IX: 9,
	XL: 40,
	XC: 90,
	CD: 400,
	CM: 900,
}

type specialCasesSet map[Roman]bool

var specialCasesMap = map[Roman]specialCasesSet{
	I: {IV: true, IX: true},
	X: {XL: true, XC: true},
	C: {CD: true, CM: true},
}

func RomanToInteger(romanNumber string) int {
	digits := strings.Split(romanNumber, "")
	var result int

	for i := 0; i < len(digits); i++ {
		value, nextI := findValue(digits, i, Roman(digits[i]))
		i = nextI
		result += value
	}

	return result
}

func findValue(digits []string, index int, number Roman) (value int, nextDigit int) {
	value = romanToIntegerMap[number]
	nextDigit = index

	especialCases, isSpecialCase := specialCasesMap[number]
	if !isSpecialCase {
		return
	}

	if nextDigit+1 >= len(digits) {
		return
	}

	composeRomanDigit := Roman(fmt.Sprintf("%s%s", digits[index], digits[nextDigit+1]))

	_, foundCase := especialCases[composeRomanDigit]
	if !foundCase {
		return
	}

	return findValue(digits, nextDigit+1, composeRomanDigit)
}

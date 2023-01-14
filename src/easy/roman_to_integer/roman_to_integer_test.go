package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestTwoSum(t *testing.T) {
	suite.Run(t, new(RomanToIntegerSuite))
}

type RomanToIntegerSuite struct {
	suite.Suite
}

func (this *RomanToIntegerSuite) Test3() {
	result := RomanToInteger("III")

	this.Equal(3, result)
}

func (this *RomanToIntegerSuite) Test58() {
	result := RomanToInteger("LVIII")

	this.Equal(58, result)
}

func (this *RomanToIntegerSuite) Test1994() {
	result := RomanToInteger("MCMXCIV")

	this.Equal(1994, result)
}

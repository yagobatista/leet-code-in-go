package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestTwoSum(t *testing.T) {
	suite.Run(t, new(RomanToIntegerSuite))
}

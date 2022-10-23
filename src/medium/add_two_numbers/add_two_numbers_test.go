package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestAddTwoNumbers(t *testing.T) {
	suite.Run(t, new(addTwoNumbersSuite))
	suite.Run(t, new(sumSuite))
}

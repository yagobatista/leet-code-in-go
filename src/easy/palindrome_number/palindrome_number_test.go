package palindrome_number

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPalindromeNumber(t *testing.T) {
	suite.Run(t, new(PalindromeNumberSuite))
}

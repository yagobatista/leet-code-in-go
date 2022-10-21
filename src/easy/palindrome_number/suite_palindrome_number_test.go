package palindrome_number

import (
	"github.com/stretchr/testify/suite"
)

type PalindromeNumberSuite struct {
	suite.Suite
}

func (this *PalindromeNumberSuite) TestSimpleDigitNumber() {
	IsPalindrome := IsPalindrome(1)

	this.True(IsPalindrome)
}

func (this *PalindromeNumberSuite) TestTwoDigitInvalidNumber() {
	IsPalindrome := IsPalindrome(12)

	this.False(IsPalindrome)
}

func (this *PalindromeNumberSuite) TestTwoDigitValidNumber() {
	IsPalindrome := IsPalindrome(44)

	this.True(IsPalindrome)
}

func (this *PalindromeNumberSuite) TestLargeValidNumber() {
	IsPalindrome := IsPalindrome(44321344432134)

	this.True(IsPalindrome)
}

func (this *PalindromeNumberSuite) TestLargeNegativeNumber() {
	IsPalindrome := IsPalindrome(-44321344432134)

	this.False(IsPalindrome)
}

func (this *PalindromeNumberSuite) TestSimpleNumber() {
	IsPalindrome := IsPalindrome(121)

	this.True(IsPalindrome)
}

func (this *PalindromeNumberSuite) TestNegativeNumber() {
	IsPalindrome := IsPalindrome(-121)

	this.False(IsPalindrome)
}

func (this *PalindromeNumberSuite) TestInvalid() {
	IsPalindrome := IsPalindrome(10)

	this.False(IsPalindrome)
}

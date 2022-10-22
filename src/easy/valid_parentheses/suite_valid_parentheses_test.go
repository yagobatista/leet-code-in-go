package valid_parentheses

import (
	"github.com/stretchr/testify/suite"
)

type ValidParenthesesSuite struct {
	suite.Suite
}

func (this *ValidParenthesesSuite) TestSimple() {
	this.True(isValid("()"))
}

func (this *ValidParenthesesSuite) TestMultipleType() {
	this.True(isValid("()[]{}"))
}

func (this *ValidParenthesesSuite) TestComplex() {
	this.True(isValid("{{}}()[()]"))
	this.True(isValid("{{{}}()[{()}]}"))
}

func (this *ValidParenthesesSuite) TestInvalidCases() {
	this.False(isValid("(]"))
	this.False(isValid("{][}"))
	this.False(isValid("()("))
	this.False(isValid("]"))
	this.False(isValid("}{"))
	this.False(isValid(")(){}"))
}

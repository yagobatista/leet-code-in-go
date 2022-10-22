package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestValidParentheses(t *testing.T) {
	suite.Run(t, new(ValidParenthesesSuite))
}

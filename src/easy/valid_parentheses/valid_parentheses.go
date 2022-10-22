package valid_parentheses

import (
	"leet-code/src/pkg/stack"
	"strings"
)

type Bracket string

const (
	PARENTHESES           Bracket = "("
	CLOSING_PARENTHESES   Bracket = ")"
	CURLY_BRACKET         Bracket = "{"
	CLOSING_CURLY_BRACKET Bracket = "}"
	BRACKET               Bracket = "["
	CLOSING_BRACKET       Bracket = "]"
)

var bracketClosingMap = map[Bracket]Bracket{
	CLOSING_PARENTHESES:   PARENTHESES,
	CLOSING_CURLY_BRACKET: CURLY_BRACKET,
	CLOSING_BRACKET:       BRACKET,
}

func isValid(expression string) bool {
	var stack stack.Stack[Bracket]

	for _, character := range strings.Split(expression, "") {
		bracket := Bracket(character)

		correspondingOpenBracket, isClosing := bracketClosingMap[bracket]

		isOpen := !isClosing
		if isOpen {
			stack.Push(bracket)
			continue
		}

		if !stack.IsEmpty() && stack.Get() == correspondingOpenBracket {
			stack.Pop()
		} else {
			return false
		}
	}

	return stack.Len() == 0
}

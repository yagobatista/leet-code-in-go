package length_of_last_word

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestLengthOfLastWord(t *testing.T) {
	suite.Run(t, new(LengthOfLastWordSuite))
}

type LengthOfLastWordSuite struct {
	suite.Suite
}

func (this *LengthOfLastWordSuite) TestExample1() {
	length := lengthOfLastWord("Hello World")

	this.Equal(5, length)
}

func (this *LengthOfLastWordSuite) TestExample2() {
	length := lengthOfLastWord("   fly me   to   the moon  ")

	this.Equal(4, length)
}

func (this *LengthOfLastWordSuite) TestExample3() {
	length := lengthOfLastWord("luffy is still joyboy")

	this.Equal(6, length)
}

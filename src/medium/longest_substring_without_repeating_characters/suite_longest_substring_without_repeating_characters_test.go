package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	suite.Run(t, new(LongestSubstringWithoutRepeatingCharactersSuite))
}

type LongestSubstringWithoutRepeatingCharactersSuite struct {
	suite.Suite
}

func (this *LongestSubstringWithoutRepeatingCharactersSuite) TestExample1() {
	this.Equal(3, longestSubstringWithoutRepeatingCharacters("abcabcbb"))
}

func (this *LongestSubstringWithoutRepeatingCharactersSuite) TestExample2() {
	this.Equal(1, longestSubstringWithoutRepeatingCharacters("bbbbb"))
}

func (this *LongestSubstringWithoutRepeatingCharactersSuite) TestExample3() {
	this.Equal(3, longestSubstringWithoutRepeatingCharacters("pwwkew"))
}

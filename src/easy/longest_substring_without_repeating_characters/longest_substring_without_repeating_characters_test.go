package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	suite.Run(t, new(LongestSubstringWithoutRepeatingCharactersSuite))
}

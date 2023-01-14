package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestLongestCommonPrefix(t *testing.T) {
	suite.Run(t, new(LongestCommonPrefixSuite))
}

type LongestCommonPrefixSuite struct {
	suite.Suite
}

func (this *LongestCommonPrefixSuite) TestFindValid() {
	prefix := longestCommonPrefix([]string{"flower", "flow", "flight"})

	this.Equal("fl", prefix)
}

func (this *LongestCommonPrefixSuite) TestNoPrefix() {
	prefix := longestCommonPrefix([]string{"dog", "racecar", "car"})

	this.Equal("", prefix)
}

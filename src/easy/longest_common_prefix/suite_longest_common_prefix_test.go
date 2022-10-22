package longest_common_prefix

import (
	"github.com/stretchr/testify/suite"
)

type LongestCommonPrefixSuite struct {
	suite.Suite
}

func (this *LongestCommonPrefixSuite) TestFindValid() {
	prefix := LongestCommonPrefix([]string{"flower", "flow", "flight"})

	this.Equal("fl", prefix)
}

func (this *LongestCommonPrefixSuite) TestNoPrefix() {
	prefix := LongestCommonPrefix([]string{"dog", "racecar", "car"})

	this.Equal("", prefix)
}

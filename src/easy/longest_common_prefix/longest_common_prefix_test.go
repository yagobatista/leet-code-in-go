package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestLongestCommonPrefix(t *testing.T) {
	suite.Run(t, new(LongestCommonPrefixSuite))
}

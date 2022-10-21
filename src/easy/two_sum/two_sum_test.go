package two_sum

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestTwoSum(t *testing.T) {
	suite.Run(t, new(TwoSumSuite))
}

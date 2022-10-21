package two_sum

import (
	"github.com/stretchr/testify/suite"
)

type TwoSumSuite struct {
	suite.Suite
}

func (this *TwoSumSuite) TestInvalidInputs() {
	firstIndex, secondIndex, err := twoSum([]int{2, 7, 11, 15}, 1)

	this.ErrorContains(err, "invalid input or target")
	this.Equal(firstIndex, 0)
	this.Equal(secondIndex, 0)
}

func (this *TwoSumSuite) TestExample1() {
	firstIndex, secondIndex, err := twoSum([]int{2, 7, 11, 15}, 9)

	this.NoError(err)
	this.Equal(0, firstIndex)
	this.Equal(1, secondIndex)
}

func (this *TwoSumSuite) TestExample2() {
	firstIndex, secondIndex, err := twoSum([]int{3, 2, 4}, 6)

	this.NoError(err)
	this.Equal(1, firstIndex)
	this.Equal(2, secondIndex)
}

func (this *TwoSumSuite) TestExample3() {
	firstIndex, secondIndex, err := twoSum([]int{3, 3}, 6)

	this.NoError(err)
	this.Equal(0, firstIndex)
	this.Equal(1, secondIndex)
}

package remove_element

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestRemoveElement(t *testing.T) {
	suite.Run(t, new(RemoveElementSuite))
}

type RemoveElementSuite struct {
	suite.Suite
}

func (this *RemoveElementSuite) TestExample1() {
	output := removeElement([]int{3, 2, 2, 3}, 3)

	this.Equal(2, output)
}

func (this *RemoveElementSuite) TestExample2() {
	output := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)

	this.Equal(5, output)
}

func (this *RemoveElementSuite) TestNoReplaceableElements() {
	output := removeElement([]int{0, 1, 5, 4, 3, 10}, 15)

	this.Equal(6, output)
}

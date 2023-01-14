package rotate_image

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestGetAllElementsFromSquareSuite(t *testing.T) {
	suite.Run(t, new(RotateImageSuite))
	suite.Run(t, new(GetAllElementsFromSquareSuite))
}

type RotateImageSuite struct {
	suite.Suite
}

func (this *RotateImageSuite) TestExample1() {
	image := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateImage(image)
	this.Equal([][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, image)
}

func (this *RotateImageSuite) TestExample2() {
	image := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotateImage(image)
	this.Equal([][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}, image)
}

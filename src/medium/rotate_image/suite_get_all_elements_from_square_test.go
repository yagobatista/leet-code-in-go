package rotate_image

import (
	"github.com/stretchr/testify/suite"
)

type GetAllElementsFromSquareSuite struct {
	suite.Suite
}

func (this *GetAllElementsFromSquareSuite) TestGetAll4X4() {
	elements := getAllElementFromSquare([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, 1, 3)

	this.Equal([]int{4, 8, 6, 3}, elements)
}

func (this *GetAllElementsFromSquareSuite) TestGetAll3X3() {
	elements := getAllElementFromSquare([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 0, 3)

	this.Equal([]int{1, 2, 3, 6, 9, 8, 7, 4}, elements)
}

func (this *GetAllElementsFromSquareSuite) TestGetAll2X2() {
	elements := getAllElementFromSquare([][]int{{4, 8}, {3, 6}}, 0, 2)

	this.Equal([]int{4, 8, 6, 3}, elements)
}

func (this *GetAllElementsFromSquareSuite) TestRotateElements4X4() {
	rotated := rotateElements([]int{4, 8, 6, 3}, 1, 3)
	this.Equal([]int{3, 4, 8, 6}, rotated)
}

func (this *GetAllElementsFromSquareSuite) TestRotateElements3X3() {
	rotated := rotateElements([]int{1, 2, 3, 6, 9, 8, 7, 4}, 0, 3)
	this.Equal([]int{7, 4, 1, 2, 3, 6, 9, 8}, rotated)
}

func (this *GetAllElementsFromSquareSuite) TestRotateElements2X2() {
	rotated := rotateElements([]int{4, 8, 6, 3}, 0, 2)

	this.Equal([]int{3, 4, 8, 6}, rotated)
}

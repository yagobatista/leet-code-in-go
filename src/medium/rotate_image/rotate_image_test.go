package rotate_image

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestRotateImage(t *testing.T) {
	suite.Run(t, new(RotateImageSuite))
	suite.Run(t, new(GetAllElementsFromSquareSuite))
}

package set

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSet(t *testing.T) {
	suite.Run(t, new(SetSuite))
}

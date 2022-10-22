package same_tree

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSameTree(t *testing.T) {
	suite.Run(t, new(SameTreeSuite))
}

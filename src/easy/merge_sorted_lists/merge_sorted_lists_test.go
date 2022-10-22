package merge_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMergeSortedLists(t *testing.T) {
	suite.Run(t, new(MergeSortedListsSuite))
}

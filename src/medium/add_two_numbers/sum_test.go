package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSumSuite(t *testing.T) {
	suite.Run(t, new(sumSuite))
}

type sumSuite struct {
	suite.Suite
}

func (this *sumSuite) Test342() {
	this.Equal(342, sum(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}))
}

func (this *sumSuite) Test465() {
	this.Equal(465, sum(&ListNode{5, &ListNode{6, &ListNode{4, nil}}}))
}

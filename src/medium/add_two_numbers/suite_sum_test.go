package add_two_numbers

import (
	"github.com/stretchr/testify/suite"
)

type sumSuite struct {
	suite.Suite
}

func (this *sumSuite) Test342() {
	this.Equal(342, sum(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}))
}

func (this *sumSuite) Test465() {
	this.Equal(465, sum(&ListNode{5, &ListNode{6, &ListNode{4, nil}}}))
}

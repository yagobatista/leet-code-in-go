package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestAddTwoNumbers(t *testing.T) {
	suite.Run(t, new(addTwoNumbersSuite))
}

type addTwoNumbersSuite struct {
	suite.Suite
}

// Input: l1 = [24,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
func (this *addTwoNumbersSuite) Test() {
	expected := &ListNode{
		7,
		&ListNode{
			0,
			&ListNode{
				8,
				nil,
			},
		},
	}

	output := addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})

	this.Equal(expected, output)
}

func (this *addTwoNumbersSuite) TestEmpty() {
	output := addTwoNumbers(&ListNode{0, nil}, &ListNode{0, nil})

	this.Equal(&ListNode{0, nil}, output)
}

func (this *addTwoNumbersSuite) Test99() {
	n1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	n2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	expected := &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}

	output := addTwoNumbers(n1, n2)

	this.Equal(expected, output)
}

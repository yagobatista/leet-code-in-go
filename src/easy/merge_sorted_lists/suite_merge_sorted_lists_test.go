package merge_sorted_lists

import (
	"github.com/stretchr/testify/suite"
)

type MergeSortedListsSuite struct {
	suite.Suite
}

func (this *MergeSortedListsSuite) Test() {
	firstList := ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				4,
				nil,
			},
		},
	}
	secondList := ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}
	list := mergeTwoLists(&firstList, &secondList)

	expectedList := &ListNode{
		1,
		&ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					&ListNode{
						4,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
		},
	}

	this.Equal(expectedList, list)
}

func (this *MergeSortedListsSuite) TestEmpty() {
	var emptyList *ListNode
	list := mergeTwoLists(emptyList, emptyList)

	this.Equal(emptyList, list)
}

func (this *MergeSortedListsSuite) TestSingleElement() {
	var emptyList *ListNode

	list := mergeTwoLists(emptyList, &ListNode{0, nil})

	this.Equal(&ListNode{0, nil}, list)
}

func (this *MergeSortedListsSuite) TestSimpleLists() {
	list := mergeTwoLists(&ListNode{1, nil}, &ListNode{2, nil})

	this.Equal(&ListNode{1, &ListNode{2, nil}}, list)
}

func (this *MergeSortedListsSuite) TestSimpleLists2() {
	expectedList := &ListNode{
		-9,
		&ListNode{3,
			&ListNode{5,
				&ListNode{7, nil},
			},
		},
	}

	list := mergeTwoLists(&ListNode{-9, &ListNode{3, nil}}, &ListNode{5, &ListNode{7, nil}})

	this.Equal(expectedList, list)
}

package same_tree

import (
	"github.com/stretchr/testify/suite"
)

type SameTreeSuite struct {
	suite.Suite
}

func (this *SameTreeSuite) TestValid() {
	p := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}
	q := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	this.True(isSameTree(p, q))
}

func (this *SameTreeSuite) TestReverseTree() {
	p := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		nil,
	}
	q := &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			nil,
			nil,
		},
	}

	this.False(isSameTree(p, q))
}

func (this *SameTreeSuite) TestInvalid() {
	p := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			1,
			nil,
			nil,
		},
	}
	q := &TreeNode{
		1,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			1,
			nil,
			nil,
		},
	}

	this.False(isSameTree(p, q))
}

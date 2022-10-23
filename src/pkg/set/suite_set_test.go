package set

import (
	"github.com/stretchr/testify/suite"
)

type SetSuite struct {
	suite.Suite

	set Set[int]
}

func (this *SetSuite) SetupTest() {
	this.set = Set[int]{}
}

func (this *SetSuite) TestAdd() {
	this.set.Add(1)

	_, exists := this.set.data[1]
	this.True(exists)
}

func (this *SetSuite) TestAddMultipleItems() {
	this.set.Add(1, 2, 3, 4, 5)

	this.Equal(map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}, this.set.data)
}

func (this *SetSuite) TestAddSameItemMultipleTimes() {
	this.set.Add(1)
	this.set.Add(1)
	this.set.Add(1)
	this.set.Add(1)

	this.Equal(map[int]bool{1: true}, this.set.data)
}

func (this *SetSuite) TestLen() {
	this.set.data = map[int]bool{
		2: true,
		3: true,
		4: true,
	}

	this.Equal(3, this.set.Len())
}

func (this *SetSuite) TestClear() {
	this.set.data = map[int]bool{
		2: true,
		3: true,
		4: true,
		5: true,
	}

	this.set.Clear()

	this.Equal(0, this.set.Len())
}

func (this *SetSuite) TestIter() {
	this.set.data = map[int]bool{
		2: true,
		3: true,
		4: true,
		5: true,
	}

	this.Equal([]int{2, 3, 4, 5}, this.set.Iter())
}

func (this *SetSuite) TestMissingElement() {
	this.False(this.set.Exist(2))
}

func (this *SetSuite) TestFindExistingElement() {
	this.set.data = map[int]bool{
		2: true,
	}
	this.True(this.set.Exist(2))
}

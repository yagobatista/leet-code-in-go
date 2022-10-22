package stack

type IntStack struct {
	data []int
}

func (this *IntStack) Push(element int) {
	this.data = append(this.data, element)
}

func (this *IntStack) Pop() int {
	lastIndex := len(this.data) - 1
	element := this.data[lastIndex]

	this.data = this.data[:lastIndex]

	return element
}

func (this *IntStack) Get() int {
	return this.data[len(this.data)-1]
}

func (this *IntStack) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *IntStack) Len() int {
	return len(this.data)
}

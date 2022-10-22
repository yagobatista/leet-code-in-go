package stack

type Stack[T any] struct {
	data []T
}

func (this *Stack[T]) Push(element T) {
	this.data = append(this.data, element)
}

func (this *Stack[T]) Pop() T {
	lastIndex := len(this.data) - 1
	element := this.data[lastIndex]

	this.data = this.data[:lastIndex]

	return element
}

func (this *Stack[T]) Get() T {
	return this.data[len(this.data)-1]
}

func (this *Stack[T]) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *Stack[T]) Len() int {
	return len(this.data)
}

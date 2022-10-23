package set

type Set[T comparable] struct {
	data map[T]bool
}

func (this *Set[T]) Add(items ...T) {
	if this.data == nil {
		this.data = map[T]bool{}
	}

	for _, item := range items {
		this.data[item] = true
	}
}

func (this *Set[T]) Exist(key T) bool {
	_, found := this.data[key]
	return found
}

func (this *Set[T]) Clear() {
	this.data = map[T]bool{}
}

func (this *Set[T]) Len() int {
	return len(this.data)
}

func (this *Set[T]) Iter() []T {
	keys := make([]T, this.Len())
	index := 0

	for key := range this.data {
		keys[index] = key
		index++
	}

	return keys
}

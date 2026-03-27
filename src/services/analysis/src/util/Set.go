package util

// TODO mutexes for access

type Set[T comparable] map[T]bool

func (set Set[T]) Add(value T) {
	set[value] = true
}

func (set Set[T]) Delete(value T) {
	delete(set, value)
}

func (set Set[T]) Includes(value T) bool {
	return set[value]
}

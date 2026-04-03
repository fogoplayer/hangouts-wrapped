package util

import "container/heap"

type Heap[T any] struct {
	innerHeap[T]
}

func (h *Heap[T]) Push(x T) {
	heap.Push(&h.innerHeap, x)
}

func (h *Heap[T]) Pop() any {
	return heap.Pop(&h.innerHeap)
}

// ///////// //
// InnerHeap //
// ///////// //
type innerHeap[T any] struct {
	comparator func(T, T) int
	data       []T
}

func (h innerHeap[T]) Len() int           { return len(h.data) }
func (h innerHeap[T]) Less(i, j int) bool { return h.comparator(h.data[i], h.data[j]) < 0 }
func (h innerHeap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *innerHeap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.data = append(h.data, x.(T))
}

func (h *innerHeap[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

var _ heap.Interface = &innerHeap[any]{}

package util

import (
	"slices"
	"sort"
	"testing"
)

// The simplest possible Pop test - create a heap of ints, then empty it out and confirm that it's in order
func TestIntHeapPopsInOrder(t *testing.T) {
	heap := CreateHeap(func(a, b int) int { return a - b })

	values := []int{14, 11, 3, 6, 7, 18}
	ListForEach(values, func(value int) {
		heap.Push(value)
	})

	popped := make([]int, 0, len(values))
	for heap.Len() > 0 {
		v := heap.Pop()
		popped = append(popped, v)
	}

	sorted := sort.IntSlice(values)
	sorted.Sort()

	for i := range len(values) {
		if sorted[i] != popped[i] {
			t.Log("popped:", popped)
			t.Log("sorted", sorted)
			t.FailNow()
		}
	}
}

// Create a heap of structs that contain ints, then empty it out and confirm that it's in order
func TestStructHeapPopsInOrder(t *testing.T) {
	type TestStruct struct{ index, value int }
	sortTestStruct := func(a, b TestStruct) int {
		return a.value - b.value
	}

	heap := CreateHeap(sortTestStruct)

	baseValues := []int{14, 11, 3, 6, 7, 18}
	i := 0
	values := ListMap(baseValues, func(value int) TestStruct {
		i += 1
		bundle := TestStruct{i, value}
		heap.Push(bundle)
		return bundle
	})

	popped := make([]TestStruct, 0, len(values))
	for heap.Len() > 0 {
		v := heap.Pop()
		popped = append(popped, v)
	}

	sorted := values
	slices.SortFunc(sorted, sortTestStruct)

	for i := range len(values) {
		if sorted[i] != popped[i] {
			t.Log("popped:", popped)
			t.Log("sorted", sorted)
			t.Fail()
		}
	}
}

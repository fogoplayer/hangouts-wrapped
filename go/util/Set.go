package util

// TODO mutexes for access

type Set[T comparable] map[T]bool

func (set Set[T]) Add(values ...T) {
	for _, value := range values {
		set[value] = true
	}
}

func (set Set[T]) Delete(value T) {
	delete(set, value)
}

func (set Set[T]) Includes(value T) bool {
	return set[value]
}

// TODO this is pass-by-value, right?
func (set1 Set[T]) Disjoint(set2 Set[T]) Set[T] {
	disjoint := Set[T](MapMap(set1, func(k T, v bool) (T, bool) { return k, v }))

	// check all values of set2 that they are in set1
	for value := range set2 {
		if set1.Includes(value) {
			disjoint.Delete(value)
		} else {
			disjoint.Add(value)
		}
	}

	return disjoint
}

// Replaces the contents of the Set with values passed in
// Returns `true` if the overwrite added or removed the contents of the set, or `false` if not
func (set Set[T]) Overwrite(values ...T) bool { // TODO support passing in a set, too... maybe?
	newValuesSet := make(Set[T])
	newValuesSet.Add(values...)

	notSharedValues := set.Disjoint(newValuesSet)

	if len(notSharedValues) == 0 {
		return false
	}

	for notSharedValue := range notSharedValues {
		if set.Includes(notSharedValue) {
			set.Delete(notSharedValue)
		} else if newValuesSet.Includes(notSharedValue) {
			set.Delete(notSharedValue)
		}
	}

	return true
}

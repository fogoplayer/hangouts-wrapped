package reportoutputs

import (
	"slices"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"zarinloosli.com/hangouts-wrapped/util"
)

// Gets values and makes sure they are sorted
func TestValues(t *testing.T) {
	barOutput := CreateBarOutput()

	baseValues := []int{14, 11, 3, 6, 7, 18}
	i := 0
	util.ListForEach(baseValues, func(value int) {
		i += 1
		bundle := ReportOutputEntry[int]{strconv.Itoa(i), value}
		barOutput.Push(bundle)
	})

	sorted := sort.IntSlice(baseValues)
	sorted.Sort()
	slices.Reverse(sorted) // BarOutput uses increasing, not decreasing order

	fromHeap := barOutput.TypedValues()
	require.True(t, util.ListsAreEqual(fromHeap, sorted), "Lists do not match")
}

// Tests if the return value of the heap changes with multiple accesses
func TestValuesStability(t *testing.T) {
	barOutput := CreateBarOutput()

	baseValues := []int{14, 11, 3, 6, 7, 18}
	i := 0
	util.ListForEach(baseValues, func(value int) {
		i += 1
		bundle := ReportOutputEntry[int]{strconv.Itoa(i), value}
		barOutput.Push(bundle)
	})

	sorted := sort.IntSlice(baseValues)
	sorted.Sort()
	slices.Reverse(sorted) // BarOutput uses increasing, not decreasing order

	fromHeap := barOutput.TypedValues()
	require.True(t, util.ListsAreEqual(fromHeap, sorted), "Lists do not match")

	fromHeap = barOutput.TypedValues()
	require.True(t, util.ListsAreEqual(fromHeap, sorted), "Lists do not match")

	fromHeap = barOutput.TypedValues()
	require.True(t, util.ListsAreEqual(fromHeap, sorted), "Lists do not match")

	fromHeap = barOutput.TypedValues()
	require.True(t, util.ListsAreEqual(fromHeap, sorted), "Lists do not match")
}

// Gets values and makes sure they are sorted, even if other values in the struct are duplicated
func TestValuesWithDuplicateKeys(t *testing.T) {
	barOutput := CreateBarOutput()

	baseValues := []int{14, 11, 3, 6, 7, 18}
	i := 0
	util.ListForEach(baseValues, func(value int) {
		i += 1
		bundle := ReportOutputEntry[int]{"key", value}
		barOutput.Push(bundle)
	})

	sorted := sort.IntSlice(baseValues)
	sorted.Sort()
	slices.Reverse(sorted) // BarOutput uses increasing, not decreasing order

	fromHeap := barOutput.TypedValues()
	assert.True(t, util.ListsAreEqual(fromHeap, sorted), "Lists do not match")

	t.Log(fromHeap)
	t.Log(sorted)
}

// Tests if the return value of the heap changes with multiple accesses, even if other values in the struct are duplicated
func TestValuesStabilityWithDuplicateKeys(t *testing.T) {
	assert := assert.New(t)

	barOutput := CreateBarOutput()

	baseValues := []int{14, 11, 3, 6, 7, 18}
	i := 0
	util.ListForEach(baseValues, func(value int) {
		i += 1
		bundle := ReportOutputEntry[int]{"key", value}
		barOutput.Push(bundle)
	})

	sorted := sort.IntSlice(baseValues)
	sorted.Sort()
	slices.Reverse(sorted) // BarOutput uses increasing, not decreasing order

	fromHeap := barOutput.TypedValues()
	assert.True(util.ListsAreEqual(fromHeap, sorted), "Lists do not match")

	fromHeap = barOutput.TypedValues()
	assert.True(util.ListsAreEqual(fromHeap, sorted), "Lists do not match")
}

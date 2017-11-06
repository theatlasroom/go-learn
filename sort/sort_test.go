package algorithms

import (
	"reflect"
	"testing"

	sort "github.com/theatlasroom/go-learn/sort/algorithms"
)

var (
	unsorted = []int{10, 2, 0, 1, 4, 6, 7, 9, 5, 3, 8}
	sorted   = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func checkResult(testData, expected, actual []int, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Error(
			"\nFor", testData,
			"\nexpected", expected,
			"\ngot", actual,
		)
	}
}

func TestEmptyDataSets(t *testing.T) {
	var empty, result []int
	var _ error
	result, _ = sort.Sort(sort.Bubble, empty)
	checkResult(empty, empty, result, t)

	result, _ = sort.Sort(sort.Selection, empty)
	checkResult(empty, empty, result, t)

	result, _ = sort.Sort(sort.Insertion, empty)
	checkResult(empty, empty, result, t)
}

func TestBubble(t *testing.T) {
	// data :=
	data := make([]int, len(unsorted))
	copy(data, unsorted)
	result, _ := sort.Sort(sort.Bubble, data)

	checkResult(unsorted, sorted, result, t)
}

func TestSelection(t *testing.T) {
	// data :=
	data := make([]int, len(unsorted))
	copy(data, unsorted)
	result, _ := sort.Sort(sort.Selection, data)

	checkResult(unsorted, sorted, result, t)
}

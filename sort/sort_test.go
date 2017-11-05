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

func TestBubble(t *testing.T) {
	data := unsorted[:]
	result, _ := sort.Sort(data, sort.Bubble)

	if !reflect.DeepEqual(result, sorted) {
		t.Error(
			"\nFor", unsorted,
			"\nexpected", sorted,
			"\ngot", result,
		)
	}
	// fmt.Println(arr)
}

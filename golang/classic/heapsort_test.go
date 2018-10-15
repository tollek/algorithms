package classic

import (
	"sort"
	"testing"
)

func TestHeapsort1(t *testing.T) {
	a := sort.IntSlice([]int{2, 8, 7, 1, 3, 5, 6, 4})
	Heapsort(a)
	if !sort.IsSorted(a) {
		t.Error(a)
	}
}

func TestHeapsort2(t *testing.T) {
	a := sort.IntSlice([]int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11})
	Heapsort(a)
	if !sort.IsSorted(a) {
		t.Error(a)
	}
}

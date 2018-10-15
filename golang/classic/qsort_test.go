package classic

import "testing"

func TestQsort1(t *testing.T) {
	a := []int{2, 8, 7, 1, 3, 5, 6, 4}
	Qsort(a)
	if !IsSorted(a) {
		t.Error(a)
	}
}

func TestQsort2(t *testing.T) {
	a := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	Qsort(a)
	if !IsSorted(a) {
		t.Error(a)
	}
}

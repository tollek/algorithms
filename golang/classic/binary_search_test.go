package classic

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		a        []int
		el       int
		expected int
	}{
		// 5 elements
		{a: []int{1, 2, 3, 4, 5}, el: 0, expected: -1},
		{a: []int{1, 2, 3, 4, 5}, el: 1, expected: 0},
		{a: []int{1, 2, 3, 4, 5}, el: 2, expected: 1},
		{a: []int{1, 2, 3, 4, 5}, el: 3, expected: 2},
		{a: []int{1, 2, 3, 4, 5}, el: 4, expected: 3},
		{a: []int{1, 2, 3, 4, 5}, el: 5, expected: 4},
		{a: []int{1, 2, 3, 4, 5}, el: 6, expected: -1},

		// 6 elements
		{a: []int{1, 2, 3, 4, 5, 6}, el: 0, expected: -1},
		{a: []int{1, 2, 3, 4, 5, 6}, el: 1, expected: 0},
		{a: []int{1, 2, 3, 4, 5, 6}, el: 2, expected: 1},
		{a: []int{1, 2, 3, 4, 5, 6}, el: 3, expected: 2},
		{a: []int{1, 2, 3, 4, 5, 6}, el: 4, expected: 3},
		{a: []int{1, 2, 3, 4, 5, 6}, el: 5, expected: 4},
		{a: []int{1, 2, 3, 4, 5, 6}, el: 6, expected: 5},
		{a: []int{1, 2, 3, 4, 5, 6}, el: 7, expected: -1},

		// http://codekata.com/kata/kata02-karate-chop/
		{a: []int{}, el: 3, expected: -1},
		{a: []int{1}, el: 3, expected: -1},
		{a: []int{1}, el: 1, expected: 0},
		//
		{a: []int{1, 3, 5}, el: 1, expected: 0},
		{a: []int{1, 3, 5}, el: 3, expected: 1},
		{a: []int{1, 3, 5}, el: 5, expected: 2},
		{a: []int{1, 3, 5}, el: 0, expected: -1},
		{a: []int{1, 3, 5}, el: 2, expected: -1},
		{a: []int{1, 3, 5}, el: 4, expected: -1},
		{a: []int{1, 3, 5}, el: 6, expected: -1},
		//
		{a: []int{1, 3, 5, 7}, el: 1, expected: 0},
		{a: []int{1, 3, 5, 7}, el: 3, expected: 1},
		{a: []int{1, 3, 5, 7}, el: 5, expected: 2},
		{a: []int{1, 3, 5, 7}, el: 7, expected: 3},
		{a: []int{1, 3, 5, 7}, el: 0, expected: -1},
		{a: []int{1, 3, 5, 7}, el: 2, expected: -1},
		{a: []int{1, 3, 5, 7}, el: 4, expected: -1},
		{a: []int{1, 3, 5, 7}, el: 6, expected: -1},
		{a: []int{1, 3, 5, 7}, el: 8, expected: -1},
	}

	for _, c := range cases {
		index := BinarySearch(c.a, c.el)
		if index != c.expected {
			t.Errorf("%d %v => %d (should be %d)", c.el, c.a, index, c.expected)
		}
	}
}

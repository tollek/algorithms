package classic

// algortihm explanation:
// https://www.nayuki.io/page/next-lexicographical-permutation-algorithm

import (
	"sort"
)

// https: //leetcode.com/problems/next-permutation/description/

// creates next (lexicographically) permutation in place.
// If this is already last permutation, return false.
func NextPermutation(data sort.Interface) bool {
	// find longest non-increasing suffix
	i := longestNonIncreasingSuffix(data)
	if i == -1 {
		return false
	}
	// find element within the suffix to be swapped with pre-suffix element
	j := findSwapElement(data, i)
	data.Swap(i-1, j)

	// reverse the suffix, make it increasing
	reverse(data, i)

	return true
}

// return begining of longest, non-increasing suffix.
// or -1 if not exists
func longestNonIncreasingSuffix(data sort.Interface) int {
	n := data.Len()
	for i := n - 1; i-1 >= 0; i-- {
		if data.Less(i-1, i) {
			return i
		}
	}
	return -1
}

func findSwapElement(data sort.Interface, suffixStart int) int {
	swapIndex := suffixStart - 1
	for j := data.Len() - 1; j >= suffixStart; j-- {
		if data.Less(swapIndex, j) {
			return j
		}
	}
	panic("should not happen")
}

func reverse(data sort.Interface, suffixStart int) {
	n := data.Len()
	// to reverse data[suffixStart..n] we need
	// this many swaps
	swaps := (n - suffixStart + 1) / 2

	for i := 0; i < swaps; i++ {
		data.Swap(suffixStart+i, n-1-i)
	}
}

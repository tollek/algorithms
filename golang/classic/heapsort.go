package classic

import "sort"

type Heap struct {
	sort.Interface
}

// sort data in place using heapsort algorithm
func Heapsort(data sort.Interface) {
	heap := BuildHeap(data)
	heap.Sort()
}

func BuildHeap(data sort.Interface) Heap {
	return Heap{data}
}

func (h Heap) Sort() {
	n := h.Len()
	for i := n / 2; i >= 0; i-- {
		h.heapify(i, n)
	}

	for i := 0; i < n-1; i++ {
		h.Swap(0, n-1-i)
		h.heapify(0, n-1-i)
	}
}

// heapify will fix the heap subtree rooted at position root
// (assuming n - size of the whole heap).
// Passing n as size explicitly allows manipulating the heapify operation range
// (required during heap sort).
func (h Heap) heapify(root int, n int) {
	largest := root
	left := root*2 + 1  // 0-based indexing
	right := root*2 + 2 // 0-based indexing

	if left < n && h.Less(root, left) {
		largest = left
	}
	if right < n && h.Less(largest, right) {
		largest = right
	}

	if largest != root {
		h.Swap(largest, root)
		// recursively heaplify the subtree that has changed
		h.heapify(largest, n)
	}
}

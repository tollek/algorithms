package classic

const binarySearchMaxRecusionDepth int = 30

func BinarySearch(a []int, el int) int {
	return binarySearchImpl(a, 0, len(a), el, binarySearchMaxRecusionDepth)
}

// performs binary search in a in [start, end) range
func binarySearchImpl(a []int, start, end, el int, recursionDepthLeft int) int {
	if recursionDepthLeft == 0 {
		panic("recusion too deep")
	}
	// fmt.Printf("\t%v (%d %d) %d\n", a[start:end], start, end, el)

	switch end - start {
	case 0:
		// this should never happen
		return -1
	case 1:
		if a[start] == el {
			return start
		} else {
			return -1
		}
	default:
		// the core problem of binary search: how to split the array
		// so that it's always smaller than it was (no recursive loop)
		//
		// the idea about the solution is to split the thinking about the problem
		// into 2 separate sub-problems:
		// - how to split array into 2 smaller ones
		// - how to pick the proper array for later
		//
		// After the first problem is solved properly, second is simple "if" on indices.
		// The idea behind the split:
		// - 1st sub-array must be equal or larger by one than the 2nd:
		//   [1 2 3 4] => [1 2] + [3 4]
		//   [1 2 3]   => [1 2] + [3]
		//   [1 2]     => [1] + [2]
		mid := (end + start + 1) / 2
		if el < a[mid] {
			return binarySearchImpl(a, start, mid, el, recursionDepthLeft-1)
		} else {
			return binarySearchImpl(a, mid, end, el, recursionDepthLeft-1)
		}
	}

}

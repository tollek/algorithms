package classic

import (
	"sort"
)

func IsSorted(a []int) bool {
	x := sort.IntSlice(a)
	return sort.IsSorted(x)
}

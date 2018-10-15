package classic

func IsSorted(a []int) bool {
	if len(a) == 0 {
		return true
	}
	last := a[0]
	for _, next := range a {
		if next < last {
			return false
		}
		last = next
	}
	return true
}

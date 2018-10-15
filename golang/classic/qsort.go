package classic

import(
  "math/rand"
)

// Qsort sorts a in place
func Qsort(a []int) {
    qsortImpl(a, 0, len(a))
}

// sorts a in place in [start, end) range
func qsortImpl(a []int, start, end int) {
  if end - start > 1 {
    q := randomPartition(a, start, end)
    qsortImpl(a, start, q)
    qsortImpl(a, q+1, end)
  }
}

// partition will rearrange the subarray a[start, end) in place
//
// algorithm splits the array into spaces:
// - if start <= k <= i, then a[k] <= pivot (elemnts that are smaller than pivot, and "already" handled)
// - if i+1 <= k <= j-1, then a[k] > pivot, but hasn't been moved to proper place yet
// - k == end-1, then a[k] = pivot elements
// - if j <= k < end-1, then a[k] is not processed yet
func partition(a []int, start, end int) int {
  pivot := a[end-1]
  i := start - 1
  for j := start; j < end-1; j++ {
    if a[j] < pivot {
      i++
      a[i], a[j] = a[j], a[i]
    }
  }
  a[i+1], a[end-1] = a[end-1], a[i+1]
  return i+1
}

// same as partition(..), but randomizes the choice of pivot element
func randomPartition(a []int, start, end int) int {
  i := rand.Intn(end-start)+start
  a[i], a[end-1] = a[end-1], a[i]
  return partition(a, start, end)
}

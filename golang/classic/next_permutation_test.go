package classic

import (
	"reflect"
	"sort"
	"testing"
)

func TestNextPerm1(t *testing.T) {
	a := []int{0, 1, 2, 5, 3, 3, 0}
	ok := NextPermutation(sort.IntSlice(a))
	if !ok {
		t.Fatal("did not create next permutation!")
	}

	expected := []int{0, 1, 3, 0, 2, 3, 5}
	if !reflect.DeepEqual(expected, a) {
		t.Log(a)
		t.FailNow()
	}
}

func TestNextPermNoPermutation(t *testing.T) {
	a := []int{2, 1, 0}
	ok := NextPermutation(sort.IntSlice(a))
	if ok {
		t.Log(a)
		t.Fatal("should not create more permutations, but did")
	}
}

func TestAllPermutations(t *testing.T) {
	a := []int{1, 2, 3, 4}
	perms := 1
	for NextPermutation(sort.IntSlice(a)) {
		t.Log(a)
		perms++
	}

	if perms != 24 {
		t.Errorf("did not generate 24 permutations, got only %d", perms)
	}
}

func TestEmptyArrayPermutation(t *testing.T) {
	a := []int{}
	ok := NextPermutation(sort.IntSlice(a))
	if ok {
		t.Log(a)
		t.Fatal("should not create more permutations, but did")
	}
}

func TestOneElementArrayPermutation(t *testing.T) {
	a := []int{123}
	ok := NextPermutation(sort.IntSlice(a))
	if ok {
		t.Log(a)
		t.Fatal("should not create more permutations, but did")
	}
}

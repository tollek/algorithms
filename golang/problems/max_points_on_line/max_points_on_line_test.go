package maxpoints

import (
	"testing"
)

func TestMaxPoints1(test *testing.T) {
	//	return
	p := []Point{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	ans := maxPoints(p)
	if ans != 3 {
		test.Errorf("%v\ngot: %d\nexpected: %d\n", p, ans, 3)
	}
}

func TestMaxPoints2(test *testing.T) {
	//return
	p := []Point{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}
	ans := maxPoints(p)
	if ans != 4 {
		test.Errorf("%v\ngot: %d\nexpected: %d\n", p, ans, 4)
	}
}

func TestMaxPoints3(test *testing.T) {
	//	return
	p := []Point{
		{1, 1},
		{1, 1},
		{2, 2},
		{2, 2},
	}
	ans := maxPoints(p)
	if ans != 4 {
		test.Errorf("%v\ngot: %d\nexpected: %d\n", p, ans, 4)
	}
}

func TestMaxPoints4(test *testing.T) {
	p := []Point{
		{0, 0},
		{0, 0},
	}
	ans := maxPoints(p)
	if ans != 2 {
		test.Errorf("%v\ngot: %d\nexpected: %d\n", p, ans, 2)
	}
}

func TestMaxPoints5(test *testing.T) {
	p := []Point{
		{0, 0},
	}
	ans := maxPoints(p)
	if ans != 1 {
		test.Errorf("%v\ngot: %d\nexpected: %d\n", p, ans, 1)
	}
}

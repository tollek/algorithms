package maxpoints

// Solution
// for each point:
//    make it center of coordinate system
//    sort all other points using cosinus-like metric (a bit modified to avoid 0 traps)
//    calculate points that have the same metric value
//      (if points are in different quarters, that's OK; we'll calculate them together when
//       coordinate center is moved to boundary point)

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/max-points-on-a-line/
type Point struct {
	X int
	Y int
}

func debugf(s string, args ...interface{}) {
	return
	fmt.Printf(s, args...)
}

func debugln(i interface{}) {
	return
	fmt.Println(i)
}

func maxPoints(points []Point) int {
	n := len(points)
	best := 0
	for i := 0; i < n; i++ {
		center := points[i]
		makeCenter(points, center)
		debugln("---------------------------")
		debugln(center)
		debugln(points)

		cand := calculate(points, i)
		if cand > best {
			best = cand
		}

		makeCenter(points, reversePoint(center))
		debugln("")
	}

	return best
}

func makeCenter(points []Point, center Point) {
	for i := range points {
		points[i].X -= center.X
		points[i].Y -= center.Y
	}
}

func reversePoint(p Point) Point {
	return Point{-p.X, -p.Y}
}

type pointMetric struct {
	u, d int
}

func less(a, b pointMetric) bool {
	p := int64(a.u) * int64(b.d)
	q := int64(b.u) * int64(a.d)
	if p == q {
		return a.u < b.u
	}
	return p < q
}

func equal(a, b pointMetric) bool {
	p := int64(a.u) * int64(b.d)
	q := int64(b.u) * int64(a.d)
	return p == q
}

type pointMetricList []pointMetric

func (p pointMetricList) Len() int           { return len(p) }
func (p pointMetricList) Less(i, j int) bool { return less(p[i], p[j]) }
func (p pointMetricList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func calculate(points []Point, centerIdx int) int {
	n := len(points)
	metrics := make(pointMetricList, n)
	for i := 0; i < n; i++ {
		metrics[i] = calcMetric(points[i])
	}
	sort.Sort(metrics)
	debugln(metrics)

	// calculate "0,0" points size
	zeroPoints := 0
	for i := 0; i < n; i++ {
		if metrics[i].u == 0 && metrics[i].d == 1 {
			zeroPoints++
		} else {
			break
		}
	}
	debugf("zeroPoints = %d\n", zeroPoints)

	// calculate biggest group
	start := 0
	size := 0
	maxSize := 0
	for i := zeroPoints; i < n; i++ {
		if equal(metrics[i], metrics[start]) {
			size++
		} else {
			if size > maxSize {
				debugf("group %d, size = %d\n", start, size)
				maxSize = size
			}
			start = i
			size = 1
		}
	}
	if size > maxSize {
		debugf("group %d, size = %d\n", start, size)
		maxSize = size
	}
	return maxSize + zeroPoints
}

func calcMetric(p Point) pointMetric {
	x, y := p.X, p.Y
	switch {
	case x == 0 && y == 0:
		return pointMetric{0, 1}
	case x > 0 && y >= 0:
		return pointMetric{y, x + y}
	case x > 0 && y < 0 || x == 0 && y < 0:
		return pointMetric{x - 2*y, x - y}
	case x < 0 && y <= 0:
		return pointMetric{-3*y - 2*x, -y - x}
	case x <= 0 && y > 0:
		return pointMetric{-3*x + 4*y, -x + y}
	}
	panic(fmt.Sprintf("should not happen %d %d", x, y))
}

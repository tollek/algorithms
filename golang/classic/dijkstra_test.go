package classic

import (
	"container/list"
	"reflect"
	"testing"
)

// example graph for Dijkstra algorithm from Cormen (3ed edition)
func TestDijkstra(t *testing.T) {
	g := Graph{
		V: 5,
		E: make([]list.List, 5),
	}
	g.AddEdges([]Edge{
		{0, 1, 10},
		{0, 3, 5},
		{1, 2, 1},
		{1, 3, 2},
		{2, 4, 4},
		{3, 1, 3},
		{3, 2, 9},
		{3, 4, 2},
		{4, 0, 7},
		{4, 2, 6},
	}...)
	t.Log(g.ToString())

	dist := Dijkstra(&g, 0)
	t.Log(dist)

	expected := []GraphDistance{
		{-1, 0},
		{3, 8},
		{1, 9},
		{0, 5},
		{3, 7},
	}
	if !reflect.DeepEqual(expected, dist) {
		t.FailNow()
	}
}

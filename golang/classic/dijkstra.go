package classic

import (
	"container/heap"
	"math"
)

func Dijkstra(g *Graph, source int) []GraphDistance {
	dist := initializeSingleSource(g, source)
	// initialize the queue
	q := &sourceDistList{}
	heap.Push(q, sourceDist{V: source, Distance: 0})

	for q.Len() > 0 {
		u := heap.Pop(q).(sourceDist)
		// fmt.Println(u)
		// extra check, as we're not updating the queue - just adding and popping
		// elements which might no longer be the best possible path
		//
		// extra check is only a performance improvement - no relaxation would happen
		// from "old" distance, but we'd traverse all edges going from u.V
		if u.Distance <= dist[u.V].Distance {
			for e := g.EdgesFrom(u.V); e != nil; e = e.Next() {
				edge := e.Value.(Edge)
				// relax the u.V ---> edge.To edge
				v := edge.To
				newDistance := u.Distance + edge.Weight
				if dist[v].Distance > newDistance {
					// fmt.Printf("\t%d -> %d relax %d (prev %d)\n", u.V, v, newDistance, dist[v].Distance)
					dist[v].Prev = u.V
					dist[v].Distance = newDistance
					heap.Push(q, sourceDist{V: v, Distance: newDistance})
				} else {
					// fmt.Printf("\t%d -> %d no relax %d (prev %d)\n", u.V, v, newDistance, dist[v].Distance)
				}
			}
		}
		// fmt.Println("\tq: ", q)
	}
	return dist
}

type sourceDist struct {
	V        int
	Distance int
}

// An sourceDistList is a min-heap of distances
type sourceDistList []sourceDist

func (h sourceDistList) Len() int           { return len(h) }
func (h sourceDistList) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h sourceDistList) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *sourceDistList) Push(x interface{}) {
	*h = append(*h, x.(sourceDist))
}

func (h *sourceDistList) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func initializeSingleSource(g *Graph, source int) []GraphDistance {
	dist := make([]GraphDistance, g.V)
	for i := range dist {
		dist[i].Prev = -1
		dist[i].Distance = math.MaxInt32
	}
	dist[source].Prev = -1
	dist[source].Distance = 0
	return dist
}

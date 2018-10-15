package classic

import (
	"bytes"
	"container/list"
	"fmt"
)

type Graph struct {
	V int
	// lists of Edge objects
	E []list.List
}

type Edge struct {
	From, To int
	Weight   int
}

type GraphDistance struct {
	Prev     int
	Distance int
}

func (g *Graph) AddEdges(edges ...Edge) {
	for _, e := range edges {
		g.E[e.From].PushBack(e)
	}
}

func (g *Graph) EdgesFrom(node int) *list.Element {
	return g.E[node].Front()
}

func (g *Graph) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d\n", g.V))
	for i := 0; i < g.V; i++ {
		buffer.WriteString(fmt.Sprintf("%d:\t", i))
		for e := g.E[i].Front(); e != nil; e = e.Next() {
			edge := e.Value.(Edge)
			buffer.WriteString(fmt.Sprintf("%d (%d) ", edge.To, edge.Weight))
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

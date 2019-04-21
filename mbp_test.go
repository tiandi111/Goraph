package goraph

import (
	"fmt"
	"testing"
)

func TestMBP(t *testing.T) {
	g := NewUndirectedGraph(5, 3, -1, 10, true)
	fmt.Println( g.AdjList )
	fmt.Println( getMBP(g, 0, 0) )
}

// A trivial but reliable test algorithm is DFS
// Simply traverse all possible paths and take the maximum bandwidth
func getMBP(g *Graph, s int, t int) float64{
	// This algorithm only cares about connected graph
	if !g.Connected {
		fmt.Println("Graph should be connected!")
		return -1 
	}
	visited := make([]bool, len(g.AdjList))
	var max float64 = 0
	getMBPHelper(g.AdjList, s, t, s, visited, 10, &max)
	return max
}

func getMBPHelper(l [][]Edge, s int, t int, cur int, visited []bool, bw float64, max *float64) {
	if visited[cur] { return }
	if cur == t {
		if bw > *max {
			*max = bw
		}
		return
	}

	visited[cur] = true

	for _, edge := range l[cur] {
		if edge.weight < bw {
			bw = edge.weight
		}
		getMBPHelper(l, s, t, edge.tail, visited, bw, max)
	}
}



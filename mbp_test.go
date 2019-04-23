package goraph

import (
	"fmt"
	"testing"
	"math/rand"
)

func TestMBP(t *testing.T) {
	testcases := []struct {
		Size 		int
		AvgENum 	int
		Density		float64
		MaxWeight	float64
		Connected	bool
	}{
		{9, 4, 0.6, 10, true},
		{10, -1, 0.6, 100, true},
		{200, -1, 0.2, 1, true},
		{5000, 6, -1.0, 2500, true},
		//{5000, -1, 0.2, 64, true},
	}
	for _, tc := range testcases {
		descr := fmt.Sprintf("Graph(size: %d, average edges: %d, density: %f, maxweight: %f, connected: %t)", tc.Size, tc.AvgENum, tc.Density, tc.MaxWeight, tc.Connected)
		// get arributes
		g := NewUndirectedGraph(tc.Size, tc.AvgENum, tc.Density, tc.MaxWeight, tc.Connected)
		sc := rand.Intn(g.Size)
		tm := rand.Intn(g.Size)
		//exp := getMBP(g, sc, tm)
		exp := MBP(g, sc, tm, "Kruskal")
		r  := MBP(g, sc, tm, "Dijkstra")
		if exp != r {
			t.Errorf("\nGraph: %q\n    Source:%d\n    Terminal:%d\n    Expect %f, but get %f", descr, sc, tm, exp, r)
		}
	}
}

//func MBP(g *Graph, s int, t int) float64 {
//	return 0
//}

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
	getMBPHelper(g.AdjList, s, t, s, visited, g.MaxWeight, &max)
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
			getMBPHelper(l ,s, t, edge.tail, visited, edge.weight, max)
		} else {
			getMBPHelper(l, s, t, edge.tail, visited, bw, max)
		}
	}
}



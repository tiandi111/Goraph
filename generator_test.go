package main

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	testcases := []struct {
		Size 		int
		AvgENum 	int
		Density		float64
		MaxWeight	float64
		Connected	bool
	}{
		{10, 2, -1.0, 10, true},
		{10, -1, 0.6, 100, true},
		{10, 2, -1.0, 1, false},
		{5000, 6, -1.0, 25, true},
		{5000, -1, 0.2, 64, true},
	}
	for _, tc := range testcases {
		descr := fmt.Sprintf("Graph(size: %d, average edges: %d, density: %f, maxweight: %f, connected: %t)", tc.Size, tc.AvgENum, tc.Density, tc.MaxWeight, tc.Connected)
		// get arributes
		g := NewUndirectedGraph(tc.Size, tc.AvgENum, tc.Density, tc.MaxWeight, tc.Connected)
		rSize := len(g.AdjList)
		rTotalENum := 0
		for _, vList := range g.AdjList {rTotalENum += len(vList)} 
		rConnected := DFS(g)
		// test size
		if rSize != tc.Size {
			t.Errorf("%q has %d vertices, expect %d", descr, rSize, tc.Size)
		}
		// test average number of edges 
		if tc.AvgENum >= 0 {
			rAvgENum := rTotalENum/rSize
			if tc.AvgENum != rAvgENum {
				t.Errorf("%q has %d edges per vertex, expect %d", descr, rAvgENum, tc.AvgENum)
			}
		}
		// test density
		if tc.Density >= 0 {
			rDensity := float64(rTotalENum)/float64(rSize*(rSize-1))
			if rDensity != tc.Density {
				t.Errorf("%q has %f edges, expect %f", descr, rDensity, tc.Density)
			}
		}
		// test connected
		if rConnected == false && tc.Connected == true {
			t.Errorf("%q isn't connected, expect connected", descr)
		}
	}
}

func DFS(g *Graph) bool {
	visited := make([]bool, len(g.AdjList))
	// Traverse the graph, if the graph is connected, we should have seen all vertices 
	// otherwise, it is un-connected
	DFShelper(g.AdjList, 0, visited)
	for _, seen := range visited {
		if !seen {
			return false
		}
	} 
	return true
}

func DFShelper(l [][]Edge, vertex int, visited []bool) {
	if visited[vertex] {return}
	visited[vertex] = true
	for _, next := range l[vertex] {
		DFShelper(l, next.tail, visited)
	}
}


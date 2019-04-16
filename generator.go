package main

import (
	"fmt"
//	"math/rand"
)

type Graph struct {
	AdjList		[][]int
	Size		int
	AvgENum		int
	Density		float64
	Connected	bool
}

func main() {
	g := NewUndirectedGraph(10, 2, -1, true)
	fmt.Print(g)
}

func NewUndirectedGraph(size int, avgenum int, density float64, connected bool) *Graph {
	g := Graph{make([][]int, size), size, avgenum, density, connected}
	for i := range g.AdjList {
		g.AdjList[i] = make([]int, 0)
	}
	if connected {
		ConnectGraph(&g)
	}
	return &g
}

func ConnectGraph(g *Graph) {
	l := g.AdjList
	unconnected := make(map[int]bool, len(l))
	for i := 0; i<len(l); i++ {unconnected[i] = true}
	prev := -1
	head := -1
	for k, _ := range unconnected {
		if prev >= 0 {
			g.add(prev, k)
		} else {
			head = k
		}
		prev = k
	}
	g.add(prev, head)
}

func (g *Graph) add(a int, b int) {
	g.AdjList[a] = append(g.AdjList[a], b)
	g.AdjList[b] = append(g.AdjList[b], a) 
}

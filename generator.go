package goraph

import (
	"fmt"
	"math/rand"
)

type Edge struct {
	tail		int
	weight		float64
}

type Graph struct {
	AdjList		[][]Edge
	Size		int
	AvgENum		int
	Density		float64
	MaxWeight	float64
	Connected	bool
}

func main() {
	g := NewUndirectedGraph(10, 2, 0.6, 10, true)
	fmt.Print(g)
}

// Create undirected graph
func NewUndirectedGraph(size int, avgenum int, density float64, maxweight float64, connected bool) *Graph {
	g := Graph{make([][]Edge, size), size, avgenum, density, maxweight, connected}
	for i := range g.AdjList {
		g.AdjList[i] = make([]Edge, 0)
	}
	// total number of edges need to create
	totalEdges := 0
	if avgenum >= 0 {totalEdges = avgenum*size/2}
	// if both avgenum and density is non-negative, use density 
	if density >= 0 {totalEdges = int(density * float64((size-1)*size/2))}
	// Keep track of which vertex pair alreay has an edge
	// so that we never create repeat edges
	exsit := make(map[int]bool, 2*totalEdges)
	if connected {
		ConnectGraph(&g, exsit)
	}
	BuildGraph(&g, totalEdges, exsit)
	return &g
}

// Randomly build edges in graph
// No two edges will have the same head and tail
func BuildGraph(g *Graph, e int, exsit map[int]bool) {
	hashFactor := g.Size
	for len(exsit)<2*e {
		head := rand.Intn(g.Size)
		tail := rand.Intn(g.Size)
		if _, ok := exsit[head * hashFactor + tail]; !ok && head!=tail {
			exsit[head * hashFactor + tail] = true
			exsit[tail * hashFactor + head] = true 
			g.add(head, tail)
		}
	}
}

// Randomly generate a connected path inside the graph
func ConnectGraph(g *Graph, exsit map[int]bool) {
	hashFactor := g.Size
	l := g.AdjList
	// A map to record which vertices havn't been connected to the graph
	unconnected := make(map[int]bool, len(l))
	for i := 0; i<len(l); i++ {unconnected[i] = true}
	prev := -1
	head := -1
	for k, _ := range unconnected {
		if prev >= 0 {
			g.add(prev, k)
			exsit[prev * hashFactor + k] = true
			exsit[k * hashFactor + prev] = true 
		} else {
			head = k
		}
		prev = k
	}
	g.add(prev, head)
	exsit[prev * hashFactor + head] = true
	exsit[head * hashFactor + prev] = true 
}

// Add edge
func (g *Graph) add(a int, b int) {
	w := rand.Float64()*g.MaxWeight
	g.AdjList[a] = append(g.AdjList[a], Edge{b, w})
	g.AdjList[b] = append(g.AdjList[b], Edge{a, w}) 
}

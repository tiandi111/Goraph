package goraph

import (
	"fmt"
	"math"
)

type gNode struct {
	v	int
	idx	int	// the heap index of the gNode
	bw	float64
}

func (n *gNode) cmp(a interface{}) int {
	x := n.bw - a.(*gNode).bw 
	switch {
		case x > 0: return 1
		case x < 0: return -1
		case x == 0: return 0
	}
	return 0
}

func (n *gNode) setIdx(i int) bool {
	n.idx = i
	return true
}

// Set the bw of the node, this operation will effect the position 
// of the given gNode inside the heap, so we need to know the current
// index 
func (n *gNode) setBw(q Heap, b float64) {
	n.bw = b
	q.set(n.idx, n)
}

// Three algorithms are provided for maximum bandwidth path problem
// "Dijkstra"      - Plain dijkstra
// "Dijkstra_Heap" - heap implementation for dijkstra
// "Kruskal"       - Kruskal
func MBP(g *Graph, s int, t int, alg string) float64 {
	switch alg {
		case "Dijkstra":
			return Dij(g, s, t)
		case "Dijkstra_Heap":
			return DijHeap(g, s, t)
		case "Kruskal":
			return Kru(g, s, t)
		default:
			fmt.Println("Please indicate a correct algorithm!")
			return -1
	}
	return -1
}

// Dijkstra without heap
func Dij(g *Graph, s int, t int) float64 {
	// Initialize two arrays
	// bw for recording maximum bandwidth
	// visited for recording which vertices havn'r been visited
	bw := make( []float64, g.Size )
	visited := make( map[int]bool, 0 )
	bw[s] = g.MaxWeight
	// Main while loop
	for len(visited) < g.Size {
		// Find the vertex with maximum bandwitdth in unvisited set
		cur := -1
		curBw := float64(-1)
		for i, mbw := range bw {
			_, seen := visited[i]
			if !seen && mbw > curBw {
				cur = i
				curBw = mbw
			}
		}
		// Mark it as visited
		visited[cur] = true
		// Update bandwidth
		for _, edge := range g.AdjList[cur] {
			nei := edge.tail
			min := math.Min( bw[cur], edge.weight )
			if min > bw[nei] {
				bw[nei] = min
			}
		}
		//fmt.Println(len(visited))
	}
	return bw[t]
}

// Modified Dijkstra algorithm, using heap
func DijHeap(g *Graph, s int, t int) float64 {
	// Array to record maximum bandwidth value
	mbw := make( []HeapNode, g.Size )
	
	q := NewHeap()
	
	// Initilization source node to have infinity bandwidth
	// and others to have 0
	for i := 0; i < len(mbw); i++ {
		mbw[i] = &gNode{ i, i, 0 }
		q.add( mbw[i] )
	}
	mbw[s].(*gNode).setBw(q, g.MaxWeight)

	// Main while loop
	for q.size() > 0 {
		// Take the node with maximum bandwidth
		// Then remove it from queue
		cur, _ := q.get(0)
		q.del(0)
		// For each edge, take the min value of bandwidth of
		// cur node and that of nei node, if it is greater 
		// than the maximum bandwidth of nei node till now,
		// replace it; otherwise, go for the next nei node
		for _, edge := range g.AdjList[cur.(*gNode).v] {
			nei := edge.tail
			min := math.Min(cur.(*gNode).bw, edge.weight)
			if min > mbw[nei].(*gNode).bw {
				mbw[nei].(*gNode).setBw(q, min)
			}
		}
	}
	
	return mbw[t].(*gNode).bw
}


// Kruskal algorithm
func Kru(g *Graph, s int, t int) float64 {
	size := g.Size
	// Initialize disjoint set 
	djs := NewDJS( size )
	// Sort edges
	// Push all edges to the heap
	// then pop them to an array
	h := NewHeap()
	for _, ver := range g.AdjList {
		for _, edge := range ver {
			h.add( edge )
		}
	}
	eSet := make( []Edge, h.size() )
	for i := 0; i < len(eSet); i++ {
		e, _ := h.get(0)
		eSet[i] = e.(Edge)
		h.del(0)
	}
	// Main while loop 
	MST := make( []Edge, 0 )
	for _, e := range eSet {
		head := e.head
		tail := e.tail
		if djs.find(head) != djs.find(tail) {
			MST = append(MST, e)
			djs.union(head, tail)
		}
	}
	// Construct Maximum spanning tree(graph)
	mstG := CreateGraph(MST, size, g.MaxWeight, g.Connected, false)
	// DFS to calculate maximum bandwidth
	return getMBP(mstG, s, t)
}

// For a tree-like graph, we can use dfs to fet the maximum bandwidth path from
// source to terminal node
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


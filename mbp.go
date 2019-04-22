package goraph

import (
	"fmt"
	"math"
)

type gNode struct {
	v	int
	idx	int
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
		//	return Dij(g, s, t)
		case "Dijkstra_Heap":
			return DijHeap(g, s, t)
		case "Kruskal":
		//	return Kru(g, s, t)
		default:
			fmt.Println("Please indicate a correct algorithm!")
			return -1
	}
	return -1
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

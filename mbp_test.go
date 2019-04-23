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
		r1  := MBP(g, sc, tm, "Dijkstra")
		r2  := MBP(g, sc, tm, "Dijkstra_Heap")
		if exp != r1 {
			t.Errorf("\nGraph: %q\n    Source:%d\n    Terminal:%d\n    Expect %f, but get %f", descr, sc, tm, exp, r1)
		}
		if exp != r2 {
			t.Errorf("\nGraph: %q\n    Source:%d\n    Terminal:%d\n    Expect %f, but get %f", descr, sc, tm, exp, r2)
		}
		//fmt.Println("Kruskal: ", exp, "Dij: ", r1, "Dij_Heap", r2)
	}
}

// Benchmark for:
//	Dijkstra without heap
//	Sparse graph
func Benchmark_Dij_Sparse(b *testing.B) {
	g := NewUndirectedGraph(5000, 6, -1, 10, true)
	sc := rand.Intn(g.Size)
	tm := rand.Intn(g.Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MBP(g, sc, tm, "Dijkstra")
	}
}

// Benchmark for:
//	Dijkstra without heap
//	Dense graph
func Benchmark_Dij_Dense(b *testing.B) {
	g := NewUndirectedGraph(5000, -1, 0.2, 10, true)
	sc := rand.Intn(g.Size)
	tm := rand.Intn(g.Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MBP(g, sc, tm, "Dijkstra")
	}
}

// Benchmark for:
//	Dijkstra with heap
//	Sparse graph
func Benchmark_Dij_Heap_Sparse(b *testing.B) {
	g := NewUndirectedGraph(5000, 6, -1, 10, true)
	sc := rand.Intn(g.Size)
	tm := rand.Intn(g.Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MBP(g, sc, tm, "Dijkstra_Heap")
	}
}

// Benchmark for:
//	Dijkstra with heap
//	Dense graph
func Benchmark_Dij_Heap_Dense(b *testing.B) {
	g := NewUndirectedGraph(5000, -1, 0.2, 10, true)
	sc := rand.Intn(g.Size)
	tm := rand.Intn(g.Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MBP(g, sc, tm, "Dijkstra_Heap")
	}
}

// Benchmark for:
//	Kruskal
//	Sparse graph
func Benchmark_Kruskal_Sparse(b *testing.B) {
	g := NewUndirectedGraph(5000, 6, -1, 10, true)
	sc := rand.Intn(g.Size)
	tm := rand.Intn(g.Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MBP(g, sc, tm, "Kruskal")
	}
}

// Benchmark for:
//	Kruskal
//	Dense graph
func Benchmark_Kruskal_Dense(b *testing.B) {
	g := NewUndirectedGraph(5000, -1, 0.2, 10, true)
	sc := rand.Intn(g.Size)
	tm := rand.Intn(g.Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MBP(g, sc, tm, "Kruskal")
	}
}


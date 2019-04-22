package goraph

import (

)

// DJS interface encapsulate the basic operaions on
// a disjoint set
// Description:
//	find  - find the root vertex of the given vertex
//	union - union two given vertex into a single set
type DJS interface {
	find (a int) int
	union(a, b int) bool
	size()	int
}

type djs struct {
	set	[]int
	Size	int
}

// Create a new DJS
func NewDJS(size int) DJS {
	set := make([]int, size)
	for i := 0; i < size; i++ {
		set[i] = i
	}
	return &djs{ set, size }
}

// Find root
func (s *djs) find(a int) int {
	// if parent == a, a is root
	p := s.set[a]
	if p == a {
		return a
	}
	// otherwise, continue to trace
	// To reduce the height of the set,
	// update the root while tracing
	s.set[a] = s.find(p)
	return s.set[a]
}

// Union two vertex
func (s *djs) union(a, b int) bool {
	pa := s.find(a)
	pb := s.find(b)
	s.set[pa] = pb
	return true
}

// get size
func (s *djs) size() int {
	return s.Size
}



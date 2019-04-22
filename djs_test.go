package goraph

import (
	"testing"
)

// Test find
func TestFind(t *testing.T) {
	// initialize a disjiont set
	s := djs { make([]int, 6), 6 }
	// set all vertex's parent to its next vertex
	for i := 0; i < 6; i++ {
		s.set[i] = i+1
	}
	// root is the last vertex
	s.set[5] = 5
	// test root
	for i := 0; i < 6; i++ {
		r := s.find(i)
		if r != 5 {
			t.Errorf("The %dth vertex has incorrect root, expect %d, but %d", i, 5, r)
		}
	}
	// test root update
	for i, root := range s.set {
		if root != 5 {
			t.Errorf("The %dth vertex'root is not updated, expect %d, but %d", i, 5, root)
		}
	}
}

// Test union
func TestUnion(t *testing.T) {
	s := djs { make([]int, 6), 6 }
	for i := 0; i < 6; i++ {
		s.set[i] = i
	}
	s.union(0, 1)
	s.union(1, 2)
	s.union(2, 3)
	s.union(4, 5)
	for i := 0; i < 3; i++ {
		pa := s.find(i)
		pb := s.find(i+1)
		if pa != pb {
			t.Errorf("The %dth vertex has incorrect root, expect %d, but %d", i, pa, pb)
		}
	}
	if s.find(4) != s.find(5) {
		t.Errorf("The 5th vertex has incorrect root, expect %d, but %d", s.set[4], s.set[5])
	}
	if s.find(3) == s.find(4) {
		t.Errorf("The 3rd and 4th verteices should not belong to the same set!")
	}
}

func TestDJS(t *testing.T) {
	size := 6
	djs := NewDJS(6)
	if size != djs.size() {
		t.Errorf("Size is incorrect, expect %d, but %d", size, djs.size())
	}
	for i := 0; i < size; i++ {
		p := djs.find(i)
		if i != p {
			t.Errorf("The %dth vertex in initialized set should have parent %d, but %d", i, i, p)
		}
	}
	djs.union(0, 1)
	if djs.find(0) != djs.find(1) {
		t.Errorf("Union failed!")
	}
}

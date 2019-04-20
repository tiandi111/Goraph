package goraph

import (
	"testing"
)

type Event struct {
	opr	byte
	opd	interface{}
	res	interface{}
}

func TestHeap(t *testing.T) {
	testcases := [][]Event {
		{
			Event{'a', 1, true}, Event{'a', 2, true}, Event{'a', 3, true},
			Event{'a', 4, true}, Event{'m', 0, 4}, Event{'a', 5, true}, 
			Event{'m', 0, 5}, Event{'d', 1, true}, Event{'m', 0, 5},
			Event{'d', 0, true}, Event{'m', 0, 4}, 
		},
	}
	for _, tc := range testcases {
		h := newMaxHeap()
		for i, e := range tc {
			if r := Execute(h, e); r != e.res {
				t.Errorf("The %dth event return %v, but expect %v", i, r, e.res)
			}
			if !isValidHeap(h) {
				t.Errorf("Heap is not valid after %dth event", i)
			}
		}
	}
}

// execute the given event 
func Execute(h *Heap, e *Event) interface{} {
	switch e.opr {
		case 'a': return h.add(e.opd) 
		case 'm': return h.max()
		case 'd': return h.del(e.opd)
	}
	return false
}

// iterate the heap to check its validity
func isValidHeap(h *Heap) bool {
	for i := 0; i < Heap.Size(); i++ {
		left := i*2+1
		right := i*2+2
		if left < Heap.Size() && Heap.get(i)<Heap.get(left) {return false}
		if right < Heap.Size() && Heap.get(i)<Heap.get(right) {return false}
	}
	return true
}

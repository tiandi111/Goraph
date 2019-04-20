package goraph

import (
	"fmt"
)

// Heap struct
// BODY is implemented with a slice
// SIZE point to the position new element will be inserted
type Heap struct {
	BODY	[]WeightedNode
	SIZE	int
}

// WeightedNode should implement two methods
// getKey return the key of the node
// getVal return the value of the node
type WeightedNode interface {
	getKey (i int) interface{}
	getVal (i int) interface{}
}

// return the size of the heap
func (h *Heap) size() int {
	return h.SIZE
}

// return the max value of the heap
func (h *Heap) max() float64, error {
	if err := h.validIndex(0); err { 
		return 0, err
	}
	return h.getVal(0)
}

// add node to the heap
// first add to the end, then do a bottom-up
// heapify called swim()
// return the size after inserting the new node
func (h *Heap) add(e WeightedNode) int{
	h.BODY = append(h.BODY, e)
	h.swim(h.size())
	return h.SIZE++
}

// return the element
func (h *Heap) get(i int) interface{}, error {
	if err := h.validIndex(i); err {
		return interface{}, err
	}
	return h.BODY[i]
}

// return the key, error is non-nil if the index is invalid
func (h *Heap) getKey(i int) interface{}, error {
	if err := h.validIndex(i); err {
		return interface{}, err 
	}
	return h.BODY[i].getKey(i)
}

// return the value, error is non-nil if the index is invalid
func (h *Heap) getVal(i int) interface{}, error {
	if err := h.validIndex(i); err {
		return interface{}, err
	}
	return h.BODY[i].getVal(i)
} 

// check if the given index is valid
func (h *Heap) validIndex(i int) bool{
	if i < 0 || i >= h.size() { return false }
	return true
}

// swap element on a and b
func (h *Heap) swap(a int, b int) bool, error {
	if err := h.validIndex(a); err {
		return false, err
	}
	if err := h.validIndex(b); err {
		return false, err
	}
	tmp := h.get(a)
	h.set(a, h.get(b))
	h.set(b, tmp)
	return true
}

func (h *Heap) del(i int) int {
	
}

// compare the val of the ith element with its parent
// if greater, swap
func (h *Heap) swim(i int) error {
	if err := validIndex(i); err {return err}
	for p := (i-1)/2; p >= 0; {
		if h.compareTo(i, p) > 0 {
			h.swap(i, p)
		}
		else {
			return nil
		}
		i = p
		p = (i-1)/2
	} 
	return nil
}

func (h *Heap) sink(i int) error {
	if err := validIndex(i); err {return err}
	for c := 2*i+1; c < h.size() {
		swp := 2*i +1
		if 2*i+2 < h.size() {
			
		}
	}
	return nil
}


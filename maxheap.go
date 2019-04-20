package goraph

import (
	"fmt"
)

// Heap interface encapsulates the basic operations on a heap
// This interface isolates the underlying heap struct to users
// Function description:
//	add:	add an element to heap
//	del:	delete the element on the given index
//	max:	return maximum element
//	size:	return size
//	get:	return element on the given index
type Heap interface {
	add (e... HeapNode)	(bool, error)
	del (i int)		(bool, error)
	max ()			(HeapNode, error)
	size()			int
	get (i int)		(HeapNode, error)
}

// The underlying heap struct
// Body is implemented as a slice of HeapNode
// The type of element inserted to heap must implement HeapNode
type heap struct {
	Body	[]HeapNode
	Size	int
}

// HeapNode should implement one function
// Function Description:
//	cmp: 	compare a HeapNode with itself
//		if a > b, return 1
//		if a = b, return 0
//		if a < b, return -1
type HeapNode interface {
	cmp (a interface{})	int
}

// Consturctor for Heap
func NewHeap() Heap{
	return &heap{make([]HeapNode, 0), 0}
}

func (h *heap) add(e ...HeapNode) (bool, error) {
	// e... unpack slice e to multiple HeapNode
	return h.addNode(e...)
}

func (h *heap) del(i int) (bool, error) {
	return h.delNode(i)
}

func (h *heap) max() (HeapNode, error) {
	return h.maxNode()
}

func (h *heap) size() int {
	return h.getSize()
}

func (h *heap) get(i int) (HeapNode, error) {
	return h.getNode(i)
}

// Add nodes to heap
func (h *heap) addNode(e ...HeapNode) (bool, error) {
	arr := e
	for _, node := range arr {
		h.addSingle(node)
	}
	return true, nil
}

// Add single node to heap
func (h *heap) addSingle(e HeapNode) (bool, error) {
	h.Body = append(h.Body, e)
	h.Size++
	h.swim(h.Size-1)
	return true, nil
}

// Delete node on given index
func (h *heap) delNode(i int) (bool, error) {
	if err := h.checkIndex(i); err != nil {
		return false, err
	}
	// Note: swap should appear before size--
	// Because swap utilize checkIndex() which return error
	// if given idex is equal to h.Size
	// ps: Never send h.Size to function swap, swim and sink!!!!!
	h.swap(i, h.Size-1)
	h.Size--
	h.swim(i)
	h.sink(i)
	return true, nil
}

// Return the maximum element
func (h *heap) maxNode() (HeapNode, error) {
	if err := h.checkIndex(0); err != nil {
		return nil, err
	}
	return h.Body[0], nil
}

// Return size of the heap 
func (h *heap) getSize() int {
	return h.Size
}

// Get element on the given index
func (h *heap) getNode(i int) (HeapNode, error) {
	if err := h.checkIndex(i); err != nil {
		return nil, err
	}
	return h.Body[i], nil
}

// Swap nodes on index a and index b
func (h *heap) swap(a int, b int) error{
	if err := h.checkIndex(a, b); err != nil {
		return err
	}
	B := h.Body
	B[a], B[b] = B[b], B[a]
	return nil
}

// Check the validity of index
func (h *heap) checkIndex(i ...int) error {
	arr := i
	for _, index := range arr {
		if index < 0 || index >= h.Size {
			return fmt.Errorf("Index out of bounds: %d", index)
		}
	}
	return nil
}

// Function swim is to Bottom-up heapify the heap 
func (h *heap) swim(i int) error {
	if err := h.checkIndex(i); err != nil {
		return err
	}
	for p := (i-1)/2; p >= 0; {
		// If ith element larger than its parent, swap
		// otherwise, return
		if h.Body[i].cmp(h.Body[p]) > 0 {
			h.swap(i, p)
		} else {
			return nil
		}
		i = p
		p = (i-1)/2
	} 
	return nil
}

// Function sink is to Top-down heapify the heap
func (h *heap) sink(i int) error {
	if err := h.checkIndex(i); err != nil {
		return err
	}
	for j := 2*i+1; j < h.Size; {
		// Choose a larger child to compare
		if j+1 < h.Size && h.Body[j+1].cmp(h.Body[j]) > 0 {
			j++
		}
		// If ith element smaller than its child, swap
		// otherwise, return
		if h.Body[i].cmp(h.Body[j]) < 0 {
			h.swap(i, j)
		} else {
			return nil
		}
		i = j
		j = 2*i+1
	}
	return nil
}


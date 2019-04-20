package goraph

import (
	"testing"
	"strings"
	"strconv"
	"fmt"
)

// Event represent an operation on heap
// opr(operation) description:
//	0:	add
//	1: 	del
//	2:	max
//	3: 	size
//	4:	get
// opd(operand)	and res(result) are
type Event struct {
	opr	int
	opd	interface{}
	exp	interface{}
}

// Node struct for test
type Node struct {
	weight	int
}

// Implement cmp 
func (n Node) cmp(a interface{}) int {
	return n.weight-a.(Node).weight
}

// Test swap
func TestSwap(t *testing.T) {
	h := heap{make([]HeapNode, 0), 0}

	h.add( Node{0} )
	h.add( Node{1} )

	oZero, _ := h.get(0)
	oOne, _  := h.get(1)
	exp := []HeapNode{oOne, oZero}

	h.swap(0, 1)

	rZero, _ := h.get(0)
	rOne, _ := h.get(1)

	if oZero != rOne || oOne != rZero {
		t.Errorf("Expect %v, but %v", exp, h.Body)
	}
}

// Test swim
func TestSwim(t *testing.T) {
	h := heap{make([]HeapNode, 0), 4}

	h.Body = []HeapNode{ Node{0}, Node{1}, Node{2}, Node{3} }
	exp := []HeapNode{ h.Body[3], h.Body[0], h.Body[2], h.Body[1] }

	h.swim(3)

	for i, n := range h.Body {
		if exp[i] != n {
			t.Errorf("Incorrect %dth element, Expect %v, but %v", i, exp[i], n)
		}
	}
}

// Test Sink
func TestSink(t *testing.T) {
	h := heap{ make([]HeapNode, 0), 5 }

	h.Body = []HeapNode { Node{0}, Node{2}, Node{1}, Node{3}, Node{4} }
	exp := []int {2, 4, 1, 3, 0}

	h.sink(0)

	for i, n := range h.Body {
		if exp[i] != n.(Node).weight {
			t.Errorf("Incorrect %dth element, Expect %v, but %v", i, exp[i], n)
		}
	}

	//fmt.Println("Expected: ", exp)
	//fmt.Println("Your out: ", h.Body)
}

// Test add
func TestAdd(t *testing.T) {
	h := heap{make([]HeapNode, 0), 0}

	for i := 0; i < 4; i++ {
		h.addSingle( Node{i} )
	}
	exp := []int{ 3, 2, 1, 0 }

	for i, n := range h.Body {
		if exp[i] != n.(Node).weight {
			t.Errorf("Incorrect %dth element, Expect %v, but %v", i, exp[i], n)
		}
	}
}

// TODO: Test del


// Integration test
func TestHeap(t *testing.T) {
	oprStr := []string{"add", "del", "max", "size", "get"}
	testcases := [][]string {
		{
			"0 1 t", "0 2 t", "0 3 t", "2 0 3", "3 0 3", "1 2 t", "1 0 t", "1 5 f",
			"1 0 t", 
		},
		{
			"2 0 n", "3 0 0", "1 0 f", "4 0 n",
		},
		{
			"0 0 t", "2 0 0", 
			"0 1 t", "2 0 1", 
			"0 2 t", "2 0 2", 
			"0 3 t", "2 0 3", 
			"0 4 t", "2 0 4",
			"1 0 t", "2 0 3",
			"1 0 t", "2 0 2",
			"1 0 t", "2 0 1",
			"1 0 t", "2 0 0",
			"1 0 t", "2 0 n",
			"1 0 f", "2 0 n",
		},
	}
	for index, tc := range testcases {
		h := heap{make([]HeapNode, 0), 0}
		
		oprOut := make([]string, 0)
		opdOut := make([]interface{}, 0)
		eOut := make([]interface{}, 0)
		rOut := make([]interface{}, 0)

		for i, c := range tc {
			e := getEvent(c)
			r := Execute(&h, &e)
			if r != e.exp {
				t.Errorf("The %dth operation failed, expect %v, but %v", i, e.exp, r)
			}

			oprOut = append(oprOut, oprStr[e.opr])
			opdOut = append(opdOut, e.opd)
			eOut = append(eOut, e.exp)
			rOut = append(rOut, r)
		}
		
		fmt.Println("The ", index, "th Testcase: ")
		fmt.Println("	Operation:", oprOut)
		fmt.Println("	Operand:  ", opdOut)
		fmt.Println("	Expected: ", eOut)
		fmt.Println("	Your out: ", rOut)
	}
}

// Convert string to Event
func getEvent(in string) Event {
	s := strings.Split(in, " ")
	e := Event{}
	e.opr, _ = strconv.Atoi(s[0])
	setOpd( &e, s[1] )
	setExp( &e, s[2] )
	return e
}

// Convert string to operand for Event
func setOpd( e *Event, opd string ) {
	switch e.opr {
		case 0:
			val, _ := strconv.Atoi(opd)
			e.opd = Node{ val }
		default:
			e.opd, _ = strconv.Atoi(opd)
	}
}

// Convert string to expectation for Event 
func setExp( e *Event, exp string ) {
	switch e.opr {
		case 0, 1:
			if exp == "t" {
				e.exp = true
			} else {
				e.exp = false
			}
		case 2, 4:
			if exp == "n" {
				e.exp = nil
			} else {
				e.exp, _ = strconv.Atoi(exp)
			}
		case 3:
			e.exp, _ = strconv.Atoi(exp)
	}
}

// Execute the given event 
func Execute(h *heap, e *Event) interface{} {
	var r interface{}
	switch e.opr {
		// add
		case 0: 
			r, _ = h.add( e.opd.(HeapNode) ) // Here need type assertion
		// del
		case 1: 
			r, _ = h.del( e.opd.(int) )
		// max
		case 2: 
			n, _ := h.max()
			if n != nil {
				r = n.(Node).weight
			} 
		// size
		case 3: 
			r = h.size()
		// get
		case 4: 
			n, _ := h.get( e.opd.(int) )
			if n != nil {
				r = n.(Node).weight
			}
		
	}
	return r
}

// iterate the heap to check its validity
func isValidHeap(h *heap) bool {
	s := h.size()
	for i := 0; i < s; i++ {
		left := i*2+1
		right := i*2+2
		fmt.Println(left, right)
		//if left < s && h.get(i).cmp( h.get(left) ) < 0 {return false}
		//if right < s && h.get(i).cmp( h.get(right) ) < 0 {return false}
	}
	return true
}

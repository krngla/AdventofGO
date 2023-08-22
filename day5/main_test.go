package main

import (
	"testing"
)

func TestInsertElement(t *testing.T) {
	type input struct {
		element     string
		destination int
	}
	type testmatrix struct {
		ipt      []input
		expected string
	}
	tc := testmatrix{
		ipt:      []input{{"A]", 2}, {"B]", 3}, {"C", 1}},
		expected: "4,,C,A,B,",
	}
	//ic := inputcase{
	//	Columns: make([]stack.Stack, 0),
	//	Ncol:    0,
	//	Ops:     make([]operation, 0),
	//}
	ic := inputcase{}
	for _, inp := range tc.ipt {
		ic.insertElement(inp.element, inp.destination)
	}
	ser := ic.Print()
	if ser != tc.expected {
		t.Errorf("Failed expected:\n%s \n got:\n%s\n", tc.expected, ser)
	}
}

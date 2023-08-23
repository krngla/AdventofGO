package main

import (
	"testing"
)

var testobject []byte = []byte(`{"Columns":[{"Head":3,"Stack":["D","N","Z"]},{"Head":2,"Stack":["C","M"]},{"Head":1,"Stack":["P"]}],"Ncol":3}`)

type input struct {
	element     string
	destination int
}
type testmatrix struct {
	ipt      []input
	expected string
}

func TestInsertElement(t *testing.T) {

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

func TestMoveElement(t *testing.T) {

	ic := inputcase{}
	ic.Deser(testobject)
	t.Errorf(ic.Print())
	ic.moveElements(2, 1, 2)
	t.Errorf(ic.Print())
}

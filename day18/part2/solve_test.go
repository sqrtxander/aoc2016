package main

import "testing"

var INPUT string = `
.^^.^.^^^^
`[1:]

var EXPECTED int = 38

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 10)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

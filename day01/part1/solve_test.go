package main

import "testing"

var INPUT string = `
R5, L5, R5, R3
`[1:]

var EXPECTED int = 12

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

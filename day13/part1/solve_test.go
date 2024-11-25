package main

import "testing"

var INPUT string = `
10
`[1:]

var EXPECTED int = 11

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 7, 4)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

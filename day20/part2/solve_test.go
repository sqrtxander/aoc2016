package main

import "testing"

var INPUT string = `
5-8
0-2
4-7
`[1:]

var EXPECTED int = 2

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 9)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

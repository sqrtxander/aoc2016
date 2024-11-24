package main

import "testing"

var INPUT string = `
rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1
`[1:]

var EXPECTED int = 6

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 7, 3)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

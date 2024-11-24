package main

import "testing"

var INPUT string = `
X(8x2)(3x3)ABCY
`[1:]

var EXPECTED int = 18

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

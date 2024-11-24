package main

import "testing"

var INPUT string = `
ULL
RRDDD
LURDL
UUUUD
`[1:]

var EXPECTED int = 1985

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

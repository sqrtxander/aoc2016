package main

import "testing"

var INPUT string = `
ULL
RRDDD
LURDL
UUUUD
`[1:]

var EXPECTED string = "5DB3"

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %s got %s\n", EXPECTED, actual)
	}
}

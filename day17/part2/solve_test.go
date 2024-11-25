package main

import "testing"

var INPUT string = `
ulqzkmiv
`[1:]

var EXPECTED int = 830

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

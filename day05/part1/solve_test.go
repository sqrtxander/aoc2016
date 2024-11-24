package main

import "testing"

var INPUT string = `
abc
`[1:]

var EXPECTED string = "18f47a30"

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}

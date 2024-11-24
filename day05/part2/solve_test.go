package main

import "testing"

var INPUT string = `
abc
`[1:]

var EXPECTED string = "05ace8e3"

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}

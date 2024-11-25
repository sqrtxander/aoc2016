package main

import "testing"

var INPUT string = `
10000
`[1:]

var EXPECTED string = "01100"

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 20)
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}

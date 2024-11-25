package main

import "testing"

var INPUT string = `
ulqzkmiv
`[1:]

var EXPECTED string = "DRURDRUDDLLDLUURRDULRLDUUDDDRR"

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}

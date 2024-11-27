package main

import "testing"

var INPUT string = `
###########
#0.1.....2#
#.#######.#
#4.......3#
###########
`[1:]

var EXPECTED int = 14

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

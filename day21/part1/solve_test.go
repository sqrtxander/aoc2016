package main

import "testing"

var INPUT string = `
swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d
`[1:]

var EXPECTED string = "decab"

func TestSolve(t *testing.T) {
	actual := solve(INPUT, "abcde")
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}

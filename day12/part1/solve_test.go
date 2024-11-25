package main

import "testing"

var INPUT string = `
cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a
`[1:]

var EXPECTED int = 42

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

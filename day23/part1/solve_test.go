package main

import "testing"

var INPUT string = `
cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a
`[1:]

var EXPECTED int = 3

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

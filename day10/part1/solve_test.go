package main

import "testing"

var INPUT string = `
value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2
`[1:]

var EXPECTED int = 2

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 2, 5)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

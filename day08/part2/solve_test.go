package main

import "testing"

var INPUT string = `
rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1
`[1:]

var EXPECTED string = `
 #  # #
# #    
 #     `[1:]

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 7, 3)
	if actual != EXPECTED {
		t.Fatalf("Expected\n%s got\n%s\n", EXPECTED, actual)
	}
}

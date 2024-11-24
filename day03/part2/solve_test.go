package main

import "testing"

var INPUT string = `
101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603
`[1:]

var EXPECTED int = 6

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}

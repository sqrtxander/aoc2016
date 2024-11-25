package main

import "testing"

var INPUT string = `
10
`[1:]

var (
	INPUTS   = []int{0, 1, 2, 3, 4, 5}
	EXPECTED = []int{1, 3, 5, 6, 9, 11}
)

func TestSolve(t *testing.T) {
	for i := 0; i < len(INPUTS); i++ {
		actual := solve(INPUT, INPUTS[i])
		if actual != EXPECTED[i] {
			t.Fatalf("Expected %d got %d\n", EXPECTED[i], actual)
		}
	}
}

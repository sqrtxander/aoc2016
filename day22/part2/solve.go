package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

type node struct {
	used  int
	avail int
}

type stateSteps struct {
	state state
	steps int
}

type state struct {
	empty utils.Point
	goal  utils.Point
}

func cmpStatesMin(a, b stateSteps) int {
	return b.steps - a.steps
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")[2:]
	nodeGrid := make(map[utils.Point]node, len(lines))
	maxX := 0
	maxY := 0
	for _, line := range lines {
		line = line[len("/dev/grid/node-"):]
		parts := strings.Fields(line)
		xStr, yStr, _ := strings.Cut(parts[0], "-")
		x, err := strconv.Atoi(xStr[1:])
		if err != nil {
			log.Fatalf("Invalid x-oordinate: %q\n", xStr)
		}
		y, err := strconv.Atoi(yStr[1:])
		if err != nil {
			log.Fatalf("Invalid y-oordinate: %q\n", yStr)
		}
		usedStr := parts[2]
		used, err := strconv.Atoi(usedStr[:len(usedStr)-1])
		if err != nil {
			log.Fatalf("Invalid used amount: %q\n", usedStr)
		}
		availStr := parts[3]
		avail, err := strconv.Atoi(availStr[:len(availStr)-1])
		if err != nil {
			log.Fatalf("Invalid avail amount: %q\n", availStr)
		}
		maxX = max(maxX, x)
		maxY = max(maxY, y)
		nodeGrid[utils.Point{X: x, Y: y}] = node{
			used:  used,
			avail: avail,
		}
	}
	var emptyP utils.Point
	var emptyN node
	for p, n := range nodeGrid {
		if n.used == 0 {
			emptyP = p
			emptyN = n
		}
	}
	immovable := []utils.Point{}
	for p, n := range nodeGrid {
		if n.used > emptyN.avail+emptyN.used {
			immovable = append(immovable, p)
		}
	}
	init := stateSteps{
		state: state{
			empty: emptyP,
			goal:  utils.Point{X: maxX, Y: 0},
		},
		steps: 0,
	}
	seen := map[state]bool{}
	heap := utils.HeapFunc[stateSteps]{
		Heap: []stateSteps{init},
		Cmp:  cmpStatesMin,
	}
	for len(heap.Heap) > 0 {
		var s stateSteps
		heap, s = heap.Pop()
		if s.state.goal.X == 0 && s.state.goal.Y == 0 {
			return s.steps
		}
		for _, p := range utils.Adjacent4(s.state.empty) {
			if p.X < 0 || p.X > maxX || p.Y < 0 || p.Y > maxY {
				continue
			}
			if slices.Contains(immovable, p) {
				continue
			}
			var newState stateSteps
			if p.X == s.state.goal.X && p.Y == s.state.goal.Y {
				newState = stateSteps{
					state: state{
						empty: p,
						goal:  s.state.empty,
					},
					steps: s.steps + 1,
				}
			} else {
				newState = stateSteps{
					state: state{
						empty: p,
						goal:  s.state.goal,
					},
					steps: s.steps + 1,
				}
			}
			if !seen[newState.state] {
				heap = heap.Push(newState)
				seen[newState.state] = true
			}
		}
	}
	return -1
}

func main() {
	var inputPath string
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	} else {
		_, currentFilePath, _, _ := runtime.Caller(0)
		dir := filepath.Dir(currentFilePath)
		dir = filepath.Dir(dir)
		inputPath = filepath.Join(dir, "input.in")
	}
	contents, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading file %s:\n%v\n", inputPath, err)
		return
	}
	fmt.Println(solve(string(contents)))
}

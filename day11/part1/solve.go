package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"sort"
	"strings"
)

type state struct {
	floors   [4][]int // generator +ve with corresponding -ve chip
	elevator int
	steps    int
}

func printState(s state) {
	fmt.Printf("The elevator is at floor %d\n", s.elevator+1)
	for i, floor := range s.floors {
		fmt.Printf("Floor %d: %d\n", i+1, floor)
	}
}

func moveObjects(s state, os []int, dir int) state {
	result := clone(s)
	for _, o := range os {
		result.floors[result.elevator] = utils.RemoveFirst(result.floors[result.elevator], o)
	}
	for _, o := range os {
		result.floors[result.elevator+dir] = append(result.floors[result.elevator+dir], o)
	}
	return result
}

func clone(s state) (result state) {
	result.steps = s.steps
	result.elevator = s.elevator
	for i, f := range s.floors {
		for _, o := range f {
			result.floors[i] = append(result.floors[i], o)
		}
	}
	return
}

func hashState(s state) string {
	genMap := map[int]int{}
	chipMap := map[int]int{}
	for i, f := range s.floors {
		for _, o := range f {
			if o > 0 {
				genMap[o] = i
			} else {
				chipMap[-o] = i
			}
		}
	}
	pairs := [][2]int{}
	for oType := range genMap {
		pairs = append(pairs, [2]int{chipMap[oType], genMap[oType]})
	}

	sort.Slice(pairs, func(i int, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
	return fmt.Sprintf("%d%d", s.elevator, pairs)
}

func isValid(s state, from int) bool {
	currFloor := s.floors[s.elevator]
	prevFloor := s.floors[s.elevator+from]
	if utils.Any(currFloor, func(o int) bool {
		if o > 0 {
			return false
		}
		if slices.Contains(currFloor, -o) {
			return false
		}
		return !utils.All(currFloor, func(p int) bool {
			return p < 0
		})
	}) {
		return false
	}
	return !utils.Any(prevFloor, func(o int) bool {
		if o > 0 {
			return false
		}
		if slices.Contains(prevFloor, -o) {
			return false
		}
		return !utils.All(prevFloor, func(p int) bool {
			return p < 0
		})
	})
}

func isFinished(s state) bool {
	return utils.All(s.floors[:3], func(f []int) bool {
		return len(f) == 0
	})
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	lines := strings.Split(s, "\n")
	stateStr := utils.Map(lines, strings.Fields)
	init := state{floors: [4][]int{}, elevator: 0, steps: 0}

	seenMap := map[string]int{}
	seenIdx := 1
	for _, line := range stateStr {
		if line[4] == "nothing" {
			continue
		}
		floorStr := line[1]
		var floorIdx int
		switch floorStr {
		case "first":
			floorIdx = 0
		case "second":
			floorIdx = 1
		case "third":
			floorIdx = 2
		case "fourth":
			floorIdx = 3
		default:
			log.Fatalf("Invalid floor %q", floorStr)
		}
		for i, word := range line {
			var oType string
			if word == "generator" {
				oType = line[i-1]
			} else if word == "microchip" {
				oType = line[i-1][:len(line[i-1])-len("-compatible")]
			} else {
				continue
			}
			var typeIdx int
			var ok bool
			if typeIdx, ok = seenMap[oType]; !ok {
				typeIdx = seenIdx
				seenMap[oType] = typeIdx
				seenIdx++
			}
			if word == "generator" {
				init.floors[floorIdx] = append(init.floors[floorIdx], typeIdx)
			} else if word == "microchip" {
				init.floors[floorIdx] = append(init.floors[floorIdx], -typeIdx)
			}
		}
	}

	states := utils.Queue[state]{init}
	seen := []string{}
	for len(states) > 0 {
		var curr state
		states, curr = states.Pop()
		if isFinished(curr) {
			return curr.steps
		}
		curr.steps++
		currFloor := curr.floors[curr.elevator]
		// single microchip
		for _, object := range currFloor {
			// move object up
			if curr.elevator < 3 {
				newS := moveObjects(curr, []int{object}, 1)
				newS.elevator++
				hash := hashState(newS)
				if isValid(newS, -1) && !slices.Contains(seen, hash) {
					states = states.Push(newS)
					seen = append(seen, hash)
				}
			}
			// move object down
			if curr.elevator > 0 {
				newS := moveObjects(curr, []int{object}, -1)
				newS.elevator--
				hash := hashState(newS)
				if isValid(newS, 1) && !slices.Contains(seen, hash) {
					states = states.Push(newS)
					seen = append(seen, hash)
				}
			}
		}
		// two objects
		for i, first := range currFloor {
			for _, second := range currFloor[i+1:] {
				// move object up
				if curr.elevator < 3 {
					newS := moveObjects(curr, []int{first, second}, 1)
					newS.elevator++
					hash := hashState(newS)
					if isValid(newS, -1) && !slices.Contains(seen, hash) {
						states = states.Push(newS)
						seen = append(seen, hash)
					}
				}
				// move object down
				if curr.elevator > 0 {
					newS := moveObjects(curr, []int{first, second}, -1)
					newS.elevator--
					hash := hashState(newS)
					if isValid(newS, 1) && !slices.Contains(seen, hash) {
						states = states.Push(newS)
						seen = append(seen, hash)
					}
				}
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

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

func solve(s string, start string) string {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	state := utils.Deque[byte]{}
	state = append(state, start...)
	for _, line := range lines {
		inst := strings.Fields(line)
		switch inst[0] + inst[1] {
		case "swapposition":
			idx1, err := strconv.Atoi(inst[2])
			if err != nil {
				log.Fatalf("Invalid index: %q\n", inst[2])
			}
			idx2, err := strconv.Atoi(inst[5])
			if err != nil {
				log.Fatalf("Invalid index: %q\n", inst[5])
			}
			state[idx1], state[idx2] = state[idx2], state[idx1]
		case "swapletter":
			idx1 := slices.Index(state, inst[2][0])
			idx2 := slices.Index(state, inst[5][0])
			state[idx1], state[idx2] = state[idx2], state[idx1]
		case "rotateleft":
			steps, err := strconv.Atoi(inst[2])
			if err != nil {
				log.Fatalf("Invalid rotation amount: %q\n", inst[2])
			}
			for range steps {
				var ch byte
				state, ch = state.PopLeft()
				state = state.PushRight(ch)
			}
		case "rotateright":
			steps, err := strconv.Atoi(inst[2])
			if err != nil {
				log.Fatalf("Invalid rotation amount: %q\n", inst[2])
			}
			for range steps {
				var ch byte
				state, ch = state.PopRight()
				state = state.PushLeft(ch)
			}
		case "rotatebased":
			steps := slices.Index(state, inst[6][0]) + 1
			if steps >= 5 {
				steps++
			}
			for range steps {
				var ch byte
				state, ch = state.PopRight()
				state = state.PushLeft(ch)
			}
		case "reversepositions":
			idx1, err := strconv.Atoi(inst[2])
			if err != nil {
				log.Fatalf("Invalid index: %q\n", inst[2])
			}
			idx2, err := strconv.Atoi(inst[4])
			if err != nil {
				log.Fatalf("Invalid index: %q\n", inst[4])
			}
			for i, j := idx1, idx2; i <= (idx1+idx2)/2; i, j = i+1, j-1 {
				state[i], state[j] = state[j], state[i]
			}
		case "moveposition":
			idx1, err := strconv.Atoi(inst[2])
			if err != nil {
				log.Fatalf("Invalid index: %q\n", inst[2])
			}
			idx2, err := strconv.Atoi(inst[5])
			if err != nil {
				log.Fatalf("Invalid index: %q\n", inst[5])
			}
			if idx1 > idx2 {
				for i := idx1; i > idx2; i-- {
					state[i], state[i-1] = state[i-1], state[i]
				}
			} else {
				for i := idx1; i < idx2; i++ {
					state[i], state[i+1] = state[i+1], state[i]
				}
			}
		default:
			log.Fatalf("Invalid instruction: %q\n", inst[0]+" "+inst[1])
		}
	}

	return string(state)
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
	fmt.Println(solve(string(contents), "fdhgacbe"))
}

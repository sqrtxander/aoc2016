package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	d, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid number input: %q\n", s)
	}
	elves := utils.Queue[int]{}
	for i := range d {
		elves = elves.Push(i + 1)
	}
	for len(elves) > 1 {
		var to int
		elves, to = elves.Pop()
		elves, _ = elves.Pop()
		elves = elves.Push(to)
	}

	return elves.Peek()
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

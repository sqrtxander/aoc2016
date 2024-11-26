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
	left := utils.Deque[int]{}
	right := utils.Queue[int]{}
	for i := 0; i <= d/2; i++ {
		left = left.PushRight(i + 1)
	}
	for i := d/2 + 1; i < d; i++ {
		right = right.Push(i + 1)
	}
	for {
		if len(left) > len(right) {
			left, _ = left.PopRight()
		} else {
			right, _ = right.Pop()
		}
		if len(left) == 0 {
			return right[0]
		} else if len(right) == 0 {
			return left[0]
		}
		var l int
		var r int
		right, r = right.Pop()
		left = left.PushRight(r)
		left, l = left.PopLeft()
		right = right.Push(l)
	}
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

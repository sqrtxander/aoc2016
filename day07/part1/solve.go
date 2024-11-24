package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	count := 0
	for _, line := range lines {
		queue := utils.Queue[rune]{}
		brackets := 0
		isValid := false
		for _, char := range line {
			if char == '[' {
				brackets++
				queue = queue.Clear()
			} else if char == ']' {
				brackets--
				queue = queue.Clear()
			} else if len(queue) < 3 {
				queue = queue.Push(char)
			} else if char == queue[0] && queue[1] == queue[2] && char != queue[1] {
				if brackets > 0 {
					isValid = false
					break
				}
				isValid = true
				queue = queue.Push(char)
				queue, _ = queue.Pop()
			} else {
				queue = queue.Push(char)
				queue, _ = queue.Pop()
			}
		}
		if isValid {
			count++
		}
	}

	return count
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

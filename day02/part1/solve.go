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

	code := 0
	p := utils.Point{X: 1, Y: 1}
	for _, line := range lines {
		for _, dir := range line {
			switch dir {
			case 'U':
				p.Y = max(0, p.Y-1)
			case 'R':
				p.X = min(2, p.X+1)
			case 'L':
				p.X = max(0, p.X-1)
			case 'D':
				p.Y = min(2, p.Y+1)
			default:
				log.Fatalf("Unknown move: '%c'\n", dir)
			}
		}
		num := 3*p.Y + p.X + 1
		code = 10*code + num
	}

	return code
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

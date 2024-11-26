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

func solve(s string, rowsReq int) int {
	s = strings.TrimSpace(s)
	traps := utils.GetBoundedHashGrid(s, '.', '^')

	rowsDone := 1
	safeCount := 0
	for _, trap := range traps.Grid {
		if !trap {
			safeCount++
		}
	}
	for from, to := 0, 1; rowsDone < rowsReq; from, to = (from+1)%2, (to+1)%2 {
		for i := range traps.W {
			l := traps.Grid[utils.Point{X: i - 1, Y: from}]
			r := traps.Grid[utils.Point{X: i + 1, Y: from}]
			if l != r { // Karnaugh maps for the win
				traps.Grid[utils.Point{X: i, Y: to}] = true
			} else {
				traps.Grid[utils.Point{X: i, Y: to}] = false
				safeCount++
			}
		}
		rowsDone++
	}
	return safeCount
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
	fmt.Println(solve(string(contents), 40))
}

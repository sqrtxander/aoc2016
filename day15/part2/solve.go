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

type disc struct {
	positions int
	init      int
}

func isSolved(discs []disc, time int) bool {
	return utils.All(discs, func(d disc) bool {
		return (time+d.positions+d.init)%d.positions == 0
	})
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ".", "", -1)
	lines := strings.Split(s, "\n")
	instructions := utils.Map(lines, strings.Fields)
	discs := make([]disc, len(instructions)+1)
	for _, line := range instructions {
		discIdxStr := line[1][1:]
		discIdx, err := strconv.Atoi(discIdxStr)
		if err != nil {
			log.Fatalf("Invalid disc number: %q\n", discIdxStr)
		}
		discIdx -= 1
		positionsStr := line[3]
		positions, err := strconv.Atoi(positionsStr)
		if err != nil {
			log.Fatalf("Invalid number of posiitions: %q\n", positionsStr)
		}
		initStr := line[11]
		init, err := strconv.Atoi(initStr)
		if err != nil {
			log.Fatalf("Invalid initial posiition: %q\n", initStr)
		}

		discs[discIdx] = disc{
			positions: positions,
			init:      (init + discIdx + 1) % positions,
		}
	}
	discs[len(discs)-1] = disc{
		positions: 11,
		init:      len(discs) % 11,
	}
	i := 0
	for !isSolved(discs, i) {
		i++
	}
	return i
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

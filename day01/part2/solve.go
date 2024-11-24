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

func solve(s string) int {
	s = strings.TrimSpace(s)
	moves := strings.Split(s, ", ")
	loc := utils.ORIGIN()
	dir := utils.NORTH
	seen := []utils.Point{}
	for _, move := range moves {
		rotateDir := move[0]
		amountStr := move[1:]
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			log.Fatalf("Invalid move amount: %q\n", amountStr)
		}
		switch rotateDir {
		case 'L':
			dir.RotateCCW()
		case 'R':
			dir.RotateCW()
		default:
			log.Fatalf("Invalid turn direction: '%c'\n", rotateDir)
		}
        for range amount {
            loc.MoveInDir(dir, 1)
            if slices.Contains(seen, loc) {
                return loc.Manhattan()
            }
            seen = append(seen, loc)
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

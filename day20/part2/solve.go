package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

type blockedRange struct {
	lo uint32
	hi uint32
}

func solve(s string, maxNum uint32) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	blocked := make(utils.Stack[blockedRange], len(lines))

	for i, line := range lines {
		loStr, hiStr, _ := strings.Cut(line, "-")
		lo, err := strconv.ParseUint(loStr, 10, 32)
		if err != nil {
			log.Fatalf("Invalid number: %q", loStr)
		}
		hi, err := strconv.ParseUint(hiStr, 10, 32)
		if err != nil {
			log.Fatalf("Invalid number: %q", hiStr)
		}
		blocked[i] = blockedRange{
			lo: uint32(lo),
			hi: uint32(hi),
		}
	}

	slices.SortFunc(blocked, func(a, b blockedRange) int {
		return int(b.lo) - int(a.lo)
	})
	cleanBlocked := utils.Stack[blockedRange]{}
	for len(blocked) > 1 {
		var a blockedRange
		var b blockedRange
		blocked, a = blocked.Pop()
		blocked, b = blocked.Pop()
		if a.hi+1 >= b.lo {
			blocked = blocked.Push(blockedRange{
				lo: a.lo,
				hi: max(a.hi, b.hi),
			})
		} else {
			cleanBlocked = cleanBlocked.Push(a)
			blocked = blocked.Push(b)
		}
	}
	cleanBlocked, a := cleanBlocked.Pop()
	blocked, b := blocked.Pop()
	if a.hi >= b.lo-1 {
		cleanBlocked = cleanBlocked.Push(blockedRange{
			lo: a.lo,
			hi: max(a.hi, b.hi),
		})
	} else {
		cleanBlocked = cleanBlocked.Push(a)
		cleanBlocked = cleanBlocked.Push(b)
	}

	allowedCount := 0
	for i := range len(cleanBlocked) - 1 {
		allowedCount += int(cleanBlocked[i+1].lo - cleanBlocked[i].hi - 1)
	}
	allowedCount += int(maxNum - cleanBlocked[len(cleanBlocked)-1].hi)
	return allowedCount
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
	fmt.Println(solve(string(contents), math.MaxUint32))
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

type node struct {
	used  int
	avail int
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")[2:]
	nodes := make([]node, len(lines))
	for i, line := range lines {
		line = line[len("/dev/grid/node-"):]
		parts := strings.Fields(line)
		usedStr := parts[2]
		used, err := strconv.Atoi(usedStr[:len(usedStr)-1])
		if err != nil {
			log.Fatalf("Invalid used amount: %q\n", usedStr)
		}
		availStr := parts[3]
		avail, err := strconv.Atoi(availStr[:len(availStr)-1])
		if err != nil {
			log.Fatalf("Invalid avail amount: %q\n", availStr)
		}
		nodes[i] = node{
			used:  used,
			avail: avail,
		}
	}
	slices.SortFunc(nodes, func(a, b node) int {
		return b.avail - a.avail
	})
	store := map[int]int{}
	count := 0
	for i, n := range nodes {
		if n.used == 0 {
			continue
		}
		if viable, ok := store[n.used]; ok {
			count += viable
			if i < viable {
				count--
			}
			continue
		}
		// binary search
		result := 0
		l, r := 0, len(nodes)-1
		for l <= r {
			m := (l + r) / 2
			if nodes[m].avail < n.used {
				result = m
				r = m - 1
			} else {
				l = m + 1
			}
		}
		store[n.used] = result
		count += result
		if i < result {
			count--
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

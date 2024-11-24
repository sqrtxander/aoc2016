package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	pattern := regexp.MustCompile(`(?P<name>[a-z-]+)-(?P<sid>\d+)\[(?P<chsum>[a-z]{5})\]`)

	result := 0
	for _, line := range lines {
		match := pattern.FindStringSubmatch(line)
		name := match[pattern.SubexpIndex("name")]
		sidStr := match[pattern.SubexpIndex("sid")]
		chsum := match[pattern.SubexpIndex("chsum")]

		sid, err := strconv.Atoi(sidStr)
		if err != nil {
			log.Fatalf("Invalid sector ID: %q\n", sidStr)
		}

		name = strings.ReplaceAll(name, "-", "")
		occursMap := map[byte]int{}
		for _, char := range name {
			occursMap[byte(char)] += 1
		}

		occurs := make([]utils.Pair[byte, int], 0, len(occursMap))
		for k, v := range occursMap {
			occurs = append(occurs, utils.Pair[byte, int]{K: k, V: v})
		}

		sort.Slice(occurs, func(a int, b int) bool {
			if occurs[a].V == occurs[b].V {
				return occurs[a].K < occurs[b].K
			}
			return occurs[a].V > occurs[b].V
		})

		expectedChsum := ""
		for _, p := range occurs[:5] {
			expectedChsum += string(p.K)
		}
		if expectedChsum == chsum {
			result += sid
		}
	}

	return result
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

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

const (
	normal int = iota
	findingMarker
)

func decrementRepeaters(repeaters []utils.Pair[int, int]) (newRepeaters []utils.Pair[int, int], amount int) {
	amount = 1
	for _, rpt := range repeaters {
		rpt.K--
		if rpt.K >= 0 {
			newRepeaters = append(newRepeaters, rpt)
			amount *= rpt.V
		}
	}
	return
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	state := normal
	result := 0
	marker := ""
	repeaters := []utils.Pair[int, int]{}
	for _, char := range s {
		if char == '(' && state == normal {
			state = findingMarker
			repeaters, _ = decrementRepeaters(repeaters)
			continue
		} else if char == ')' && state == findingMarker {
			repeaters, _ = decrementRepeaters(repeaters)
			amtStr, rptStr, _ := strings.Cut(marker, "x")
			amt, err := strconv.Atoi(amtStr)
			if err != nil {
				log.Fatalf("Bad amount in %q", marker)
			}
			rpt, err := strconv.Atoi(rptStr)
			if err != nil {
				log.Fatalf("Bad repeat in %q", marker)
			}
			marker = ""
			repeaters = append(repeaters, utils.Pair[int, int]{K: amt, V: rpt})
			state = normal
			continue
		}
		var amount int
		repeaters, amount = decrementRepeaters(repeaters)
		if state == normal {
			result += amount
		} else if state == findingMarker {
			marker += string(char)
		} else {
			log.Fatalln("Unknown state")
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

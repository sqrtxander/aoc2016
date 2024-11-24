package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"maps"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

func areRemainingActions(bots map[string][]int) bool {
	return utils.Any(slices.Collect(maps.Values(bots)), func(b []int) bool { return len(b) == 2 })
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	instructions := utils.Map(lines, strings.Fields)
	init := utils.Filter(instructions, func(strs []string) bool {
		return strs[0] == "value"
	})
	_ = init
	instructions = utils.Filter(instructions, func(strs []string) bool {
		return strs[0] == "bot"
	})
	bots := map[string][]int{}
	outputs := map[string][]int{}

	for _, in := range init {
		valStr := in[1]
		to := in[5]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			log.Fatalf("Invalid value: %q\n", valStr)
		}
		bots[to] = append(bots[to], val)
	}

	for areRemainingActions(bots) {
		for _, inst := range instructions {
			from := inst[1]
			if len(bots[from]) != 2 {
				continue
			}
			toTypeL := inst[5]
			toL := inst[6]
			toTypeH := inst[10]
			toH := inst[11]

			low := slices.Min(bots[from])
			high := slices.Max(bots[from])
			if toTypeL == "output" {
				outputs[toL] = append(outputs[toL], low)
			} else if toTypeL == "bot" {
				bots[toL] = append(bots[toL], low)
			} else {
				log.Fatalf("Invalid position for low: %q\n", toTypeL)
			}
			if toTypeH == "output" {
				outputs[toH] = append(outputs[toH], high)
			} else if toTypeH == "bot" {
				bots[toH] = append(bots[toH], high)
			} else {
				log.Fatalf("Invalid position for high: %q\n", toTypeL)
			}
			bots[from] = []int{}
		}
	}
	return outputs["0"][0] * outputs["1"][0] * outputs["2"][0]
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

package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	count := 0
	ready := 0
	triangles := [3][]int{{}, {}, {}}
	for _, line := range lines {
		ready = (ready + 1) % 3
		lengths, err := utils.CastToInts(strings.Fields(line))
		if err != nil {
			log.Fatalln(err)
		}
		for i := range 3 {
			triangles[i] = append(triangles[i], lengths[i])
		}
		if ready == 0 {
			for _, lengths := range triangles {
				m := slices.Max(lengths)
				if utils.Sum(lengths...) > 2*m {
					count++
				}
			}
			triangles = [3][]int{{}, {}, {}}
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

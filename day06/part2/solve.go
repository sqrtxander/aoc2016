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

func solve(s string) string {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	cols := [][]byte{}
	for range len(lines[0]) {
		cols = append(cols, []byte{})
	}
	for _, line := range lines {
		for i, char := range line {
			cols[i] = append(cols[i], byte(char))
		}
	}

	result := ""
	for _, col := range cols {
		result += string(utils.LeastFrequent(col))
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

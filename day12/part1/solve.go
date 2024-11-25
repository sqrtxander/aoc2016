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

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	instructions := utils.Map(lines, strings.Fields)

	values := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	}
	i := 0
	for i < len(instructions) {
		inst := instructions[i]
		switch inst[0] {
		case "cpy":
			val, err := strconv.Atoi(inst[1])
			if err != nil {
				values[inst[2]] = values[inst[1]]
				break
			}
			values[inst[2]] = val
		case "inc":
			values[inst[1]]++
		case "dec":
			values[inst[1]]--
		case "jnz":
			nz, err := strconv.Atoi(inst[1])
			if err != nil {
				nz = values[inst[1]]
			}
			if nz == 0 {
				break
			}
			by, err := strconv.Atoi(inst[2])
			if err != nil {
				log.Fatalf("Invalid jump amount: %q", inst[2])
			}
			i += by - 1
		default:
			log.Fatalf("Invalid instruction: %q", inst[0])
		}
		i++
	}

	return values["a"]
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

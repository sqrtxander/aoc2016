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

func getValOrLookup(num string, values map[string]int) int {
	val, err := strconv.Atoi(num)
	if err != nil {
		if !strings.Contains("abcd", num) {
			log.Fatalf("Invalid register: %q\n", num)
		}
		val = values[num]
	}
	return val
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	instructions := utils.Map(lines, strings.Fields)

	values := map[string]int{
		"a": 7,
		"b": 0,
		"c": 0,
		"d": 0,
	}

	ip := 0
	for ip < len(instructions) {
		inst := instructions[ip]
		switch inst[0] {
		case "cpy":
			val := getValOrLookup(inst[1], values)
			_, err := strconv.Atoi(inst[2])
			if err == nil {
				break
			}
			values[inst[2]] = val
		case "inc":
			values[inst[1]]++
		case "dec":
			values[inst[1]]--
		case "jnz":
			nz := getValOrLookup(inst[1], values)
			if nz == 0 {
				break
			}
			by := getValOrLookup(inst[2], values)
			ip += by - 1
		case "tgl":
			toggleIdx := getValOrLookup(inst[1], values) + ip
			if toggleIdx >= len(instructions) {
				break
			}
			toToggle := instructions[toggleIdx][0]
			if len(instructions[toggleIdx]) == 2 {
				if toToggle == "inc" {
					instructions[toggleIdx][0] = "dec"
				} else {
					instructions[toggleIdx][0] = "inc"
				}
			} else {
				if toToggle == "jnz" {
					instructions[toggleIdx][0] = "cpy"
				} else {
					instructions[toggleIdx][0] = "jnz"
				}
			}
		default:
			log.Fatalf("Invalid instruction: %q", inst[0])
		}
		ip++
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

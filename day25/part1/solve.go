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

func hashValues(values map[string]int) string {
	return fmt.Sprintf("%d;%d;%d;%d",
		values["a"],
		values["b"],
		values["c"],
		values["d"],
	)
}

func run(initA int, instructions [][]string) bool {
	values := map[string]int{
		"a": initA,
		"b": 0,
		"c": 0,
		"d": 0,
	}

	seen := map[string]int{}
	rets := []int{}

	ip := 0
	for ip >= 0 && ip < len(instructions) {
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
		case "add":
			val := getValOrLookup(inst[1], values)
			_, err := strconv.Atoi(inst[2])
			if err == nil {
				break
			}
			values[inst[2]] += val
		case "sar":
			val := getValOrLookup(inst[1], values)
			_, err := strconv.Atoi(inst[2])
			if err == nil {
				break
			}
			values[inst[2]] >>= val
		case "mod":
			val := getValOrLookup(inst[1], values)
			_, err := strconv.Atoi(inst[2])
			if err == nil {
				break
			}
			values[inst[2]] = (val-1)%2 + 1 // adjust for when 0 going to 2
		case "nop":
			break
		case "out":
			ret := getValOrLookup(inst[1], values)
			if len(rets) == 0 && ret != 0 {
				return false
			}
			if len(rets) >= 2 {
				if !(rets[len(rets)-1] == 0 && rets[len(rets)-2] == 1) &&
					!(rets[len(rets)-1] == 1 && rets[len(rets)-2] == 0) {
					return false
				}
				if val, ok := seen[hashValues(values)]; ret == 0 && rets[len(rets)-1] == 1 && val == ret && ok {
					return true
				} else if val, ok := seen[hashValues(values)]; ret == 1 && rets[len(rets)-1] == 0 && val == ret && ok {
					return true
				}
			}
			rets = append(rets, ret)
			seen[hashValues(values)] = ret
		default:
			log.Fatalf("Invalid instruction: %q", inst[0])
		}
		ip++
	}
	return false
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	lines = append([]string{
		"add 2572 a",
		"cpy a d",
	}, lines[8:]...)
	lines = append(
		lines[:4],
		append(
			append(
				[]string{
					"mod a c",
					"sar 1 a",
				},
				slices.Repeat([]string{"nop"}, 8)...,
			),
			lines[14:]...,
		)...,
	)

	instructions := utils.Map(lines, strings.Fields)

	for i := 1; ; i++ {
		if run(i, instructions) {
			return i
		}
	}
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

package main

import (
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
	repeating
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	state := normal
	result := ""
	marker := ""
	repeat := ""
	var amt, rpt int
	for _, char := range s {
		if char == '(' && state == normal {
			state = findingMarker
			continue
		} else if char == ')' && state == findingMarker {
			amtStr, rptStr, _ := strings.Cut(marker, "x")
			var err error
			amt, err = strconv.Atoi(amtStr)
			if err != nil {
				log.Fatalf("Bad amount in %q", marker)
			}
			rpt, err = strconv.Atoi(rptStr)
			if err != nil {
				log.Fatalf("Bad repeat in %q", marker)
			}
			marker = ""
			state = repeating
			continue
		}
		if state == normal {
			result += string(char)
		} else if state == findingMarker {
			marker += string(char)
		} else if state == repeating {
			repeat += string(char)
			amt--
			if amt == 0 {
				result += strings.Repeat(repeat, rpt)
				repeat = ""
				state = normal
			}
		} else {
			log.Fatalln("Unknown state")
		}
	}
	return len(result)
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

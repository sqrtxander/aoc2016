package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func solve(s string) int {
	_ = s
	a := 12
	b := a - 1
	for b > 1 {
		a *= b
		b--
	}
	a += 84 * 76
	return a // decompiled the code
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

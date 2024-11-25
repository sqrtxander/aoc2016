package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func dragonify(a string) string {
	var sb strings.Builder
	sb.WriteString(a)
	sb.WriteByte('0')
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '0' {
			sb.WriteByte('1')
		} else if a[i] == '1' {
			sb.WriteByte('0')
		} else {
			log.Fatalf("Invalid digit: %q\n", a[i])
		}
	}
	return sb.String()
}

func getChecksum(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s)-1; i += 2 {
		if s[i] == s[i+1] {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}

func solve(s string, length int) string {
	s = strings.TrimSpace(s)
	result := s
	for len(result) < length {
		result = dragonify(result)
	}
	result = result[:length]
	for len(result)%2 == 0 {
		result = getChecksum(result)
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
	fmt.Println(solve(string(contents), 272))
}

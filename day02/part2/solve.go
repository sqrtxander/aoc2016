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

	code := ""
	p := utils.Point{X: 0, Y: 2}
	keypad := [][]byte{
		{' ', ' ', '1', ' ', ' '},
		{' ', '2', '3', '4', ' '},
		{'5', '6', '7', '8', '9'},
		{' ', 'A', 'B', 'C', ' '},
		{' ', ' ', 'D', ' ', ' '},
	}
	for _, line := range lines {
		for _, dir := range line {
			switch dir {
			case 'U':
				p.Y = p.Y - 1
				if p.Y < 0 || keypad[p.Y][p.X] == ' ' {
					p.Y = p.Y + 1
				}
			case 'R':
				p.X = p.X + 1
				if p.X >= len(keypad[p.Y]) || keypad[p.Y][p.X] == ' ' {
					p.X = p.X - 1
				}
			case 'L':
				p.X = p.X - 1
				if p.X < 0 || keypad[p.Y][p.X] == ' ' {
					p.X = p.X + 1
				}
			case 'D':
				p.Y = p.Y + 1
				if p.Y >= len(keypad) || keypad[p.Y][p.X] == ' ' {
					p.Y = p.Y - 1
				}
			default:
				log.Fatalf("Unknown move: '%c'\n", dir)
			}
		}
		digit := keypad[p.Y][p.X]
		code += string(digit)
	}

	return code
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

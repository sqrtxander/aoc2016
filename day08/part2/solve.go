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

func solve(s string, width int, height int) string {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	screen := utils.BoundedHashGrid{
		Grid: utils.HashGrid{},
		W:    width,
		H:    height,
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "rect") {
			line = line[len("rect "):]
			xwStr, yhStr, _ := strings.Cut(line, "x")
			xw, err := strconv.Atoi(xwStr)
			if err != nil {
				log.Fatalf("Invalid dimensions in line %q\n", line)
			}
			yh, err := strconv.Atoi(yhStr)
			if err != nil {
				log.Fatalf("Invalid dimensions in line %q\n", line)
			}

			for x := range xw {
				for y := range yh {
					screen.Grid[utils.Point{X: x, Y: y}] = true
				}
			}
		} else if strings.HasPrefix(line, "rotate row") {
			line = line[len("rotate row y="):]
			yStr, byStr, _ := strings.Cut(line, " by ")
			by, err := strconv.Atoi(byStr)
			by %= width
			if err != nil {
				log.Fatalf("Invalid by amount in line %q\n", line)
			}
			y, err := strconv.Atoi(yStr)
			if err != nil {
				log.Fatalf("Invalid column in line %q\n", line)
			}
			store := utils.HashGrid{}
			for x := range width {
				store[utils.Point{X: x, Y: y}] = screen.Grid[utils.Point{X: x, Y: y}]
			}
			for x := range width {
				screen.Grid[utils.Point{X: x, Y: y}] = store[utils.Point{X: ((x-by)%width + width) % width, Y: y}]
			}
		} else if strings.HasPrefix(line, "rotate column") {
			line = line[len("rotate column x="):]
			xStr, byStr, _ := strings.Cut(line, " by ")
			by, err := strconv.Atoi(byStr)
			by %= height
			if err != nil {
				log.Fatalf("Invalid by amount in line %q\n", line)
			}
			x, err := strconv.Atoi(xStr)
			if err != nil {
				log.Fatalf("Invalid column in line %q\n", line)
			}
			store := utils.HashGrid{}
			for y := range height {
				store[utils.Point{X: x, Y: y}] = screen.Grid[utils.Point{X: x, Y: y}]
			}
			for y := range height {
				screen.Grid[utils.Point{X: x, Y: y}] = store[utils.Point{X: x, Y: ((y-by)%height + height) % height}]
			}
		} else {
			log.Fatalf("Invalid instruction: %q", line)
		}
	}

	return screen.GetBoundedHash()
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
	fmt.Println(solve(string(contents), 50, 6))
}

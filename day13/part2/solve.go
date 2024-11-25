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

func isOpen(p utils.Point, inp int) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}
	val := (p.X+p.Y)*(p.X+p.Y) + 3*p.X + p.Y + inp
	result := true
	for val > 0 {
		if val%2 == 1 {
			result = !result
		}
		val /= 2
	}
	return result
}

func solve(s string, maxDist int) int {
	inp, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatalln(err)
	}

	p := utils.Point{X: 1, Y: 1}
	dists := map[utils.Point]int{
		p: 0,
	}
	queue := utils.Queue[utils.Point]{p}
	for {
		queue, p = queue.Pop()
		if dists[p] >= maxDist {
			break
		}
		for _, q := range utils.Adjacent4(p) {
			if _, contains := dists[q]; !contains && isOpen(q, inp) {
				queue = queue.Push(q)
				dists[q] = dists[p] + 1
			}
		}
	}

	return len(dists)
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
	fmt.Println(solve(string(contents), 50))
}

package main

import (
	"aoc2016/utils"
	"fmt"
	"log"
	"maps"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

type path struct {
	loc     utils.Point
	visited map[utils.Point]bool
	dist    int
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	pattern := regexp.MustCompile(`\d`)
	numberless := pattern.ReplaceAllString(s, ".")

	grid := utils.GetHashGrid(numberless, '#', '.')
	numbers := map[utils.Point]bool{}
	dists := map[utils.Point]map[utils.Point]int{}
	var start utils.Point
	for y, line := range lines {
		for x, ch := range line {
			if ch >= '0' && ch <= '9' {
				numbers[utils.Point{X: x, Y: y}] = true
			}
			if ch == '0' {
				start = utils.Point{X: x, Y: y}
			}
		}
	}
	for n := range numbers {
		dists[n] = map[utils.Point]int{}
	}
	for p := range numbers {
		ds := map[utils.Point]int{}
		ds[p] = 0
		queue := utils.Queue[utils.Point]{p}
		for len(dists[p]) < len(numbers)-1 && len(queue) > 0 {
			var q utils.Point
			queue, q = queue.Pop()
			if _, ok := numbers[q]; ok && p != q {
				dists[q][p] = ds[q]
				dists[p][q] = ds[q]
			}
			for _, r := range utils.Adjacent4(q) {
				if !grid[r] {
					continue
				}
				if _, ok := ds[r]; !ok {
					queue = queue.Push(r)
					ds[r] = ds[q] + 1
				}
			}
		}
	}

	init := path{
		loc:     start,
		visited: map[utils.Point]bool{start: true},
		dist:    0,
	}
	queue := utils.Queue[path]{init}
	shortest := math.MaxInt
	for len(queue) > 0 {
		var p path
		queue, p = queue.Pop()
		for nextLoc := range numbers {
			if !p.visited[nextLoc] {
				vis := maps.Clone(p.visited)
				vis[nextLoc] = true
				d := p.dist + dists[p.loc][nextLoc]
				if len(p.visited) == len(numbers)-1 {
					shortest = min(shortest, d)
				}
				queue = queue.Push(path{
					loc:     nextLoc,
					visited: vis,
					dist:    d,
				})
			}
		}
	}

	return shortest
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

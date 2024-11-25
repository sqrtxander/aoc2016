package main

import (
	"aoc2016/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type path struct {
	path string
	pos  utils.Point
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func isValid(inp string, path path, move path) bool {
	newPos := utils.Add(path.pos, move.pos)
	if newPos.X < 0 || newPos.X >= 4 || newPos.Y < 0 || newPos.Y >= 4 {
		return false
	}
	idxMap := map[string]int{
		"U": 0,
		"D": 1,
		"L": 2,
		"R": 3,
	}
	hash := getMD5Hash(inp + path.path)
	char := hash[idxMap[move.path]]
	if (char >= '0' && char <= '9') || char == 'a' {
		return false
	} else if char >= 'b' && char <= 'f' {
		return true
	}
	log.Fatalf("Invalid hash character '%c'\n", char)
	return false
}

func solve(s string) string {
	s = strings.TrimSpace(s)

	init := path{
		path: "",
		pos:  utils.Point{X: 0, Y: 0},
	}
	moves := []path{
		{
			path: "U",
			pos:  utils.Point{X: 0, Y: -1},
		},
		{
			path: "R",
			pos:  utils.Point{X: 1, Y: 0},
		},
		{
			path: "D",
			pos:  utils.Point{X: 0, Y: 1},
		},
		{
			path: "L",
			pos:  utils.Point{X: -1, Y: 0},
		},
	}

	queue := utils.Queue[path]{init}
	for len(queue) > 0 {
		var curr path
		queue, curr = queue.Pop()
		if curr.pos.X == 3 && curr.pos.Y == 3 {
			return curr.path
		}
		for _, move := range moves {
			if isValid(s, curr, move) {
				newPath := path{
					path: curr.path + move.path,
					pos:  utils.Add(curr.pos, move.pos),
				}
				queue = queue.Push(newPath)
			}
		}
	}

	return ""
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

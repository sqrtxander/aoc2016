package main

import (
	"aoc2016/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

type key struct {
	i         int
	quintuple string
	found     bool
}

func getStretchedMD5Hash(hashStr string) string {
	for range 2017 {
		hash := md5.Sum([]byte(hashStr))
		hashStr = hex.EncodeToString(hash[:])
	}
	return hashStr
}

func hasTriple(hash string) (string, bool) {
	for i := range len(hash) - 2 {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return strings.Repeat(string(hash[i]), 5), true
		}
	}
	return "", false
}

func solve(s string) int {
	s = strings.TrimSpace(s)

	potentialKeys := []key{}
	keyIdxs := []int{}
	idxCap := math.MaxInt
	for i := 0; i < idxCap; i++ {
		toHash := s + strconv.Itoa(i)
		hash := getStretchedMD5Hash(toHash)
		// check if a key has been found
		for j, k := range potentialKeys {
			if !strings.Contains(hash, k.quintuple) {
				continue
			}
			keyIdxs = append(keyIdxs, k.i)
			potentialKeys[j].found = true
			if len(keyIdxs) == 64 {
				idxCap = i + 1000
			}
		}

		// remove expired potential keys
		potentialKeys = utils.Filter(potentialKeys, func(k key) bool {
			return i-k.i <= 1000 && !k.found
		})

		// add potential keys
		if val, has := hasTriple(hash); has {
			potentialKeys = append(potentialKeys, key{
				i:         i,
				quintuple: val,
				found:     false,
			})
		}
	}
	slices.Sort(keyIdxs)
	return keyIdxs[63]
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

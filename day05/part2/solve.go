package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func solve(s string) string {
	s = strings.TrimSpace(s)
	i := 0
	done := 0
	password := [8]byte{}
	for done < 8 {
		toHash := s + strconv.Itoa(i)
		hash := getMD5Hash(toHash)
		if strings.HasPrefix(hash, "00000") && '0' <= hash[5] && hash[5] <= '7' {
			idx, _ := strconv.Atoi(string(hash[5]))
			if password[idx] == 0 {
				password[idx] = hash[6]
				done++
			}
		}
		i++
	}
	return string(password[:])
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

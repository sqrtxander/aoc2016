package utils

import "strconv"

type Pair[T, U any] struct {
	K T
	V U
}

func CastToInts(strs []string) ([]int, error) {
	ints := []int{}
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func Sum(is ...int) int {
	result := 0
	for _, i := range is {
		result += i
	}
	return result
}

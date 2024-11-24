package utils

import "strconv"

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

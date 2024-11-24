package utils

func Filter[T any](slice []T, test func(T) bool) (result []T) {
	for _, s := range slice {
		if test(s) {
			result = append(result, s)
		}
	}
	return
}

func Map[T any, S any](slice []T, change func(T) S) (result []S) {
    for _, s := range slice {
        result = append(result, change(s))
        }
    return
}

func Any[T any](slice []T, test func(T) bool) bool {
    for _, s := range slice {
        if test(s) {
            return true
        }
    }
    return false
}

func All[T any](slice []T, test func(T) bool) bool {
    for _, s := range slice {
        if !test(s) {
            return false
        }
    }
    return true
}

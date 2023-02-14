package search

import "errors"

// SequentialSearch searches for target in given list and returns its position if it is presented there. Returns -1 if not.
func SequentialSearch(arr []int, target int) (int, error) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i, nil
		}
	}

	return -1, errors.New("target value is not present in the list")
}
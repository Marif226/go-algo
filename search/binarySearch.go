package search

import "errors"

// BinarySearch searches for target in given list and returns its position if it is presented there. Returns -1 if not.
func BinarySearch(arr []int, target int) (int, error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		switch {
			case arr[mid] == target:
				return mid, nil
			case arr[mid] > target:
				high = mid - 1
				continue
			case arr[mid] < target: 
				low = mid + 1
				continue
			default:
				continue
		}
	}

	return -1, errors.New("target value is not present in the list")
}
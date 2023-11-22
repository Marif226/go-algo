package sort

import "errors"

func SelectionSort(arr []int) error {
	if len(arr) == 0 {
		return errors.New("array is empty")
	}

	
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return nil
}
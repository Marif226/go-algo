package sort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]

	less := make([]int, 0)
	greater := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}

	res := append(QuickSort(less), pivot)
	res = append(res, QuickSort(greater)...)

	return res
}
package binarysearch

func BBinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}

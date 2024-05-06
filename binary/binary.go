package binary

// BinarySearch performs a binary search on a sorted array of ints
func BinarySearch(data []int, target int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + (high-low)/2 // Prevents overflow
		if data[mid] < target {
			low = mid + 1
		} else if data[mid] > target {
			high = mid - 1
		} else {
			return mid // Target found
		}
	}

	return -1 // Target not found
}

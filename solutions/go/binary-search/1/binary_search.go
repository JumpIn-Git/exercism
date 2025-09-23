package binarysearch

func SearchInts(list []int, key int) int {
	start, end := 0, len(list)-1
	for start <= end {
		middle := (start + end) / 2
		if list[middle] == key {
			return middle
		}
		if list[middle] > key {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return -1
}

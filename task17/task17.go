package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, target int) int {
	low, high, mid := 0, len(arr)-1, 0

	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 4
	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= target })
	if index < len(arr) && arr[index] == target {
		fmt.Println(index)
	} else {
		fmt.Println(-1)
	}

	index = binarySearch(arr, target)
	fmt.Println(index)
}

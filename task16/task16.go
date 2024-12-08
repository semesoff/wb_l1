package main

import (
	"fmt"
	"sort"
)

func task16_1() {
	arr := []int{5, 4, 3, 2, 1}
	sort.Ints(arr)
	fmt.Println(arr)

	arr1 := []int{5, 4, 3, 2, 1}
	quickSort(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // Последний элемент в качестве опорного
	i := low - 1       // Индекс меньшего элемента

	for j := low; j < high; j++ {
		if arr[j] < pivot { // Если текущий элемент меньше опорного
			i++                             // Увеличиваем индекс меньшего элемента
			arr[i], arr[j] = arr[j], arr[i] // Меняем местами
		}
	}
	// Меняем местами опорный элемент и элемент, находящийся на позиции i+1
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // Возвращаем индекс опорного элемента pivot
}

func main() {
	task16_1()
}

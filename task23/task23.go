package main

import (
	"fmt"
)

/*
Удалить i-ый элемент из слайса.
*/

func remove[T any](slice *[]T, i int) {
	// Удаляем i-ый элемент из слайса
	*slice = append((*slice)[:i], (*slice)[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice)

	i := 2
	remove(&slice, i)
	fmt.Println("Slice after removal:", slice)
}

package main

import (
	"fmt"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func task11() []int {
	arr1 := [5]int{1, 3, 5, 7, 9}
	arr2 := [5]int{2, 3, 4, 5, 6}

	// Создание хеш-таблицы для обнаружения пересечений
	hash := make(map[int]bool)

	// Заполнение хеш-таблицы значениями из первого массива
	for _, v := range arr1 {
		hash[v] = true
	}

	res := make([]int, 0)
	// Проверка наличия значений из второго массива в хеш-таблице
	for _, v := range arr2 {
		if _, ok := hash[v]; ok {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	fmt.Println(task11())
}

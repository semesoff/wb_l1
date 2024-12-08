package main

import (
	"fmt"
)

func task12() {
	strings := [5]string{"cat", "cat", "dog", "cat", "tree"}

	hash := make(map[string]bool)

	// Добавление элементов в хэш-таблицу
	for _, v := range strings {
		hash[v] = true
	}

	// Вывод ключей хэш-таблицы
	for k := range hash {
		fmt.Println(k)
	}
}

func main() {
	task12()
}

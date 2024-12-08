package main

import (
	"fmt"
)

func task19(line string) string {
	// Используем массив рун т.к. они могут хранить символы Юникода
	runes := []rune(line)
	// i это индекс первого элемента, j это индекс последнего элемента
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// меняем местами элементы с индексами i и j
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(task19("Hello, World! 🌍✨🚀"))
}

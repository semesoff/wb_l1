package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func task20(input string) []string {
	// разбиваем строку на слова по пробелу
	words := strings.Split(input, " ")
	// i - индекс первого слова, j - индекс последнего слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		// меняем местами слова
		words[i], words[j] = words[j], words[i]
	}
	return words
}

func main() {
	fmt.Println(task20("snow dog sun"))
}

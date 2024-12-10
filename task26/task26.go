package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func isUnique(s string) bool {
	// преобразуем строку в нижний регистр для регистронезависимой проверки
	s = strings.ToLower(s)
	// Используем hash-карту для отслеживания встреченных символов
	hash := make(map[rune]bool)

	for _, c := range s {
		// если такой же символ уже есть в hash, значит он повторяется
		if hash[c] {
			return false // false - строка не уникальна
		}
		// добавление символа в hash
		hash[c] = true
	}
	return true // true - строка уникальна
}

func main() {
	arr := [3]string{"abcd", "abCdefAaf", "aabcd"}
	for _, s := range arr {
		fmt.Printf("%s - %t\n", s, isUnique(s))
	}
}

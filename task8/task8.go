package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(num int64, pos uint, value int) int64 {
	if value == 1 {
		// Устанавливаем i-й бит в 1
		return num | (1 << pos)
	} else {
		// Устанавливаем i-й бит в 0
		return num &^ (1 << pos) // &^ это "очистка бита"
	}
}

func main() {
	var number int64 = 52 // Число, которое будем менять (101010 в двоичной системе)
	position := uint(3)   // Позиция бита (нумерация с 0)
	value := 1            // Значение, которое нужно установить: 1 или 0

	fmt.Printf("Исходное число: %b\n", number)
	number = setBit(number, position, value)
	fmt.Printf("Результат после устновки бита в 1: %b\n", number)

	value = 0 // Меняем значение на 0
	number = setBit(number, position, value)
	fmt.Printf("Результат после установки бита в 0: %b\n", number)
}

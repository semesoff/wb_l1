package main

import (
	"fmt"
)

/*
Поменять местами два числа без создания временной переменной.
*/

// Решение с свапом переменных
func task13_1() {
	x := 5
	y := 10
	x, y = y, x
	fmt.Println(x, y)
}

// Решение с арифметическими операциями - +
func task13_2() {
	x := 5
	y := 10
	x = x + y
	y = x - y
	x = x - y
	fmt.Println(x, y)
}

// Решение с арифметическими операциями * /
func task13_3() {
	x := 5
	y := 10
	x = x * y
	y = x / y
	x = x / y
	fmt.Println(x, y)
}

// Решение через побитовые операции XOR
func task13_4() {
	x := 5
	y := 10
	x = x ^ y
	y = x ^ y
	x = x ^ y
	fmt.Println(x, y)
}

func main() {
	task13_1()
	task13_2()
	task13_3()
	task13_4()
}

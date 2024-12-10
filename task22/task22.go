package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, 
вычитает две числовых переменных a и b, значение которых > 2^20.
*/

func main() {
	// Создаю два числа a и b, каждое из которых равно 2^30
	a := new(big.Int).Lsh(big.NewInt(1), 30)
	b := new(big.Int).Lsh(big.NewInt(1), 30)

	// Создаю переменные для хранения результата умножения, деления, сложения и вычитания
	var mul, div, add, sub big.Int
	mul.Mul(a, b)
	div.Div(a, b)
	add.Add(a, b)
	sub.Sub(a, b)

	fmt.Println("Multiplication: ", &mul)
	fmt.Println("Division: ", &div)
	fmt.Println("Addition: ", &add)
	fmt.Println("Subtraction: ", &sub)
}

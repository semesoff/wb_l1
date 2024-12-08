package main

import (
	"fmt"
	"reflect"
)

// Решение с использованием switch
func task14_1(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Println("unknown")
	}
}

// Решение с использованием reflect.TypeOf
func task14_2(variable interface{}) {
	fmt.Println(reflect.TypeOf(variable))
}

// Решение с использованием %T
func task14_3(variable interface{}) {
	fmt.Printf("%T\n", variable)
}

func main() {
	task14_1(5)
	task14_2("Hello")
	task14_3(true)
	task14_1(make(chan int))
}

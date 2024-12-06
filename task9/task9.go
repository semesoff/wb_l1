package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	chx := make(chan int)
	chx2 := make(chan int)

	// заполняем канал значениями из массива
	go func() {
		for _, i := range arr {
			chx <- i
		}
		close(chx) // закрываем канал после отправки всех значений
	}()

	// возводим в квадрат значения из канала и отправляем во второй канал
	go func() {
		for i := range chx {
			chx2 <- i * i
		}
		close(chx2)
	}()

	// выводим значения из второго канала
	for i := range chx2 {
		fmt.Println(i)
	}
}

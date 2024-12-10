package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение
 квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

// Реализация с использованием каналов
func solution1(numbers [5]int) {
	c := make(chan int, len(numbers))

	// Запускаем горутины для каждого числа
	for _, i := range numbers {
		go func(i int) {
			c <- i * i
		}(i)
	}

	// Выводим результаты
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-c)
	}
}

// Реализация с использованием sync.WaitGroup
func solution2(numbers [5]int) {
	wg := new(sync.WaitGroup)

	// Запускаем горутины для каждого числа
	for _, i := range numbers {
		// Увеличиваем счетчик горутин
		wg.Add(1)
		go func(i int) {
			defer wg.Done()    // Уменьшаем счетчик горутин
			fmt.Println(i * i) // Выводим результат
		}(i)
	}
	wg.Wait() // Ожидаем завершения всех горутин
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	solution1(numbers)
}

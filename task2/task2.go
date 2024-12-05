package main

import (
	"fmt"
	"sync"
)

// Реализация с использованием каналов
func solution1(numbers [5]int) {
	c := make(chan int, len(numbers))

	for _, i := range numbers {
		go func(i int) {
			c <- i * i
		}(i)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-c)
	}
}

// Реализация с использованием sync.WaitGroup
func solution2(numbers [5]int) {
	wg := new(sync.WaitGroup)

	for _, i := range numbers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}
	wg.Wait()
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	solution1(numbers)
}

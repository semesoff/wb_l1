package main

import (
	"fmt"
	"sync"
)

func solution1(numbers [5]int) int {
	c := make(chan int, len(numbers))

	for _, i := range numbers {
		go func(i int) {
			c <- i * i
		}(i)
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-c
	}

	return sum
}

func solution2(numbers [5]int) int {
	sum := 0

	wg := new(sync.WaitGroup)

	for _, i := range numbers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sum += i * i
		}(i)
	}

	wg.Wait()
	return sum
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	fmt.Println(solution1(numbers))
}

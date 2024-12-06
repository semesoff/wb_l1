package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать конкурентную запись данных в map.
*/

func writer(mu *sync.Mutex, hash *map[int]int, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done(): // Проверяем контекст на завершение
			wg.Done() // Уменьшаем счетчик горутин
			fmt.Println("Writer stopped.")
			return
		default:
			mu.Lock() // Блокируем доступ к map
			number := rand.Intn(100)
			(*hash)[number] = number
			mu.Unlock() // Разблокируем доступ к map
			fmt.Printf("Writer sent data: [%d]: [%d]\n", number, number)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	mu := new(sync.Mutex)
	hash := make(map[int]int)
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())
	N := 5

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt) // Перехватываем сигналы на завершение программы

	go func() {
		<-sigChan
		fmt.Println("Received signal to stop.")
		cancel()
	}()

	for i := 0; i < N; i++ {
		wg.Add(1)
		go writer(mu, &hash, ctx, wg)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("All writers stopped.")
}

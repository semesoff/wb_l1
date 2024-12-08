package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// CounterProvider Интерфейс для счетчика
type CounterProvider interface {
	Increment()
}

// Counter Счетчик
type Counter struct {
	count int
	mu    sync.Mutex
}

// Increment Инкремент счетчика
func (c *Counter) Increment() {
	c.count++
}

// worker Воркер, который инкрементирует счетчик
func worker(c CounterProvider, ctx context.Context, id int, wg *sync.WaitGroup, mu *sync.Mutex) {
	fmt.Printf("Worker started [%d]\n", id)
	for {
		select {
		case <-ctx.Done(): // Сигнал завершения работы
			wg.Done() // Уменьшаем счетчик воркеров
			fmt.Printf("Worker stopped [%d]\n", id)
			return
		default:
			mu.Lock()     // Блокируем доступ к счетчику
			c.Increment() // Инкрементируем счетчик
			fmt.Printf("Worker [%d] incremented counter.\n", id)
			mu.Unlock() // Разблокируем доступ к счетчику
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	n := 3
	counter := &Counter{}

	wg := new(sync.WaitGroup)
	sigChan := make(chan os.Signal, 1)                      // Канал для сигналов
	ctx, cancel := context.WithCancel(context.Background()) // Контекст для завершения работы
	signal.Notify(sigChan, os.Interrupt)                    // Подписываемся на сигналы завершения

	go func() {
		<-sigChan
		cancel()
	}()

	// Запускаем воркеров
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(counter, ctx, i, wg, &(counter.mu))
	}

	wg.Wait()
	fmt.Println("All worker stopped.")
	fmt.Printf("Counter value: %d\n", counter.count)
}

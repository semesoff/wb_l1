package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// StopGoroutineWithChannel Решение с каналом
func StopGoroutineWithChannel(ch <-chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Goroutine stopped")
			return
		default:
			fmt.Println("Goroutine is working")
			time.Sleep(1 * time.Second)
		}
	}
}

// StopGoroutineWithContext Решение с контекстом
func StopGoroutineWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped")
			return
		default:
			fmt.Println("Goroutine is working")
			time.Sleep(1 * time.Second)
		}
	}
}

// StopGoroutineWithAfter Решение с After
func StopGoroutineWithAfter() {
	ch := time.After(1 * time.Second)
	for {
		select {
		case <-ch:
			fmt.Println("Goroutine stopped")
			return
		default:
			fmt.Println("Goroutine is working")
			time.Sleep(1 * time.Second)
		}
	}
}

// StopGoroutineWithOnce Решение с sync.Once
func StopGoroutineWithOnce() {
	var once sync.Once
	ch := make(chan struct{})

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Goroutine stopped")
				return
			default:
				fmt.Println("Goroutine is working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	once.Do(func() {
		close(ch)
	})
}

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go StopGoroutineWithChannel(ch)
	go StopGoroutineWithContext(ctx)
	go StopGoroutineWithAfter()
	go StopGoroutineWithOnce()

	ch <- struct{}{}
	cancel()
	time.Sleep(1 * time.Second)
}

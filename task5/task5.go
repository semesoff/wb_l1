package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sender(ctx context.Context, ch chan<- int) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Sender stopped.")
			close(ch)
			return
		default:
			ch <- i
			i++
			fmt.Println("Sent data:", i)
			time.Sleep(time.Second)
		}
	}
}

func receiver(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Receiver stopped.")
			return
		case data := <-ch:
			fmt.Println("Received data:", data)
		}
	}
}

func main() {
	N := 10
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer cancel()

	wg.Add(1)
	go receiver(ctx, wg, ch)
	go sender(ctx, ch)

	wg.Wait()
	fmt.Println("Program finished.")
}

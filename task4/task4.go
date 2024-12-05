package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
*/

func task4(countWorkers int) {
	mainChannel := make(chan string)
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем воркеров
	for i := 0; i < countWorkers; i++ {
		wg.Add(1)
		go worker(mainChannel, wg, ctx, i)
	}

	// Хендл сигнала для завершения работы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		fmt.Println("Received signal to stop.")
		cancel()           // Говорим всем воркерам заканчивать работу
		close(mainChannel) // Закрываем канал, чтобы предотвратить дальнейшую запись в него
	}()

	// Записываем данные в канал
	i := 0
	for {
		select {
		case <-ctx.Done(): // Сигнал завершения работы
			wg.Wait()
			fmt.Println("All workers stopped.")
			return
		default:
			mainChannel <- fmt.Sprintf("Data %d", i)
			fmt.Printf("Main thread sent data: Data %d\n", i)
			i++
			time.Sleep(time.Second) // Чтобы не заспамить stdout
		}
	}
}

func worker(ch <-chan string, wg *sync.WaitGroup, ctx context.Context, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // Сигнал завершения работы
			fmt.Printf("Worker [%d] stopped.\n", id)
			return
		case data := <-ch: // Получаем данные из канала
			fmt.Printf("Worker [%d] received data: %s\n", id, data)
		}
	}
}

func main() {
	numWorkers := 5
	task4(numWorkers)
}

package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func sleep(duration time.Duration) {
	// Ждем указанное время duration, чтобы получить данные из канала
	<-time.After(duration)
}

func main() {
	n := 2
	fmt.Printf("Sleeping for %d seconds...\n", n)
	sleep(time.Duration(n) * time.Second)
	fmt.Println("Awake!")
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Неверное число воркеров")
		return
	}

	dataCh := make(chan int)

	// Запускаем N воркеров
	for i := 1; i <= n; i++ {
		go worker(i, dataCh)
	}

	// Главная горутина пишет в канал бесконечно
	for i := 1; ; i++ {
		dataCh <- i
		time.Sleep(300 * time.Millisecond)
	}
}

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d получил: %d\n", id, job)
	}
}

/*
L1.3
Работа нескольких воркеров:

Реализовать постоянную запись данных в канал (в главной горутине).
Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
*/

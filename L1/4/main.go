package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	// создаём контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// ловим SIGINT (Ctrl+C)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// отдельная горутина ждёт сигнал
	go func() {
		<-sigs
		fmt.Println("\nПолучен сигнал завершения...")
		cancel() // отменяем контекст
	}()

	dataCh := make(chan int)

	// Запускаем N воркеров
	for i := 1; i <= n; i++ {
		go worker(ctx, i, dataCh)
	}

	// Главная горутина пишет в канал пока не отменён контекст
	i := 1
	for {
		select {
		case <-ctx.Done():
			close(dataCh)
			fmt.Println("Главная горутина завершена")
			return
		default:
			dataCh <- i
			i++
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func worker(ctx context.Context, id int, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d завершает работу\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: канал закрыт, выхожу\n", id)
				return
			}
			fmt.Printf("Worker %d получил: %d\n", id, job)
		}
	}
}

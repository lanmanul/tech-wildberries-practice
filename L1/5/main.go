package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan int)
	// Горутина — отправляет числа в канал
	go func() {
		i := 1
		for {
			out <- i
			i++
			time.Sleep(300 * time.Millisecond) // задержка для наглядности
		}
	}()

	// Таймер на 5 секунд
	timeout := time.After(5 * time.Second)

	for {
		select {
		case v := <-out:
			fmt.Println("Получено:", v)
		case <-timeout:
			fmt.Println("Время вышло, завершаем.")
			return
		}
	}
}

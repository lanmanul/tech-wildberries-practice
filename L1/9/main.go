package main

import (
	"fmt"
)

func main() {
	// Исходный массив чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Два канала
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Генерация чисел
	go func() {
		for _, n := range numbers {
			ch1 <- n
		}
		close(ch1) // закрываем первый канал после записи всех чисел
	}()

	// Обработка чисел (x * 2)
	go func() {
		for n := range ch1 { // читаем пока ch1 открыт
			ch2 <- n * 2
		}
		close(ch2) // закрываем второй канал после обработки всех чисел
	}()

	// Вывод результата
	for res := range ch2 {
		fmt.Println(res)
	}
}

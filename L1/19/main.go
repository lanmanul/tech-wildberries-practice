package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите строку: ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("Ошибка: пустая строка")
		return
	}

	fmt.Println("Перевёрнутая строка через билдер:", revStrBuilder(input))
	fmt.Println("Перевёрнутая строка на месте:", revStrSwap(input))
}

/*
Вариант через билдер
*/
func revStrBuilder(s string) string {
	// Преобразуем строку в срез рун для корректной работы с Unicode
	runes := []rune(s)
	// Инициализируем Builder для сборки строки
	var builder strings.Builder
	// Заранее выделяем нужное количество памяти
	builder.Grow(len(runes))
	// Идём с конца к началу и добавляем символы в Builder
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	// Возвращаем собранную строку
	return builder.String()
}

/*
Вариант через Two-pointer
*/
func revStrSwap(s string) string {
	// Заводим срез рун
	runes := []rune(s)
	// Меняем порядок элементов на противоположный
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

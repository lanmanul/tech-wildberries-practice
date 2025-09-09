package main

import (
	"fmt"
)

func main() {
	var num int64
	var bitVal int

	// Просим ввести число
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	// Просим ввести значение для 1-го бита (0 или 1)
	for {
		fmt.Print("Введите новое значение для 1-го бита (0 или 1): ")
		fmt.Scan(&bitVal)

		if bitVal == 0 || bitVal == 1 {
			break
		}
		fmt.Println("Ошибка: нужно ввести только 0 или 1.")
	}

	fmt.Printf("Исходное число: %d (%04b)\n", num, num)

	if bitVal == 0 {
		// Сбрасываем младший бит
		num = num &^ 1
	} else {
		// Устанавливаем младший бит
		num = num | 1
	}

	fmt.Printf("Результат: %d (%04b)\n", num, num)
}

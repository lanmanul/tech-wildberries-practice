package main

import (
	"fmt"
)

// setBit устанавливает i-й бит числа n в значение bit (0 или 1)
func setBit(n int64, i uint, bit int) int64 {
	if bit == 1 {
		// Устанавливаем i-й бит в 1 (операция OR)
		n |= 1 << (i - 1)
	} else {
		// Сбрасываем i-й бит в 0 (операция AND NOT)
		n &^= 1 << (i - 1)
	}
	return n
}

func main() {
	var num int64 = 5 // 0101₂

	fmt.Printf("Исходное число: %d (%04b)\n", num, num)

	// Установим 1-й бит в 0
	res1 := setBit(num, 1, 0)
	fmt.Printf("Установка 1-го бита в 0: %d (%04b)\n", res1, res1)

	// Установим 2-й бит в 1
	res2 := setBit(num, 2, 1)
	fmt.Printf("Установка 2-го бита в 1: %d (%04b)\n", res2, res2)
}

package main

import "fmt"

func main() {
	a, b := 5, 8
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Обмен без временной переменной
	a = a + b
	b = a - b
	a = a - b

	// XOR-обмен
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}

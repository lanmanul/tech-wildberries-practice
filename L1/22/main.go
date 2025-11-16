package main

import (
	"fmt"
	"math/big"
)

func main() {
	// заводим big.Int
	a := new(big.Int)
	b := new(big.Int)

	// Присваиваем ОЧЕНЬ большие значения > 2^20
	a.SetString("123456789012345678901234567890", 10)
	b.SetString("98765432109876543210987654321", 10)

	// операции
	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	// вывод
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("Сложение =", sum)
	fmt.Println("Вычитание =", sub)
	fmt.Println("Умножение =", mul)
	fmt.Println("Деление =", div)
}

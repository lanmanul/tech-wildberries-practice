package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50} // Заводим слайс

	fmt.Println("До :", s)
	s = removeAtIndex(s, 2) // Удаляем элемент с индексом 2 это 30
	fmt.Println("После:", s)
}

func removeAtIndex(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice // Индекс вне диапазона — возвращаем как есть
	}

	// Сдвигаем хвост влево на 1
	copy(slice[i:], slice[i+1:])

	// Уменьшаем длину слайса на 1
	slice = slice[:len(slice)-1]

	return slice
}

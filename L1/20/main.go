package main

import "fmt"

func main() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	if len(input) == 0 {
		fmt.Println("Ошибка: пустая строка")
		return
	}

	result := reverseWords(input)
	fmt.Println("Результат:", result)
}

func reverseWords(s string) string {
	runes := []rune(s)

	// 1. Переворачиваем всю строку
	reverse(runes, 0, len(runes)-1)

	// 2. Переворачиваем каждое слово отдельно
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func reverse(r []rune, left, right int) {
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
}

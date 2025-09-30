package main

import (
	"fmt"
	"sort"
)

func main() {
	// Создаем слайс слов с повторами, где "cat" встречается несколько раз
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаём map для хранения уникальных строк, используем пустую структуру потому что в отличие от bool
	// она не занимает память, мы используем её просто для проверки наличия ключа
	set := make(map[string]struct{})

	// Проходим по всем словам в исходном слайсе
	for _, word := range words {
		// Записываем слово в map как ключ, где значение всегда одно и то же (struct{}{}), а если слово уже есть,
		// то оно не перезаписывается - просто игнорируется, что и даёт уникальность
		set[word] = struct{}{}
	}

	// Переносим ключи в слайс для сортировки
	var keys []string
	for k := range set {
		keys = append(keys, k)
	}

	// Сортируем ключи по алфавиту
	sort.Strings(keys)

	// Красивый вывод в стиле множества
	fmt.Print("Множество строк: {")
	for i, k := range keys {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(k)
	}
	fmt.Println("}")
}

//Вариант без сортировки и выделения дополнительной памяти под ключи, при котором порядок не гарантируется
//	words := []string{"cat", "cat", "dog", "cat", "tree"}
//
//	set := make(map[string]struct{})
//
//	for _, word := range words {
//		set[word] = struct{}{}
//	}
//
//	fmt.Println("Множество строк:")
//	for key := range set {
//		// Ключ — это уникальное слово
//		fmt.Println(key)
//	}

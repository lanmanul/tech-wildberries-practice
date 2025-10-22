package main

import "fmt"

// quickSort — рекурсивная реализация быстрой сортировки
func quickSort(arr []int) []int {
	// Обрабатываем базовый случай, если длина массива ≤ 1 - он уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Выбираем опорный элемент (pivot)
	// Можно выбрать первый, последний или середину
	pivot := arr[len(arr)/2]

	// Создаём три подмассива:
	// меньше опорного, равные опорному, больше опорного
	var left, middle, right []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}

	// Рекурсивно сортируем левую и правую части и объединяем результат
	return append(append(quickSort(left), middle...), quickSort(right)...)
}

func main() {
	arr := []int{10, 5, 2, 3, 7, 6, 8, 1, 9, 4}

	fmt.Println("Исходный массив:", arr)

	sorted := quickSort(arr)

	fmt.Println("Отсортированный массив:", sorted)
}

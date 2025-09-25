package main

import "fmt"

// Универсальная функция пересечения через generics
func intersection[T comparable](a, b []T) []T {
	set := make(map[T]bool)
	for _, v := range a {
		set[v] = true
	}

	var result []T
	for _, v := range b {
		if set[v] {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Пример с int
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	interInts := intersection(A, B)
	if len(interInts) == 0 {
		fmt.Println("Пересечений нет (int)")
	} else {
		fmt.Println("Пересечение (int):", interInts)
	}
	// Пример с float
	C := []float64{5.5, 6.3, 7.2, 9.0}
	D := []float64{6.3, 7.2, 8.2, 9.0}

	interFloats := intersection(C, D)
	if len(interFloats) == 0 {
		fmt.Println("Пересечений нет (float)")
	} else {
		fmt.Println("Пересечение (float):", interFloats)
	}

	// Пример со string
	X := []string{"apple", "banana", "cherry"}
	Y := []string{"banana", "cherry", "orange"}

	interStrings := intersection(X, Y)
	if len(interStrings) == 0 {
		fmt.Println("Пересечений нет (string)")
	} else {
		fmt.Println("Пересечение (string):", interStrings)
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {

	//Задаём массив температур
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, t := range temps {
		//Для каждой температуры определяем ключ группы
		key := int(math.Floor(t/10.0) * 10)
		//Добавляем температуру в groups
		groups[key] = append(groups[key], t)
	}

	// Выводим результат
	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, v)
	}
}

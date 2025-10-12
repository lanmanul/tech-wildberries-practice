package main

import "fmt"

// Функция принимает любое значение (interface{})
func detectType(x interface{}) {
	// Type switch — проверка конкретного типа значения
	switch v := x.(type) {
	case int:
		fmt.Printf("Тип: int, значение: %d\n", v)
	case string:
		fmt.Printf("Тип: string, значение: %q\n", v)
	case bool:
		fmt.Printf("Тип: bool, значение: %t\n", v)
	case chan int:
		fmt.Println("Тип: chan int, значение:", v)
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	detectType(42)             // int
	detectType("hello")        // string
	detectType(true)           // bool
	detectType(make(chan int)) // chan int
	detectType(3.14)           // неизвестный тип (float64)
}

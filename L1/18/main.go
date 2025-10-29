package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Заводим структуру Counter, содержащую значение счётчика
type Counter struct {
	value int64
}

func main() {
	var wg sync.WaitGroup // Используется, чтобы дождаться завершения всех горутин
	counter := Counter{}  // Заводим экземпляр счётчика с начальным значением 0

	// Запускаем 100 горутин
	for i := 0; i < 100; i++ {
		wg.Add(1) // Увеличиваем счётчик ожидания WaitGroup
		go func() {
			defer wg.Done() // После завершения горутины уменьшаем счётчик
			// Каждая горутина увеличивает счётчик 1000 раз
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait() // Ждём завершения всех горутин
	// Выводим итоговое значение (100 * 1000 = 100000)
	fmt.Println("Итоговое значение счётчика:", counter.Value())
}

// Функция Increment безопасно увеличивает значение счётчика на 1
func (c *Counter) Increment() {
	// atomic.AddInt64 выполняет операцию атомарно (т.е. неделимо):
	// ни одна другая горутина не сможет вмешаться посередине операции.
	atomic.AddInt64(&c.value, 1)
}

// Функция Value безопасно возвращает текущее значение счётчика
func (c *Counter) Value() int64 {
	// atomic.LoadInt64 гарантирует, что мы читаем значение без "гонок данных"
	return atomic.LoadInt64(&c.value)
}

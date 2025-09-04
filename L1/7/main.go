package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup

	// 10 горутин пишут в sync.Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			sm.Store(key, i) // безопасная запись
			val, _ := sm.Load(key)
			fmt.Println("Записано:", key, "=", val)
		}(i)
	}

	wg.Wait()
}

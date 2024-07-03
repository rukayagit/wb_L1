package main

import (
	"fmt"
	"sync"
)

// создаем структуру-счетчик с мьютексом для синхронизации
type Counter struct {
	mu    sync.Mutex // мьютекс для синхронизации доступа
	count int        // счетчик
}

// увеличиваем значение счетчика на 1
func (c *Counter) Increment() {
	c.mu.Lock()   // блокируем мьютекс перед изменением счетчика
	c.count++     // увеличиваем значение счетчика
	c.mu.Unlock() // разблокируем мьютекс после изменения счетчика
}

// возвращаем текущее значение счетчика
func (c *Counter) Value() int {
	c.mu.Lock()         // блокируем мьютекс перед получением значения счетчика
	defer c.mu.Unlock() // разблокируем мьютекс после получения значения счетчика
	return c.count
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// запускаем горутины, которые будут увеличивать значение счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // ожидаем завершения горутин
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}

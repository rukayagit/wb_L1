package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// создаем структуру-счетчик с атомарным целым числом
type Counter struct {
	count int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.count, 1) // увеличиваем счетчик на 1
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.count) // считываем текущее значение счетчика
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// запусскаем горутины, которые будут увеличивать счетчик
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // ждем завершения всех горутин
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}

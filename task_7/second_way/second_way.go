package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map       // для хранения данных
		wg sync.WaitGroup // группа ожидания для синхронизации горутин
		n  = 10           // количество горутин
	)

	// функция для записи в мапу
	write_to_map := func(key, value int) {
		defer wg.Done()
		m.Store(key, value)
	}

	// запуск горутин для конкурентной записи данных
	for i := 0; i < n; i++ {
		wg.Add(1)
		go write_to_map(i, i*i)
	}

	wg.Wait() // ждем завершения горутин

	// вывод данных
	m.Range(func(k, v interface{}) bool {
		fmt.Printf("Ключ: %v, значение: %v\n", k, v)
		return true
	})
}

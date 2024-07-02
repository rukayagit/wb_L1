package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu sync.Mutex          // mutex для синхронизации доступа к карте
		m  = make(map[int]int) //карта для хранения данных
		wg sync.WaitGroup      // группа ожидания для синхронизации горутин
		n  = 10                // количество горутин
	)

	// функция для записи данных в мапу
	write_to_map := func(key, value int) {
		defer wg.Done()
		mu.Lock()      // блокировка Mutex перед записью в мапу
		m[key] = value // запись в мапу
		mu.Unlock()    // разблокировка Mutex после записи в мапу
	}

	// запуск нескольких горутин для конкурентной записи данных в мапу
	for i := 0; i < n; i++ {
		wg.Add(1)
		go write_to_map(i, i*i)
	}

	wg.Wait() /// ожидание завершения всех горутин

	// вывод данных
	for k, v := range m {
		fmt.Printf("Ключ: %d, значение:  %d\n", k, v)
	}
}

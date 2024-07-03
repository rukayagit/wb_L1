package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu   sync.Mutex // мьютекс для синхронизации доступа к переменной stop
	stop bool       // флаг для остановки горутины
)

func worker() {
	for {
		mu.Lock() // блокируем мьютекс для доступа к переменной stop
		if stop { // проверяем флаг остановки
			mu.Unlock() // разблокируем мьютекс
			fmt.Println("Worker остановлен")
			return // выходим из функции, прекращаем выполнение горутины
		}
		mu.Unlock() // разблокируем мьютекс
		fmt.Println("Worker работает")
		time.Sleep(500 * time.Millisecond) // ждем полсекунды
	}
}

func main() {
	go worker() // запускаем worker в горутине

	time.Sleep(2 * time.Second) // ждем 2 секунды
	mu.Lock()                   // блокируем мьютекс для изменения флага stop
	stop = true                 // устанавливаем флаг остановки
	mu.Unlock()                 // разблокируем мьютекс
	time.Sleep(1 * time.Second) // ждем еще секунду, чтобы увидеть, что горутина остановилась
}

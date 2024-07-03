package main

import (
	"fmt"
	"time"
)

// горутина, которая будет останавливаться при получении сигнала
func worker(done chan bool) {
	for {
		select {
		case <-done: // при получения сигнала done, пишем, что:
			fmt.Println("Worker остановлен")
			return // выходим из функции и прекращаем выполнение горутины
		default: // если не получили сигнала, то пишем, что:
			fmt.Println("Worker работает")
			time.Sleep(500 * time.Millisecond) // ждем полсекунды
		}
	}
}

func main() {
	done := make(chan bool) // создаем канал для сигнала done
	go worker(done)         // запускаем worker в горутине

	time.Sleep(2 * time.Second) // ждем 2 секунды
	done <- true                // отправляем сигнал для остановки через канал
	time.Sleep(1 * time.Second) // ждем 1 секунду, чтобы увидеть, что горутина остановилась
}

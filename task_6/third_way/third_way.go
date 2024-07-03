package main

import (
	"fmt"
	"time"
)

func worker(stop <-chan time.Time) {
	for {
		select {
		case <-stop: // если таймер завершился, пишем, что:
			fmt.Println("Worker остановлен")
			return // выходим из функции и прекращаем выполнение горутины
		default: // если таймер еще не завершился, то пишем, что:
			fmt.Println("Worker работает")
			time.Sleep(500 * time.Millisecond) // ждем полсекунды
		}
	}
}

func main() {
	stop := time.After(2 * time.Second) // создадим таймер, который сработает через 2 секунды
	go worker(stop)                     // запускаем worker в горутине

	time.Sleep(3 * time.Second) // ждем 3 секунды чтобы увидеть, что горутина остановилась
}

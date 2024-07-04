package main

import (
	"fmt"
	"time"
)

func main() {
	// создаем канал для передачи данных:
	ch := make(chan int)

	// длительность выполнения программы в секундах:
	var N int
	fmt.Println("Введите длительность выполнения программы:")
	fmt.Scan(&N)

	// запуск горутины для отправки значений в канал:
	go func() {
		i := 0
		for {
			select {
			case ch <- i:
				i++
				time.Sleep(500 * time.Millisecond) // задержка для демонстрации
			}
		}
	}()

	// создание таймера завершения:
	timeout := time.After(time.Duration(N) * time.Second)

	// чтение значений из канала и завершение по таймеру:
	for {
		select {
		case val := <-ch:
			fmt.Println("Получено:", val)
		case <-timeout:
			fmt.Println("Время вышло")
			close(ch)
			return
		}
	}
}

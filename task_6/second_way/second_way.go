package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // если контекст отменен, пишем, что:
			fmt.Println("Worker остановлен")
			return // выходим из функции и прекращаем выполнение горутины
		default: // если контекст не отменен, то пишем, что:
			fmt.Println("Worker работает")
			time.Sleep(500 * time.Millisecond) // ждем полсекунды
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // создадим контекст с возможностью отмены
	go worker(ctx)                                          // запускаем worker в горутине

	time.Sleep(2 * time.Second) // ждем 2 секунды
	cancel()                    // отменяем контекст
	time.Sleep(1 * time.Second) // ждем 1 секунду, чтобы увидеть, что горутина остановилась
}

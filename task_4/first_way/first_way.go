package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// получаем данные из канала и выводим их в stdout
func worker(id int, data_chan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range data_chan {
		fmt.Printf("Воркер %d получил данные %d\n", id, data)
	}
}

func main() {
	// устанавливаем количество воркеров
	var num_workers int
	fmt.Println("Введите количество воркеров:")
	fmt.Scan(&num_workers)

	// создаем канал для данных
	data_chan := make(chan int)

	// используем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// запускаем воркеры
	for i := 0; i < num_workers; i++ {
		wg.Add(1)
		go worker(i, data_chan, &wg)
	}

	// канал для перехвата сигнала завершения (CTRL+C)
	sig_chan := make(chan os.Signal, 1)
	signal.Notify(sig_chan, syscall.SIGINT, syscall.SIGTERM)

	// создаем контекст с тайм-аутом для симуляции завершения программы через 10 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// горутина для записи данных в канал
	go func() {
		for {
			select {
			case <-sig_chan:
				// закрываем канал и завершаем работу при получении сигнала:
				close(data_chan)
				return
			case <-ctx.Done():
				// закрываем канал и завершаем работу при истечении времени:
				close(data_chan)
				return
			default:
				// записываем случайные данные в канал:
				data_chan <- rand.Intn(100)
				time.Sleep(1 * time.Second) // Пауза между записями
			}
		}
	}()

	// ожидаем завершения всех воркеров:
	wg.Wait()
	fmt.Println("Воркеры завершили работу")
}

package main

import (
	"fmt"
	"sync"
)

// функция, которая вычисляет квадрат числа и отправляет результат в канал
func worker(num int, result_chan chan int, wg *sync.WaitGroup) {
	defer wg.Done()          // уменьшаем счетчик WaitGroup после завершения функции
	result_chan <- num * num // отправляем квадрат числа в канал
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	result_chan := make(chan int, len(nums)) // создаем канал с буфером размером равным количеству чисел
	var wg sync.WaitGroup                    // создаем счетчик WaitGroup для ожидания завершения всех горутин

	// запускаем горутины для вычисления квадратов чисел
	for _, num := range nums {
		wg.Add(1)                        // увеличиваем счетчик WaitGroup
		go worker(num, result_chan, &wg) // запускаем горутину
	}

	wg.Wait()          // ожидаем завершения всех горутин
	close(result_chan) // закрываем канал

	// суммируем результаты из канала
	sum := 0
	for result := range result_chan {
		sum += result
	}
	fmt.Println("Сумма квадратов:", sum)
}

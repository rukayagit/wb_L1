package main

import (
	"fmt"
	"sync"
)

// send_numbers отправляет числа из массива в первый канал
func send_numbers(arr []int, ch1 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик WaitGroup по завершении функции
	for _, num := range arr {
		ch1 <- num // отправляем число в канал ch1
	}
	close(ch1) // закрываем канал после отправки всех чисел
}

// multiply_by_two читает числа из первого канала, умножает их на 2 и отправляет во второй канал
func multiply_by_two(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик WaitGroup по завершении функции
	for num := range ch1 {
		ch2 <- num * 2 // умножаем число на 2 и отправляем в канал ch2
	}
	close(ch2) // закрываем канал после обработки всех чисел
}

// print_numbers читает числа из второго канала и выводит их на стандартный вывод
func print_numbers(ch2 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик WaitGroup по завершении функции
	for num := range ch2 {
		fmt.Println(num) // выводим число на стандартный вывод
	}
}

func main() {
	// массив чисел для обработки
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// создаем два канала
	ch1 := make(chan int)
	ch2 := make(chan int)

	// создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(3) // устанавливаем счетчик WaitGroup на количество горутин

	// запускаем горутины для обработки данных в конвейере
	go send_numbers(arr, ch1, &wg)    // запускаем отправку чисел в первый канал
	go multiply_by_two(ch1, ch2, &wg) // запускаем умножение чисел на 2 и отправку во второй канал
	go print_numbers(ch2, &wg)        // читаем и выводим числа из второго канала

	// ожидаем завершения всех горутин
	wg.Wait()
}

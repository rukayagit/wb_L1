package main

import (
	"fmt"
)

// send_numbers отправляет числа из массива в первый канал
func send_numbers(arr []int, ch1 chan<- int) {
	for _, num := range arr {
		ch1 <- num //отправляем число в канал
	}
	close(ch1) //закрываем канал после отправки всех чисел
}

/*
multiply_numbers читает числа из первого канала,
умножает их на 2 и отправляет во второй канал:
*/
func multiply_numbers(ch1 <-chan int, ch2 chan<- int) {
	for num := range ch1 {
		ch2 <- num * 2 //умножаем число на 2 и отправляем во второй канал
	}
	close(ch2) //закрываем канал после отправки всех чисел
}

// print_numbers читает числа из второго канала и выводит их в stdout
func print_numbers(ch2 <-chan int) {
	for num := range ch2 {
		fmt.Println(num) //выводим число в stdout
	}
}

func main() {
	// инициализируем массив чисел
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// создаем два канала
	ch1 := make(chan int)
	ch2 := make(chan int)

	// запускаем горутины для отправки, умножения и печати чисел
	go send_numbers(arr, ch1)
	go multiply_numbers(ch1, ch2)
	go print_numbers(ch2)

	// ждем закрытия второго канала
	<-ch2

	// запускаем горутины для обработки данных в конвейере:
	go send_numbers(arr, ch1)
	go multiply_numbers(ch1, ch2)
	print_numbers(ch2)
}

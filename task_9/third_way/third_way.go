package main

import (
	"fmt"
)

func main() {
	// массив чисел для обработки
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// создаем два канала и дополнительные каналы для синхронизации
	ch1 := make(chan int)
	ch2 := make(chan int)
	done1 := make(chan struct{})
	done2 := make(chan struct{})

	// запускаем горутину для отправки чисел в первый канал
	go func(arr []int, ch1 chan<- int, done chan<- struct{}) {
		for _, num := range arr {
			ch1 <- num // отправляем число в канал ch1
		}
		close(ch1)         // закрываем канал после отправки всех чисел
		done <- struct{}{} // сообщаем, что горутина завершена
	}(arr, ch1, done1)

	// запускаем горутину для умножения чисел на 2 и отправки во второй канал
	go func(ch1 <-chan int, ch2 chan<- int, done chan<- struct{}) {
		for num := range ch1 {
			ch2 <- num * 2 // умножаем число на 2 и отправляем в канал ch2
		}
		close(ch2)         // закрываем канал после обработки всех чисел
		done <- struct{}{} // сообщаем, что горутина завершена
	}(ch1, ch2, done2)

	// запускаем горутину для чтения и вывода чисел из второго канала
	go func(ch2 <-chan int) {
		for num := range ch2 {
			fmt.Println(num) // выводим число на стандартный вывод
		}
	}(ch2)

	// ожидаем завершения всех горутин
	<-done1
	<-done2
}

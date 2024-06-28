package main

import "fmt"

func main() {
	//массив с числами:
	numbers := []int{2, 4, 6, 8, 10}

	//создаем буферизированный канал, через который будут передаваться результаты:
	results := make(chan int, len(numbers))

	//горутина для вычисления квадратов:
	for _, num := range numbers {
		go func(x int) {
			results <- x * x //вычисляем квадрат числа и отправляем его в канал
		}(num)
	}

	//выводим результаты:
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-results)
	}
}

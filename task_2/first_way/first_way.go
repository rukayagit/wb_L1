package main

import (
	"fmt"
	"sync"
)

func main() {
	//массив с числами:
	numbers := []int{2, 4, 6, 8, 10}

	//создаем канал, через который будут передаваться результаты:
	results := make(chan int)

	//создаем WaitGroup, который будет ждать завершения всех горутин:
	var wg sync.WaitGroup

	//горутина для вычисления квадратов:
	for _, num := range numbers {
		wg.Add(1) //увеличиваем счетчик горутин
		go func(x int) {
			defer wg.Done()  //уменьшаем счетчик горутин после завершения функции
			results <- x * x //вычисляем квадрат числа и отправляем его в канал
		}(num)
	}

	//горутина для закрытия канала после завершения всех горутин:
	go func() {
		wg.Wait()      //ждем окончания горутин
		close(results) //закрываем канал
	}()

	//выводим результаты:
	for result := range results {
		fmt.Println(result)
	}
}

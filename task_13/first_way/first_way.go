package main

import "fmt"

func main() {
	var a, b int //создаем переменные a и b
	fmt.Println("Введите числа a и b:")
	fmt.Scan(&a, &b) //вводим значения a и b
	fmt.Printf("До перестановки: a = %d, b =  %d\n", a, b)

	a = a + b //теперь а содержит сумму a и b
	b = a - b //b содержит исходное значение a
	a = a - b //a содержит исходное значение b

	fmt.Printf("После перестановки: a = %d, b = %d\n", a, b)
}

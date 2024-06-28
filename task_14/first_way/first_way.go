package main

import "fmt"

func determine_type(v interface{}) {
	//используем type switch для определения типа переменной
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	//введем значения переменных различных типов
	var a int = 10
	var b string = "Wildberries"
	var c bool = true
	var d chan int = make(chan int)

	//определим тип каждой переменной
	determine_type(a)
	determine_type(b)
	determine_type(c)
	determine_type(d)
}

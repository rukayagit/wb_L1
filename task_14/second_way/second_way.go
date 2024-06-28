package main

import (
	"fmt"
	"reflect"
)

func determine_type(v interface{}) {
	//используем reflect.TypeOf для получения типа переменной:
	t := reflect.TypeOf(v)
	//определяем тип переменной:
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		//если тип переменноц канал, проверяем тип элементов канала:
		if t.Elem().Kind() == reflect.Int {
			fmt.Println("Канал элементов типа int")
		} else {
			fmt.Println("Канал элементов неизвестного типа")
		}
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

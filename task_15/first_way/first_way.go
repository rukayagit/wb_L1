package main

import (
	"fmt"
)

// глобальная переменная для хранения подстроки
var justString string

// функция для создания строки заданного размера, заполненной символами 'W'
func createHugeString(size int) string {
	hugeString := make([]rune, size)
	for i := range hugeString {
		hugeString[i] = 'W'
	}
	return string(hugeString)
}

// функция для создания подстроки из большой строки:
func someFunc() {
	v := createHugeString(1 << 10) // создаём большую строку размером 1024 символа

	// проверка, чтобы избежать выхода за границы строки
	if len(v) < 100 {
		justString = v
	} else {
		// создание новой строки из первых 100 символов
		justString = string([]rune(v[:100]))
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}

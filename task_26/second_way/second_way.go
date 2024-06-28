package main

import (
	"fmt"     //импортируем пакет для форматирования и вывода данных
	"strings" //импортируем пакет для работы с строками
)

func is_unique_string(s string) bool {
	var ascii_set [128]bool

	for _, char := range s { //перебираем все символы в строке
		index := strings.ToLower(string(char))[0] //преобразуем символ в нижний регистр

		if ascii_set[index] {
			return false //если символ уже встречался, возвращаем false
		}
		ascii_set[index] = true //если символ не встречался, добавляем его в ascii_set
	}
	return true
}

func main() {
	var s string
	fmt.Println("Введите строку:")
	fmt.Scan(&s)
	fmt.Println("Введенная строка:", s)

	answer := is_unique_string(s)
	if answer {
		fmt.Println("Все символы в этой строке уникальны")
	} else {
		fmt.Println("В этой строке есть повторяющиеся символы")
	}
}

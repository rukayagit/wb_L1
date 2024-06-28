package main

import (
	"fmt" //импортируем пакет для форматирования и вывода данных
	"unicode"
)

func is_unique_string(s string) bool {
	seen := make(map[rune]bool) //создаем мапу для отслеживания символов в строке

	for _, char := range s { //перебираем все символы в строке
		char := unicode.ToLower(char) //приводим все символы к нижнему регистру

		if seen[char] {
			return false //если символ уже есть в мапе, возвращаем false
		}
		seen[char] = true //если символа нет в мапе, добавляем его
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

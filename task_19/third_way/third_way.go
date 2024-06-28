package main

import (
	"bytes" //импортируем пакет для работы с буферами
	"fmt"   //импортируем пакет для форматирования и вывода данных
)

func reverse_string(s string) string {
	var buffer bytes.Buffer //создаем буфер для построения перевернутой строки

	runes := []rune(s) //преобразуем строку в символы(срез рун)

	for i := len(runes) - 1; i >= 0; i-- {
		buffer.WriteRune(runes[i]) //записываем каждый символ в буфер в обратном порядку
	}

	return buffer.String() //преобразуем буфер в строку и возвращаем ее
}

func main() {
	var s string
	fmt.Println("Введите строку:")
	fmt.Scan(&s)
	fmt.Println("Введенная строка:", s)

	reversed := reverse_string(s)
	fmt.Println("Перевернутая строка:", reversed)
}

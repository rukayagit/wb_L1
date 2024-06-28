package main

import "fmt" //импортируем пакет для форматирования и вывода данных

func reverse_string(s string) string {
	runes := []rune(s) //преобразуем строку в символы(срез рун)
	n := len(runes)    //найдем длину строки

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i] //перевернем символы
	}
	return string(runes) //преобразуем символы в строку и возвращаем ее
}

func main() {
	var s string
	fmt.Println("Введите строку:")
	fmt.Scan(&s)
	fmt.Println("Введенная строка:", s)

	reversed := reverse_string(s)
	fmt.Println("Перевернутая строка:", reversed)
}

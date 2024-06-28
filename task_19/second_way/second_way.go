package main

import "fmt" //импортируем пакет для форматированного ввода и вывода данных

func reverse_string(s string) string {
	if len(s) == 0 {
		/* базовый случай, рекурсия будет продолжаться
		до тех пор, пока это условие не выполнится */
		return s
	}
	runes := []rune(s)                                          //преобразуем строку в символы
	return reverse_string(string(runes[1:])) + string(runes[0]) //переворачиваем строку с помощью рекурсии
}

func main() {
	var s string
	fmt.Println("Введите строку:")
	fmt.Scan(&s)
	fmt.Println("Введенная строка:", s)

	reversed := reverse_string(s)
	fmt.Println("Перевернутая строка:", reversed)
}

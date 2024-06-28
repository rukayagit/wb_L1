package main

import (
	"bufio"   //импортируем пакет для ввода данных с клавиатуры
	"fmt"     //импортируем пакет для форматирования и вывода данных
	"os"      //импортируем пакет для доступа к фунциональности операционной системы
	"strings" //импортируем пакет для работы с строками
)

func reverse_words_in_string(s string) string {
	words := strings.Fields(s) //делим строку на слова

	if len(words) <= 1 {
		return s //базовый случай. рекурсия будет продолжаться, пока это условие не выполнится
	}
	//переворачиваем оставшуюся часть строки и соединяем ее с первым словом:
	return reverse_words_in_string(strings.Join(words[1:], " ")) + " " + words[0]
}

func main() {
	fmt.Println("Введите строку:")
	scanner := bufio.NewScanner(os.Stdin) //инициализируем буфер для ввода
	for scanner.Scan() {                  //пока не конец строки
		s := scanner.Text() //получаем строку
		fmt.Println("Введенная строка:", s)

		reversed := reverse_words_in_string(s)
		fmt.Println("Перевернутая строка:", reversed)
	}

}

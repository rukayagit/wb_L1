package main

import (
	"bufio"   //импортируем пакет для ввода данных с клавиатуры
	"fmt"     //импортируем пакет для форматирования и вывода данных
	"os"      //импортируем пакет для доступа к фунциональности операционной системы
	"strings" //импортируем пакет для работы с строками
)

func reverse_words_in_string(s string) string {
	words := strings.Fields(s) //делим строку на слова

	for i, j := 0, (len(words) - 1); i < j; i, j = (i + 1), (j - 1) {
		words[i], words[j] = words[j], words[i] //переворачиваем порядок слов
	}

	return strings.Join(words, " ") //соединяем слова обратно в строку и возвращаем ее
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

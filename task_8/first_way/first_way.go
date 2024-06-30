package main

import "fmt"

// функция для установки i-го бита в значении n в 1 или 0
func set_bit(n int64, i uint, bit_value uint) int64 {
	mask := int64(1 << i) // создаем маску, сдвигая число 1 влево на i позиций. Это создаёт число, у которого только i-й бит равен 1.
	if bit_value == 1 {
		n = n | mask // устанавливаем i-й бит в 1
	} else {
		n = n &^ mask // устанавливаем i-й бит в 0
	}
	return n // возвращаем обновленное n
}

func main() {
	var number int64 = 97 // в двоичной системе: 1100001

	number = set_bit(number, 6, 0)
	fmt.Printf("Число после установки 7-го бита в 0: %d\n", number)

	number = set_bit(number, 3, 1)
	fmt.Printf("Число после установки 4-го бита в 1: %d\n", number)
}

package main

import "fmt"

func quicksort(arr []int) {
	if len(arr) < 2 {
		return //базовый случай: массив с 0 или 1 элементом уже отсортирован
	}

	//выбираем опорный элемент
	oe := arr[len(arr)/2]

	//делим массив на две части: меньше и больше опорного элемента
	left, right := partition(arr, oe)

	//с помощью рекурсии сортируем левую и правую части
	quicksort(left)
	quicksort(right)

	//объединяем левую и правую часть
	copy(arr, append(append(left, oe), right...))
}

// функция для разделения массива на части
func partition(arr []int, oe int) (left, right []int) {
	for _, v := range arr {
		if v < oe {
			left = append(left, v) //элементы меньше опорного переносятся в левую часть
		} else if v > oe {
			right = append(right, v) //соответственно, элементы больше опорного переносятся в правую часть
		}
	}
	return left, right
}

func main() {
	//создаем массив для сортировки:
	nums := []int{10, 7, 15, 6, 31, 98, 3, 24, 69, 43}
	fmt.Println("Исходный массив:", nums)

	quicksort(nums)

	fmt.Println("Отсортированный массив:", nums)
}

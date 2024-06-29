package main

import "fmt"

func bynary_search(num []int, target int) int {
	left, right := 0, len(num)-1 // устанавливаем границы для поиска

	for left <= right {
		mid := left + (right-left)/2 //вычисляем индекс среднего элемента
		if num[mid] == target {      //если средний элемент равен тому, котороый мы ищем,возвращаем его индекс
			return mid
		} else if num[mid] < target { //если средний элемент меньше целевого, сдвигаем левую границу
			left = mid + 1
		} else { //если средний элемент больше целевого, сдвигаем правую границу
			right = mid - 1
		}
	}
	return -1 //если элемент не найден, возвращаем -1
}

func main() {
	//создаем отсортированный срез чисел:
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var target int //число, которое мы хотим найти
	fmt.Println("Введите число, которое хотите найти:")
	fmt.Scan(&target)

	index := bynary_search(nums, target)

	if index == -1 {
		fmt.Printf("Число %d не найдено", target)
	} else {
		fmt.Printf("Число %d найдено по индексу %d", target, index)
	}
}

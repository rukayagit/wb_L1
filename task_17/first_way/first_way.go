package main

import (
	"fmt"
	"sort"
)

func main() {
	//создаем отсортированный срез чисел:
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var target int //число, которое мы хотим найти
	fmt.Println("Введите число, которое хотите найти:")
	fmt.Scan(&target)

	//используем цикл, который будет искать в массиве target с помощью sort.Search
	index := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })

	//проверяем, есть ли элемент в массиве
	if index < len(nums) && nums[index] == target {
		fmt.Printf("Число %d найдено с индексом %d\n", target, index)
	} else {
		fmt.Printf("Число %d не найдено в массиве\n", target)
	}
}

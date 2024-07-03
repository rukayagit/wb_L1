package main

import "fmt"

// функция, которая возвращает пересечение двух множеств
func intersect(set1, set2 []int) []int {
	// создаем карту для хранения элементов первого множества
	map1 := make(map[int]struct{})
	for _, val := range set1 {
		map1[val] = struct{}{}
	}

	// создаем срез для хранения пересечений
	intersection := []int{}

	// создаем карту для хранения элементов второго множества
	for _, val := range set2 {
		// если элемент есть в первом множестве, то добавляем его в срез
		if _, found := map1[val]; found {
			intersection = append(intersection, val)
			// удаляем этот элемент из среза
			delete(map1, val)
		}
	}
	return intersection // возвращаем срез
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	result := intersect(set1, set2)
	fmt.Printf("Пересечение множеств: %v", result)
}

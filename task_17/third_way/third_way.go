package main

import (
	"fmt"
	"sort"
)

// определяем структуру пончиков
type Donuts struct {
	Name string
}

func main() {
	//создаем отсортированный срез структур пончиков для поиска
	donuts := []Donuts{
		{"blueberry"},
		{"chocolate"},
		{"raspberry"},
		{"strawberry"},
		{"vanilla"},
	}

	//сортируем срез по алфавиту
	sort.Slice(donuts, func(i, j int) bool {
		return donuts[i].Name < donuts[j].Name
	})

	var target string //вкус пончика, который мы хотим найти
	fmt.Println("Введите название вкуса, который хотите найти:")
	fmt.Scan(&target)

	index := sort.Search(len(donuts), func(i int) bool { return donuts[i].Name >= target })

	if index < len(donuts) && donuts[index].Name == target {
		fmt.Printf("Вкус %s найден с индексом: %d\n", target, index)
	} else {
		fmt.Printf("Вкус %s не найден\n", target)
	}
}

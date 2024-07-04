package main

import (
	"fmt"
	"math"
)

// группировка температур с шагом в 10 градусов
func group_temperatures(temperatures []float64) map[int][]float64 {
	// создаем карту для хранения групп температур
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		// вычисляем ключ для текущей температуры
		key := int(math.Floor(temp/10.0) * 10)
		if temp < 0 {
			key = int(math.Ceil(temp/10.0) * 10)
		}
		// добавляем температуру в соответствующую группу
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	// исходная последовательность температурных колебаний:
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// группируем температуры:
	groups := group_temperatures(temperatures)

	// выводим результат:
	for key, temps := range groups {
		fmt.Printf("Группа %d: %v\n", key, temps)
	}
}

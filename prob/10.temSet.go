/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

// реализован генератор случайных чисел
func main() {
	var temperatures []float64

	for i := 0; i < 100; i++ {
		nums := rand.Float64()*80 - 40
		temperatures = append(temperatures, nums)
	}
	fmt.Printf("Сгенерированный слайс случайных чисел: %.2f", temperatures)

	// Группировка значений
	grouped := groupTemperatures(temperatures)

	// Печать результатов
	for key, values := range grouped {
		fmt.Printf("%d: ", key)
		for _, value := range values {
			fmt.Printf("%.2f ", value)
		}
		fmt.Println()
	}
}

// groupTemperatures группирует температуры с шагом в 10 градусов
func groupTemperatures(temps []float64) map[int][]float64 {
	grouped := make(map[int][]float64)
	for _, temp := range temps {
		// Округление до ближайшего нижнего десятка
		key := int(math.Floor(temp/10)) * 10
		grouped[key] = append(grouped[key], temp)
	}
	return grouped
}

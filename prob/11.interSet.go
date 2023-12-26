/*
Реализовать пересечение двух неупорядоченных множеств.
*/

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}
	intersection := make([]int, 0)
	for _, num := range slice1 { // Перебор каждого числа из первого слайса
		for _, num2 := range slice2 { // Перебор каждого числа из второго слайса
			if num == num2 { // Проверка на совпадение чисел
				intersection = append(intersection, num) // Добавление числа в пересечение
				break                                    // Выход из внутреннего цикла, так как совпадение найдено
			}
		}
	}
	fmt.Println(intersection)

	// Создание карты для хранения элементов первого слайса
	slice1Map := make(map[int]bool)
	for _, num := range slice1 {
		slice1Map[num] = true
	}

	// Создание карты для хранения элементов второго слайса
	slice2Map := make(map[int]bool)
	for _, num := range slice2 {
		slice2Map[num] = true
	}

	// Пересечение слайсов с использованием карт
	intersection = make([]int, 0)
	for num := range slice1Map {
		if slice2Map[num] {
			intersection = append(intersection, num)
		}
	}
	fmt.Println(intersection)
}

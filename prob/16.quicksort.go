/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

package main

import (
	"fmt"
	"math/rand"
)

// randNumGenerator генерирует слайс случайных целых чисел заданной длины.
// Принимает целое число n и возвращает слайс из n случайных целых чисел.
// Числа генерируются функцией rand.Intn, и добавляются в слайс nums.
func randNumGenerator(n int) []int {
	var nums []int
	for i := 0; i < n; i++ {
		num := rand.Intn(n)      // Генерация случайного числа от 0 до n-1
		nums = append(nums, num) // Добавление числа в слайс
	}
	return nums
}

// quicksort реализует алгоритм быстрой сортировки.
// Если слайс содержит 0 или 1 элемент, он уже отсортирован.
// Выбирается опорный элемент (здесь первый элемент слайса).
// Элементы меньше опорного помещаются в слайс left, больше — в right.
// Рекурсивно сортируются слайсы left и right, и они объединяются с опорным элементом.
func quicksort(num []int) []int {
	if len(num) <= 1 {
		return num // Слайс из 0 или 1 элемента уже отсортирован
	}
	pivot := num[0] // Опорный элемент
	var left []int  // Элементы меньше опорного
	var right []int // Элементы больше опорного

	// Разделение элементов на подмножества left и right
	for _, v := range num[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	// Рекурсивная сортировка подмножеств
	left = quicksort(left)
	right = quicksort(right)

	// Объединение отсортированных подмножеств с опорным элементом
	return append(append(left, pivot), right...)
}
func main() {
	num := randNumGenerator(10) // Генерация случайного набора чисел
	fmt.Println(num)
	sortedNum := quicksort(num) // Сортировка чисел
	fmt.Println(sortedNum)
}

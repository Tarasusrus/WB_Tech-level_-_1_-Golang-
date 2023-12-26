/*
Реализовать бинарный поиск встроенными методами языка.
Бинарный — двоичный, представленный элементами всего двух видов.
Главное правило бинарного поиска — он работает только с отсортированными данными.


*/

package main

import (
	"fmt"
	"math/rand"
)

func randNumGenerator1(n int) []int {
	var nums []int
	for i := 0; i < n; i++ {
		num := rand.Intn(n)      // Генерация случайного числа от 0 до n-1
		nums = append(nums, num) // Добавление числа в слайс
	}
	return nums
}

func quicksort1(num []int) []int {
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
	left = quicksort1(left)
	right = quicksort1(right)

	// Объединение отсортированных подмножеств с опорным элементом
	return append(append(left, pivot), right...)
}

func binarySerch(nums []int, lookingNum int) {
	low := 0              // Начальный индекс массива
	high := len(nums) - 1 // Конечный индекс массива

	// Цикл выполняется, пока диапазон поиска не сузится до нуля
	for low <= high {
		mid := (low + high) / 2 // Вычисление середины диапазона

		// Проверка, равно ли среднее значение искомому числу
		if nums[mid] == lookingNum {
			fmt.Println("Число найдено в индексе", mid) // Вывод индекса найденного числа
			return                                      // Завершение функции, так как число найдено
		} else if nums[mid] < lookingNum {
			fmt.Println("Увеличиваю индекс") // Информирование о ходе поиска
			low = mid + 1                    // Сужение диапазона поиска к более высоким индексам
		} else {
			fmt.Println("Уменьшаю индекс") // Информирование о ходе поиска
			high = mid - 1                 // Сужение диапазона поиска к более низким индексам
		}
	}

	// Если выполнение дошло до этого момента, число в массиве не найдено
	fmt.Println("Число не найдено")
}

func main() {
	nums := randNumGenerator1(100)
	sortedNums := quicksort1(nums) // Сортировка массива
	fmt.Println(sortedNums)
	binarySerch(sortedNums, 57) // Использование отсортированного массива для поиска

}

/*Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и
выведет их квадраты в stdout.
1. Создать массив (2,4,6,8,10)
2. Создать функцию для перебора цикла
3. создать функцию для возведения в квадрат n * n
4. Вывести n * n в stdout
*/

package main

import "fmt"

var nums = []int{2, 4, 6, 8, 10} // Инициализация массива чисел.

// Функция generation создает канал, который отправляет числа из массива.
func generation(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // Закрытие канала при завершении
		for _, n := range nums {
			out <- n // Отправка чисел в канал.
		}
	}()
	return out // Возврат канала для чтения.
}

// Функция square принимает канал чисел, возводит их в квадрат и отправляет в новый канал.
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // Закрытие канала при завершении
		for n := range in {
			out <- n * n // Возведение в квадрат и отправка.
		}
	}()
	return out // Возврат канала с квадратами чисел.
}

// Функция sumOfSquare суммирует квадраты чисел из канала и отправляет сумму в новый канал.
func sumOfSquare(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // Закрытие канала при завершении
		sum := 0
		for n := range in {
			sum += n // Суммирование чисел.
		}
		out <- sum // Отправка результата.
	}()
	return out // Возврат канала с суммой квадратов.
}

func main() {
	// Вывод квадратов чисел.
	for n := range square(generation(nums...)) {
		fmt.Println(n)
	}
	// Вывод суммы квадратов чисел.
	for n := range sumOfSquare(square(generation(nums...))) {
		fmt.Println(n)
	}
}

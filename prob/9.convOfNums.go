/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// maker — функция, которая генерирует случайные числа и отправляет их в указанный канал.
// Она работает бесконечно, засыпая на одну секунду между каждой операцией отправки.
// Генерируемые числа находятся в диапазоне от 0 до 49 (включительно).
// Функция принимает один аргумент, numbers, который является каналом типа int.
// Это однонаправленный канал, который позволяет только отправлять значения в канал.
func maker(numbers chan<- int) {
	for {
		num := rand.Intn(50)
		numbers <- num
		time.Sleep(time.Second)
	}
}

// square1 функция, которая читает целые числа из канала "numbers" и удваивает их.
// Полученные числа отправляются в канал "squares".
// Эта функция работает в цикле, непрерывно читая целые числа из канала "numbers"
// до его закрытия. Затем она вычисляет квадрат каждого целого числа и отправляет
// результат в канал "squares".
func square1(numbers <-chan int, squares chan<- int) {
	for num := range numbers {
		squares <- num * 2
	}
}

// collector — функция, которая получает квадраты чисел из указанного канала и выводит их в консоль.
// Она работает до тех пор, пока не закроется канал squares.
// При получении каждого квадрата числа, функция выводит сообщение в консоль.
// Функция принимает один аргумент, squares, который является каналом типа int.
// Это однонаправленный канал, который позволяет только получать значения из канала.
func collector(squares <-chan int) {
	for square := range squares {
		fmt.Println("сообщение: ", square)
	}
}

// main - функция, координирующая работу программы.
// Она создает каналы для передачи чисел и результата,
// запускает в фоне 3 горутины: maker, square1 и collector,
// и ожидает выполнения программы на протяжении 10 секунд.
// Все числа, сгенерированные функцией maker,
// будут отправлены в канал numbers,
// удвоены функцией square1 и отправлены в канал squares,
// а затем выведены на экран функцией collector.
//
// Функция main не имеет аргументов и не возвращает значений.
// Ее единственная цель - запустить и координировать
// работу других функций в программе и обеспечить ее завершение.
// Чтобы программа не завершилась сразу, добавлено ожидание
// выполнения программы в виде паузы на 10 секунд.
func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go maker(numbers)
	go square1(numbers, squares)
	go collector(squares)

	// Чтобы программа не завершилась сразу, добавим ожидание
	time.Sleep(10 * time.Second)
}
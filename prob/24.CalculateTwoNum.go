//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.

package main

import "fmt"

// это битовый сдвиг влево == 2**20, так же можно использовать math.pow
// var a float64 = 1 << 20
// var b float64 = 1 << 20

// возведение в степень через цикл
func pow(power, num int) int {
	result := num
	for i := 1; i < power; i++ {
		result *= num
	}
	return result
}
func main() {
	var a, b int64 = 1 << 21, 1 << 22 // значения > 2^20
	fmt.Println(pow(20, 2))
	// Арифметические операции
	quot := b / a // Деление
	prod := a * b // Умножение
	sum := a + b  // Сложение
	diff := a - b // Вычитание

	fmt.Println("Сложение:", sum)
	fmt.Println("Вычитание:", diff)
	fmt.Println("Умножение:", prod)
	fmt.Println("Деление:", quot)
}

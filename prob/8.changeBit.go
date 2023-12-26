/*
Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.

В Go битовые операции включают & (И),
| (ИЛИ),
^ (исключающее ИЛИ),
и &^ (И-НЕ),
а также битовые сдвиги << и >>
*/

package main

import "fmt"

func setBit(n int64, pos uint, val int) int64 {
	if val == 1 {
		// Устанавливаем бит в 1
		n |= 1 << pos
	} else {
		// Устанавливаем бит в 0
		n &^= 1 << pos
	}
	return n
}

func main() {
	var n int64 = 55
	var pos uint = 5 // Позиция бита, который мы хотим изменить
	var val int = 0
	fmt.Printf("Исходное число: %d (%b)\n", n, n)
	n = setBit(n, pos, val)
	fmt.Printf("Измененное число: %d (%b)\n", n, n)
}

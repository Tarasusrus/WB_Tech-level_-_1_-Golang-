//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 1
	b := 2
	a, b = b, a // множественное присваивание
	fmt.Println(a, b)
}

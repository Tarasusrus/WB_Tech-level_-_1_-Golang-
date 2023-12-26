/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

package main

import "fmt"

// делаем функцию для определения типа через переменную итерфейс
func indentifyType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Тип int:", v)
	case string:
		fmt.Println("Тип string:", v)
	case bool:
		fmt.Println("Тип bool:", v)
	case chan int:
		fmt.Println("Тип channel (chan int):", v)
	default:
		fmt.Println("Неизвестный тип")
	}

}

func main() {
	indentifyType(42)
	indentifyType("hello")
	indentifyType(true)
	indentifyType(make(chan int))
	indentifyType(3.14)

}

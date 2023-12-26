/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package main

import "fmt"

func main() {
	// Исходная последовательность строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание множества с помощью карты
	set := make(map[string]bool)

	// Добавление элементов в множество
	for _, str := range arr {
		set[str] = true
	}

	// Вывод множества
	for str := range set {
		fmt.Println(str)
	}

}

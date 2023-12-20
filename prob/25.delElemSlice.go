package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2 // Удаляем элемент на индексе 2 (число 3)

	slice = append(slice[:index], slice[index+1:]...)

	fmt.Printf("%#v", slice)
}

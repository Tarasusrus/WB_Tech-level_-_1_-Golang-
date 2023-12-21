// Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s)               // в переменной words хранится слайс со строкой разботой по пробелу
	for i := len(words)/2 - 1; i >= 0; i-- { // чтобы не выходить за границы массива двигаемся от середины к началу
		opp := len(words) - 1 - i                   // вычисляется как индекс противоположного элемента, который нужно поменять местами.
		words[i], words[opp] = words[opp], words[i] // множественное присваивание
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ReverseWords("snow dog sun")) //
}

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».

Разделить строку на слова
перевернуть слово
склеить строку
*/

package main

import (
	"fmt"
	"strings"
)

func reversWord(s string) string {
	//Разделить строку на слова
	// Функция strings.Fields автоматически разделяет по пробелам и игнорирует лишние пробелы.
	words := strings.Fields(s)

	// Переворачиваем слова.
	// Идем по слайсу слов до середины, меняя местами элементы с начала и конца.
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	// Объединяем перевернутые слова в одну строку, разделяя их пробелами.
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reversWord(input)
	fmt.Println("Оригинальная строка:", input)
	fmt.Println("Перевернутая строка:", reversed)
}

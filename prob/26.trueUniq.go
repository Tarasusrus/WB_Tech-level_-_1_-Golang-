/* Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

// функция применяет хеш-таблицу для отслеживанияу никальных значений
func areAllCharactersUniq(input string) bool {
	input = strings.ToLower(input) // Приведение строки к нижнему регистру
	seen := make(map[rune]bool)    // Хеш-таблица для отслеживания уникальных символов
	for _, char := range input {   // Итерация по каждому символу в строке
		if _, ok := seen[char]; ok { // Проверка, если символ уже в хеш-таблице
			return false // Найден повторяющийся символ
		}
		seen[char] = true // Добавление символа в хеш-таблицу
	}
	return true // все символы уникальны
}

func main() {
	fmt.Println(areAllCharactersUniq("abcd"))
	fmt.Println(areAllCharactersUniq("abCdefAaf"))
	fmt.Println(areAllCharactersUniq("aabcd"))

}

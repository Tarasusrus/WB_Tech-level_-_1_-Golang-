/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.


var justString string - глобальная переменная
func someFunc() {
	v := createHugeString(1 << 10) - взяли огромную строку и не использовали всю
	justString = v[:100] - ссылка на строку не даст ГК почистить память
}
func main() {
	someFunc()
}


*/

package main

func createHugeString1(size int) string {
	return string(make([]byte, size))
}

func someFunc1() string {
	v := createHugeString1(1 << 10) // Создание большой строки
	justString := string(v[:100])   // Создание новой строки с нужными данными
	return justString               // Возврат строки, теперь память может быть освобождена
}

func main() {
	result := someFunc1() // Работаем с результатом функции
	_ = result            // Для примера, используем переменную, чтобы избежать ошибки
}

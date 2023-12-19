/*

Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).

*/

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() string {
	return "Привет эта функция определена в родительском классе и возвращает имя " + h.Name
}

/*
Теперь, если мы хотим создать структуру Action,
которая встраивает Human и добавляет свои уникальные поля или методы:
*/

type Action struct {
	// Чтобы методы родительского класса были доступнты мы наследуем встраивая в структуру это реализация одного из приципов ООП
	Human
	Action string
}

func (a Action) PerformAction() string {
	return " я вызываю метод унаследованную от класса Human" + a.Name + " выполняется в " + a.Action
}

func main() {
	a := Action{
		Human{Name: "Taras", Age: 33},
		"Метод определенный в структуре Action",
	}
	fmt.Println(a.Speak())         // Вызов метода Speak от Human
	fmt.Println(a.PerformAction()) // Вызов метода PerformAction от Action
}

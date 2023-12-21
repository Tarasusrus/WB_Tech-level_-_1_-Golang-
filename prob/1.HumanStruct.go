package main

import "fmt"

const WelcomeString = "Привет, эта функция определена в родительском классе и возвращает данные: "

type Human struct {
	Name string
	Age  int
	Sex  string
}

func DefaultHuman(name string) Human {
	h := Human{Name: name, Age: 45, Sex: "Мужской"} // использование одиночной присваивающей операции
	return h
}

func (h *Human) Speak() string {
	return WelcomeString + h.Name
}

type Action struct {
	Human             // Реализация наследования путем встраивания в структуру
	ActionDescription string
}

func (a Action) PerformAction() string {
	return "Я вызываю метод, унаследованный от класса Human. " + a.Name + " выполняет действие: " + a.ActionDescription
}

func main() {
	person := DefaultHuman("Иван Иванов")
	action := Action{
		Human{Name: "Тарас", Age: 33, Sex: "Мужской"},
		"Метод, определенный в структуре Action",
	}
	fmt.Println(action.Speak())         // вызываем метод Speak() от Human
	fmt.Println(action.PerformAction()) // вызываем метод PerformAction() от Action
	fmt.Println(person.Speak())         // вызываем метод Speak() от Human
}

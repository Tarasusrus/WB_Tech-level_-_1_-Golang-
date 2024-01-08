package main

import (
	"fmt"
	"strings"
)

/*
Паттерн Facade относится к структурным паттернам уровня объекта.

Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде
набора имен методов к набору взаимосвязанных классов или объектов некоторой
подсистемы, что облегчает ее использование.

Разбиение сложной системы на подсистемы позволяет упростить процесс разработки, а
также помогает максимально снизить зависимости одной подсистемы от другой. Однако
использовать такие подсистемы становиться довольно сложно. Один из способов решения
этой проблемы является паттерн Facade.

Требуется для реализации:
1. Класс Facade предоставляющий унифицированный доступ для классов подсистемы;
2. Класс подсистемы SubSystemA;
3. Класс подсистемы SubSystemB;
4. Класс подсистемы SubSystemC.

Идея реализации фасада может быть применена в различных областях:
API-интерфейсы: Создание оболочки для внешнего API с целью облегчения работы с ним.
Фасад может скрыть сложности и специфику внешнего интерфейса, предоставив простой набор методов для взаимодействия.

Управление ресурсами: Например, система управления ресурсами памяти или файловой системы может использовать фасад
для обеспечения простого доступа и манипуляций с ресурсами.

UI-интерфейсы: Фасад может быть применен для создания прослойки,
скрывающей сложные взаимодействия с пользовательским интерфейсом.

Например, если калькулятор имеет множество операций (сложение, вычитание, умножение, деление, возведение в степень и т.д.)
и дополнительные функции (извлечение квадратного корня, вычисление факториала и т.п.),
можно создать фасад, который обеспечит более простой и понятный интерфейс для работы с этими функциями.
*/

// Создадим калькулятор
func NewCalc() *Calculator {
	return &Calculator{
		add:  &Additione{},
		sub:  &Subtraction{},
		mult: &Multiplication{},
		div:  &Division{},
	}

}

// Имплементим калькулятор с фасадом
type Calculator struct {
	add  *Additione
	sub  *Subtraction
	mult *Multiplication
	div  *Division
}

// CanDo возвращает результат выполнения всех операций калькулятора.
func (c Calculator) CanDo() string {
	// CanDo возвращает результат выполнения всех операций калькулятора.
	result := []string{
		c.add.Add(),
		c.sub.Subtraction(),
		c.mult.Multiplication(),
		c.div.Division(),
	}
	return strings.Join(result, "\n")
}

// Additione - класс для выполнения операции сложения.
type Additione struct {
	nums []int
}

// Add выполняет операцию сложения для чисел, хранящихся в nums.
func (a Additione) Add() string {
	return fmt.Sprintf("Performing addition of %d and %d", a.nums[0], a.nums[1])
}

// Subtraction represents a type for performing subtraction operations.
type Subtraction struct {
}

// Subtraction is a method of the Subtraction struct.
// It performs subtraction.
func (s Subtraction) Subtraction() string {
	return "Performing subtraction"
}

type Multiplication struct {
}

func (m Multiplication) Multiplication() string {
	return "Performing multiplication"
}

type Division struct {
}

func (d Division) Division() string {
	return "Performing division"

}

func main() {
	// Создаем экземпляр калькулятора.
	calc := NewCalc()

	// Задаем числа для операции сложения через поле nums.
	calc.add.nums = []int{5, 10}

	// Выводим результат выполнения операции деления.
	fmt.Println(calc.div.Division())

	// Выводим результат выполнения всех операций калькулятора.
	fmt.Println(calc.CanDo())
}

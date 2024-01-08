package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн Builder относится к порождающим паттернам уровня объекта

Паттерн Builder определяет процесс поэтапного построения сложного продукта. После того
как будет построена последняя его часть, продукт можно использовать.

Важно понимать, что сложный объект это не обязательно объект оперирующий несколькими
другими объектами в смысле ООП. Например, нам нужно получить документ состоящий из
заголовка, введения, содержания и заключения. Наш документ, это сложный объект. Что бы
был какой-то единый порядок составления документа, мы будем использовать паттерн
Builder.

*/

type Page struct {
	Head string
	Body string
	Foot string
}

// PageBuilderI определяет контракт для построения страницы
type PageBuilderI interface {
	Head(val string) PageBuilderI // Устанавливает заголовок страницы
	Body(val string) PageBuilderI // Устанавливает тело страницы
	Foot(val string) PageBuilderI // Устанавливает футер страницы
	Build() Page                  // Собирает страницу с установленными значениями
}

// pageBuilder - стандартная реализация PageBuilderI
type pageBuilder struct {
	head string
	body string
	foot string
}

// NewPageBuilder создает новый экземпляр pageBuilder
func NewPageBuilder() PageBuilderI {
	return &pageBuilder{} // Возвращает новый экземпляр pageBuilder
}

// Head добавляет заголовок страницы
func (b pageBuilder) Head(val string) PageBuilderI {
	b.head = val // Устанавливает заголовок
	return b     // Возвращает PageBuilderI (реализация pageBuilder)
}

// Body добавляет тело страницы
func (b pageBuilder) Body(val string) PageBuilderI {
	b.body = val // Устанавливает тело
	return b     // Возвращает PageBuilderI (реализация pageBuilder)
}

// Foot добавляет футер страницы
func (b pageBuilder) Foot(val string) PageBuilderI {
	b.foot = val // Устанавливает футер
	return b     // Возвращает PageBuilderI (реализация pageBuilder)
}

// Build завершает построение страницы
func (b pageBuilder) Build() Page {
	return Page{ // Возвращает новый экземпляр Page с установленными значениями
		Head: b.head,
		Body: b.body,
		Foot: b.foot,
	}
}

// defaultPageBuilder - реализация PageBuilderI с значениями по умолчанию
type defaultPageBuilder struct {
	pageBuilder // Встраивает pageBuilder
}

// NewDefaultPageBuilder создает новый экземпляр defaultPageBuilder
func NewDefaultPageBuilder() PageBuilderI {
	return &defaultPageBuilder{} // Возвращает новый экземпляр defaultPageBuilder
}

// Build создает страницу с заданными значениями по умолчанию
func (b defaultPageBuilder) Build() Page {
	return Page{ // Возвращает новый экземпляр Page с установленными значениями по умолчанию
		Head: "Стандартный заголовок",
		Body: "Стандартный боди",
		Foot: "Стандартный футер",
	}
}

func main() {
	// Создаем страницу с пользовательскими значениями
	pageBuilder := NewPageBuilder()
	page := pageBuilder.Head("Заголовок").Body("Тело письма").Foot("Подвал").Build() // Строим страницу с установленными значениями
	fmt.Println(page)                                                                // Выводим результат

	// Создаем страницу с значениями по умолчанию
	defPageBuilder := NewDefaultPageBuilder()
	defPage := defPageBuilder.Build() // Строим страницу со значениями по умолчанию
	fmt.Println(defPage)              // Выводим результат
}

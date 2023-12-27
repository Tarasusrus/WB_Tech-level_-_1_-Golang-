/*
Реализовать паттерн «адаптер» на любом примере.
*/

package main

import "fmt"

// OldTextDocument интерфейс старой системы для работы с текстовыми документами
type OldTextDocument interface {
	OpenText() string
	SaveText(content string)
}

// NewDocument интерфейс новой системы для работы с документами
type NewDocument interface {
	Open() string
	Save(content string)
}

// TextDoc структура, реализующая OldTextDocument
type TextDoc struct {
	content string
}

func (d *TextDoc) OpenText() string {
	return d.content
}

func (d *TextDoc) SaveText(content string) {
	d.content = content
}

// DocumentAdapter структура, которая адаптирует NewDocument к OldTextDocument
type DocumentAdapter struct {
	doc NewDocument
}

func (a *DocumentAdapter) OpenText() string {
	return a.doc.Open()
}
func (a *DocumentAdapter) SaveText(content string) {
	a.doc.Save(content)
}

// FancyDoc структура, реализующая NewDocument
type FancyDoc struct {
	data string
}

func (d *FancyDoc) Open() string {
	return d.data
}

func (d *FancyDoc) Save(content string) {
	d.data = content
}
func main() {
	// Создаем новый документ с использованием новой системы
	newDoc := &FancyDoc{data: "Hello, New World!"}
	// Адаптируем новый документ к старой системе
	adapter := &DocumentAdapter{doc: newDoc}
	// Теперь мы можем использовать адаптер с интерфейсом старой системы
	fmt.Println("Content:", adapter.OpenText())
	adapter.SaveText("Hello, Adapted World!")
	fmt.Println("Updated Content:", adapter.OpenText())
}

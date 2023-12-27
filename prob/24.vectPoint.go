/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

package main

import (
	"fmt"
	"math"
)

// Point - структура, представляющая точку в двумерном пространстве
type Point struct {
	x, y float64 // Инкапсулированные параметры x и y
}

// NewPoint - конструктор для Point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Distance - метод, вычисляющий расстояние между двумя точками
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))

}
func main() {
	// Создаем две точки с использованием конструктора
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Вычисляем и выводим расстояние между двумя точками
	distance := point1.Distance(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

// Тип для представления точки на плоскости
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Метод для вычисления расстояния между двумя точками
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(0, 0)
	p1.x = 5
	fmt.Printf("Расстояние между точками (%.2f, %.2f) и (%.2f, %.2f) равно %.2f\n", p1.x, p1.y, p2.x, p2.y, p1.Distance(p2))
}

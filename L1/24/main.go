package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	fmt.Println("Расстояние:", p1.Distance(p2)) // 5
}

// Point — структура с приватными полями
type Point struct {
	x float64
	y float64
}

// Метод Distance — расстояние до другой точки
func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

// Конструктор
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

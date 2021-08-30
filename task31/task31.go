package main

import (
	"fmt"
	"math"
)

// Point - структура точки
type Point struct {
	X, Y float64
}

// NewPoint - "конструктор" точки
func NewPoint (x, y float64) *Point {
	return &Point{x, y}
}

// DistanceTo - метод расчета расстояния до другой точки
func (a *Point) DistanceTo(b *Point) float64 {
	return math.Sqrt(math.Pow(a.X - b.X, 2) + math.Pow(a.Y - b.Y, 2))
}

func main() {
	var a = NewPoint(10, 10)
	var b = NewPoint(5, 10)

	fmt.Println("Point a:", a)
	fmt.Println("Point b:", b)
	fmt.Println("Distance between a and b:", a.DistanceTo(b))

}
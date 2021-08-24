package main

import (
	"fmt"
)

// =============== Интерйес точки ======================
type iPoint interface {
	getX() float64
	getY() float64
	move(float64, float64)
}


// =========== Обычная точка с методами ====================
// Реализует интерфейс точки
type point struct {
	x float64
	y float64
}

func (p *point) getX() float64 {
	return p.x
}

func (p *point) getY() float64 {
	return p.y
}

func (p *point) move(dx, dy float64) {
	p.x += dx
	p.y += dy
}

// =================== Город с методами ======================
// Не совместим с интерфейсом точки
type city struct {
	name string
	latitude float64
	longitude float64
}

func (c *city) getLatitude() float64 {
	return c.latitude
}

func (c *city) getLongutude() float64 {
	return c.longitude
}

// ====================== Адаптор города под интерфейс точки ====================
type cityAdapter struct {
	city
}

func (ca *cityAdapter) getX() float64 {
	return ca.getLatitude()
}

func (ca *cityAdapter) getY() float64 {
	return ca.getLongutude()
}

func (ca *cityAdapter) move(dx, dy float64) {
	ca.latitude += dx
	ca.longitude += dy
}


func main() {
	var p1 iPoint = &point{354, 538}				// Обычаня точка в переменной интерфейса точки
	var c = city{"Moscow", 55.755833, 37.617222}	// Город

	fmt.Printf("Point: point.x = %f, point.y = %f\n", p1.getX(), p1.getY())
	fmt.Printf("City: city.name = %s, city.latitude = %f, city.longitude = %f\n", c.name, c.getLatitude(), c.getLongutude())

	var p2 iPoint = &cityAdapter{c}	// Адаптор города, записанный в переменную интерфейса точки
	fmt.Printf("CityAdapter: point2.x = %f, point2.y = %f", p2.getX(), p2.getY())
}
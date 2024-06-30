package main

import (
	"fmt"
	"math"
)

// создаем структуру Point , которая будет хранить координаты точки
type Point struct {
	x, y float64
}

// функция принимает два числв 'x' и 'y' и возвращает ссылку на новую точку
func new_point(x, y float64) *Point {
	return &Point{x: x, y: y} // создаем новую точку с координатами x и y и возвращаем ее
}

// метод, который возвращает расстояние между двумя точками
func (p *Point) distance(q *Point) float64 { // p и q - точки, между которымм мы хотим найти расстояние
	// формула для вычисления расстояния между двумя точками:
	return math.Sqrt((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y))
}

func main() {
	var x1, y1, x2, y2 float64
	fmt.Println("Введите координаты точки x1, y1 и x2, y2:")
	fmt.Scan(&x1, &y1, &x2, &y2)

	p := new_point(x1, y1)
	q := new_point(x2, y2)
	fmt.Println(p.distance(q))
}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	x, y float64
}

func (p Point) getX() float64 {
	return p.x
}

func (p Point) getY() float64 {
	return p.y
}

func calculateDistance(p1, p2 Point) float64 {
	x := p1.getX() - p2.getX()
	y := p1.getY() - p2.getY()
	return math.Sqrt(x*x + y*y)
}

func createNewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p1 := createNewPoint(2.3, 5.4)
	p2 := createNewPoint(4.7, -1.2)
	fmt.Println(calculateDistance(p1, p2))
}

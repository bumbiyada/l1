package main

import (
	"fmt"
	"math"
)

// structure of point
type Point struct {
	x float64
	y float64
}

// set point to axis
func (p *Point) Set(x, y float64) {
	p.x = x
	p.y = y
}

// find out distance between two points
func (p *Point) Distance(t *Point) (res float64) {
	abs_x := math.Abs(p.x + t.x)
	abs_y := math.Abs(p.y + t.y)
	return math.Hypot(abs_x, abs_y)
}
func main() {
	var a, b Point
	a.Set(0, -3)
	b.Set(-4, 0)
	result := a.Distance(&b)
	fmt.Println(result)
}

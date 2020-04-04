package structs

import (
	"math"
)

type Shape interface {
	Area() (float64)
}

type Rectangle struct {
	Base, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Base + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Base * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

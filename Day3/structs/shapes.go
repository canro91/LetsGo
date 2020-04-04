package structs

import (
	"math"
)

type Rectangle struct {
	Base, Height float64
}

type Circle struct {
	Radius float64
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

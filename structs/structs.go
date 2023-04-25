package structs

import (
	"math"
)

type Rectangle struct {
	Height float64
	Width float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
package geometry

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

// First letter is a convention
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}

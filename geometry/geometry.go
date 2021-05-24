package geometry

import "math"

// Interface resolution is implicit
// If the type you pass in matches what the interface is asking for, it will compile
type Shape interface {
	Area() float64
}

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

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}

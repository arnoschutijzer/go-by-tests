package geometry

type Rectangle struct {
	Height float64
	Width  float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}

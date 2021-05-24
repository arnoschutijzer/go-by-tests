package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	t.Run("Calculates perimeter of a square", func(t *testing.T) {
		rectangle := Rectangle{
			Height: 10.0,
			Width:  10.0,
		}
		got := Perimeter(rectangle)
		want := 40.0

		assert.Equal(t, want, got)
	})

	t.Run("Calculates perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			Height: 20.0,
			Width:  10.0,
		}
		got := Perimeter(rectangle)
		want := 60.0

		assert.Equal(t, want, got)
	})
}

func TestArea(t *testing.T) {

	// see https://github.com/golang/go/wiki/TableDrivenTests
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Height: 10.0, Width: 10.0}, want: 100.0},
		{shape: Rectangle{Height: 10.0, Width: 6.0}, want: 60.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assert.Equal(t, tt.want, got)
	}
}

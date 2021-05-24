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
		{Rectangle{10.0, 10.0}, 100.0},
		{Rectangle{10.0, 6.0}, 60.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assert.Equal(t, tt.want, got)
	}
}

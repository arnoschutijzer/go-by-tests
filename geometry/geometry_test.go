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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		assert.Equal(t, want, got)
	}

	t.Run("Calculates area of a square", func(t *testing.T) {
		rectangle := Rectangle{
			Height: 10.0,
			Width:  10.0,
		}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Calculates area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			Height: 10.0,
			Width:  6.0,
		}
		checkArea(t, rectangle, 60.0)
	})

	t.Run("Calculates area of a circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumEquals22(t *testing.T) {
	elements := [...]int{
		4,
		5,
		6,
		7,
	}
	got := Sum(elements)
	want := 22

	assert.Equal(t, want, got)
}

func TestSumEquals44(t *testing.T) {
	elements := [...]int{
		22,
		11,
		10,
		1,
	}

	got := Sum(elements)
	want := 44

	assert.Equal(t, want, got)
}

func TestSumSlice(t *testing.T) {
	t.Run("Slice with size 4", func(t *testing.T) {
		elements := []int{
			1, 2, 3, 4,
		}

		got := SumSlice(elements)
		want := 10

		assert.Equal(t, want, got)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("calculates total of 2 slices", func(t *testing.T) {
		firstSlice := []int{
			1, 2,
		}
		secondSlice := []int{
			3, 4,
		}

		got := SumAll(firstSlice, secondSlice)
		want := []int{3, 7}

		assert.Equal(t, want, got)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("total some slices", func(t *testing.T) {
		got := SumAllTails(
			[]int{
				1, 2, 3,
			},
			[]int{
				4, 5, 6,
			},
		)
		want := []int{5, 11}

		assert.Equal(t, want, got)
	})

	t.Run("nil-safe slice summing", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 9})
		want := []int{0, 9}

		assert.Equal(t, want, got)
	})

	t.Run("doesnt manipulate my original slices", func(t *testing.T) {
		firstSlice := []int{1, 2, 3, 4}
		secondSlice := []int{5, 6, 7, 8}

		got := SumAllTails(firstSlice, secondSlice)
		want := []int{9, 21}

		assert.Equal(t, want, got)
		assert.Equal(t, 4, len(firstSlice))
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{12, 3, 12341234, 3213}, []int{123123, 12312312, 3123213})
	}
}

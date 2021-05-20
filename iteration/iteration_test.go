package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat A 5 times", func(t *testing.T) {
		got := Repeat("A", 5)
		want := "AAAAA"

		assert.Equal(t, got, want)
	})

	t.Run("repeat A 6 times", func(t *testing.T) {
		got := Repeat("A", 6)
		want := "AAAAAA"

		assert.Equal(t, got, want)
	})
}

func ExampleRepeat() {
	theLetterAFiveTimes := Repeat("A", 5)
	fmt.Println(theLetterAFiveTimes)
	// Output: AAAAA
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

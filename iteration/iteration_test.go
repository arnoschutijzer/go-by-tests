package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat A 5 times", func(t *testing.T) {
		got := Repeat("A")
		want := "AAAAA"

		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

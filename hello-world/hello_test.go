package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		// report the line number of the source code
		// and not from inside this helper.
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Arno", "")
		want := "Hello, Arno"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", SpanishSelector)
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
}

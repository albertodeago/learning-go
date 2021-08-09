package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("It should let saying 'Hello' to people", func(t *testing.T) {
		got := Hello("Gamera", "")
		want := "Hello, Gamera"

		assertCorrectMessage(t, got, want)
	})

	t.Run("It should fallback to 'Hello, world' if an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("It should work in Spanish", func(t *testing.T) {
		got := Hello("Maria", "Spanish")
		want := "Hola, Maria"

		assertCorrectMessage(t, got, want)
	})

	t.Run("It should work in French", func(t *testing.T) {
		got := Hello("Baguette", "French")
		want := "Bonjour, Baguette"

		assertCorrectMessage(t, got, want)
	})
}

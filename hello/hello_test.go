package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Tobi", "")
		want := "Hello, Tobi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to world when empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Carlos", "Spanish")
		want := "Hola, Carlos"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Anna", "French")
		want := "Allo, Anna"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French with empty string passed to name", func(t *testing.T) {
		got := Hello("", "French")
		want := "Allo, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

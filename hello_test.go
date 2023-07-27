package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Tobi")
		want := "Hello, Tobi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to world when empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tobi")
	want := "Hello, Tobi"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

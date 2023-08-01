package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Says greet", func(t *testing.T) {
		// create an empty buffer
		// buffers obey the Writer interface
		// they implement Write(p []byte) (n int, err error)
		buffer := bytes.Buffer{}

		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("wanted %q but got %q", got, want)
		}
	})

}

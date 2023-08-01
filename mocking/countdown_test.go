package countdown

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCoutndown(t *testing.T) {
	t.Run("Prints all three numbers to the buffer", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		numberOfCalls := 3

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("Insufficient calls (%d) to sleeper; wanted %d", spySleeper.Calls, numberOfCalls)
		}
	})
}

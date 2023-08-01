package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "Sleep"
const write = "Write"

type SpyCountdownOperations struct {
	Calls []string
}

type SpySleeper struct {
	Calls int
}

type SpyTime struct {
	DurationSlept time.Duration
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write([]byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.DurationSlept = duration
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

	t.Run("Calls the operations in the right order", func(t *testing.T) {
		spy := &SpyCountdownOperations{}

		Countdown(spy, spy)

		want := []string{
			"Write",
			"Sleep",
			"Write",
			"Sleep",
			"Write",
			"Sleep",
			"Write",
		}

		if !reflect.DeepEqual(want, spy.Calls) {
			t.Errorf("wanted calls to %v, got %v", want, spy.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("It sleeps for the right duration", func(t *testing.T) {
		sleepDuration := 5 * time.Second
		spyTime := &SpyTime{}

		configurableSleeper := &ConfigurableSleeper{sleepDuration, spyTime.Sleep}
		configurableSleeper.Sleep()

		if spyTime.DurationSlept != sleepDuration {
			t.Errorf("should have slept for %d instead of %v", sleepDuration, spyTime.DurationSlept)
		}
	})
}

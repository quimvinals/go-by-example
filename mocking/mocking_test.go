package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpyCountdownOperations{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

	if len(spySleeper.Calls) != 3 {
		t.Errorf("Expected 3 calls to sleep function, got %d", len(spySleeper.Calls))
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeper := &SpyCountdownOperations{}

		Countdown(spySleeper, spySleeper)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
		}

	})
}

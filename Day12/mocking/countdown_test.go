package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyCountdown struct {
	Calls []string
}

type SpyTime struct {
	durationSlep time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlep = duration
}

func (s *SpyCountdown) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdown) Write(p []byte) (n int, e error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("Output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdown{}

		Countdown(buffer, spySleeper)

		actual := buffer.String()
		expected := `3
2
1
Go!`

		if expected != actual {
			t.Errorf("Expected %q but was %q", expected, actual)
		}
	})

	t.Run("Sleep", func(t *testing.T) {
		spySleeper := &SpyCountdown{}

		Countdown(spySleeper, spySleeper)

		expected := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expected, spySleeper.Calls) {
			t.Errorf("Expected %v but was %v", expected, spySleeper.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlep != sleepTime {
		t.Errorf("Expected %v but was %v", sleepTime, spyTime.durationSlep)
	}
}

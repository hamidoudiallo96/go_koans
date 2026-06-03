package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownCounter = 3
	finalWord        = "Go!"
	sleep            = "sleep"
	write            = "write"
)

type Sleeper interface {
	Sleep()
}

// SpySleeper mocks Sleep functionality
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownCounter; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}

	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 5 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}

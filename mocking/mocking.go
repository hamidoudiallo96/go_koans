package main

import (
	"fmt"
	"io"
	"os"
)

const (
	countdownCounter = 3
	finalWord        = "Go!"
)

type Sleeper interface {
	Sleep()
}

type SleepCounter struct {
	Calls int
}

func (s *SleepCounter) Sleep() {
	s.Calls++
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownCounter; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleepCounter := SleepCounter{}
	Countdown(os.Stdout, &sleepCounter)
}

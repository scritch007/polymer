package main

import (
	"code.palmstonegames.com/polymer"
	"time"
)

func init() {
	polymer.Register("tick-timer", &Timer{})
}

type Timer struct {
	*polymer.Proto

	Time time.Time `polymer:"bind"`
}

func (t *Timer) Created() {
	go func() {
		for {
			// Set the clock
			t.Time = time.Now()

			// Notify
			t.Notify("Time")

			// Wait
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func (t *Timer) ComputeTime() string {
	return t.Time.String()
}

func main() {}

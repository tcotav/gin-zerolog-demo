package data

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

// Two sample functions to show how we're logging

func GetThing() (string, error) {
	log.Info().Str("foo", "bar").Msg("Hello Thing")
	return "thing", nil
}

func TimeSomething() {
	// show how to do a timing
	start := time.Now()
	time.Sleep(1 * time.Second)
	// we opt to include which app, which function call, and then the time elapsed
	log.Info().Str("App", "logtest").Str("Call", "TimeSomething").Float32("Duration", float32(time.Since(start).Seconds())).Msg("")

	// this is how we'd log an error including app and function call name
	log.Error().Err(errors.New("this is an error")).Str("App", "logtest").Str("Call", "TimeSomething").Msg("error")
}
